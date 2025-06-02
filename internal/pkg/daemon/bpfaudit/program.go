//go:build linux && !no_bpf && (amd64 || arm64)

package bpfaudit

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"sync"

	"github.com/aquasecurity/libbpfgo"
	"github.com/go-logr/logr"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/config"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/daemon/bpfaudit/container_id"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/daemon/enricher/source"
)

type BpfAudit struct {
	logger            logr.Logger
	containerInfo     map[container_id.ContainerId]*ContainerInfo
	lockContainerInfo sync.Mutex
	mntnsResolver     *container_id.MountNamespaceResolver
}

type ContainerInfo struct {
	totalViolations uint32
	records         []AppArmorAuditRecord
}

type AppArmorAuditRecord struct {
	request uint32
	name    string
}

func New(logger logr.Logger) *BpfAudit {
	return &BpfAudit{
		logger:        logger,
		containerInfo: make(map[container_id.ContainerId]*ContainerInfo),
		mntnsResolver: container_id.NewMountNamespaceResolver(logger),
	}
}

func (b *BpfAudit) Load() error {

	b.logger.Info("Loading bpf module...")
	module, err := libbpfgo.NewModuleFromBufferArgs(libbpfgo.NewModuleArgs{
		BPFObjBuff: source.AuditProgram,
		BPFObjName: "enricher.bpf.o",
	})
	if err != nil {
		return fmt.Errorf("load bpf module: %w", err)
	}
	if err := module.BPFLoadObject(); err != nil {
		return fmt.Errorf("load bpf object: %w", err)
	}
	if err := module.AttachPrograms(); err != nil {
		return fmt.Errorf("load bpf object: %w", err)
	}

	events := make(chan []byte)
	buf, err := module.InitRingBuf("audit_log", events)
	if err != nil {
		return fmt.Errorf("init ringbuf: %w", err)
	}
	buf.Poll(300)

	go func() {
		for val := range events {
			mntns := binary.LittleEndian.Uint32(val[0:4])
			pid := int(binary.LittleEndian.Uint32(val[4:8]))
			request := binary.LittleEndian.Uint32(val[8:12])
			name := string(val[12 : len(val)-1])
			containerId, err := b.mntnsResolver.Get(mntns, pid)

			if err != nil {
				b.logger.V(config.VerboseLevel).Info("unable to determine container id",
					"mntns", mntns,
					"pid", pid,
					"err", err,
				)
			}
			b.logger.Info("audit log event received",
				"mntns", mntns,
				"pid", pid,
				"request", request,
				"name", name,
				"containerId", containerId,
			)

			b.lockContainerInfo.Lock()
			containerInfo, ok := b.containerInfo[containerId]
			if !ok {
				containerInfo = &ContainerInfo{}
				b.containerInfo[containerId] = containerInfo
			}
			containerInfo.totalViolations += 1
			if len(containerInfo.records) < 50 {
				if len(containerInfo.records) > 0 {
					latest := containerInfo.records[len(containerInfo.records)-1]
					if latest.name == name && latest.request == request {
						continue
					}
				}
				containerInfo.records = append(containerInfo.records, AppArmorAuditRecord{
					request,
					name,
				})
			}

			b.lockContainerInfo.Unlock()
		}
	}()

	b.logger.Info("BPF module successfully loaded.")
	return nil
}

func (b *BpfAudit) GetViolationCount(id container_id.ContainerId) *ContainerInfo {
	b.lockContainerInfo.Lock()
	defer b.lockContainerInfo.Unlock()

	info, ok := b.containerInfo[id]
	if !ok {
		return nil
	}
	return info
}

func (r *AppArmorAuditRecord) Request() string {
	switch r.request {
	case 1:
		return "x"
	case 2:
		return "w"
	case 3:
		return "wx"
	case 4:
		return "r"
	case 5:
		return "rx"
	case 6:
		return "rw"
	case 7:
		return "rwx"
	default:
		return strconv.Itoa(int(r.request))
	}
}
