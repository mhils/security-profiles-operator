//go:build linux && !no_bpf && (amd64 || arm64)

package bpfaudit

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	"github.com/aquasecurity/libbpfgo"
	"github.com/go-logr/logr"
)

type BpfAudit struct {
	logger              logr.Logger
	apparmor_violations *libbpfgo.BPFMap
}

func New(logger logr.Logger) *BpfAudit {
	return &BpfAudit{
		logger: logger,
	}
}

func (b *BpfAudit) Load() error {

	b.logger.Info("Loading bpf module...")
	module, err := libbpfgo.NewModuleFromBufferArgs(libbpfgo.NewModuleArgs{
		BPFObjBuff: AuditProgram,
		BPFObjName: "audit.bpf.o",
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
	if b.apparmor_violations, err = module.GetMap("apparmor_violations"); err != nil {
		return fmt.Errorf("get bpf map: %w", err)
	}

	events := make(chan []byte)
	buf, err := module.InitRingBuf("audit_log", events)
	if err != nil {
		return fmt.Errorf("init ringbuf: %w", err)
	}
	buf.Poll(300)

	go func() {
		for val := range events {
			fmt.Printf("Ringbuf event received: %d %v %s\n", len(val), val, val)
		}
	}()

	b.logger.Info("BPF module successfully loaded.")
	return nil
}

func (b *BpfAudit) GetViolationCount(mntns uint32) uint32 {
	val, err := b.apparmor_violations.GetValue(unsafe.Pointer(&mntns))
	if err != nil {
		return 0 // not found = no violations
	}
	return binary.LittleEndian.Uint32(val)
}
