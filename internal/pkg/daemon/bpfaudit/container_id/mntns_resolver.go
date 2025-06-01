package container_id

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/jellydator/ttlcache/v3"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/util"
)

type Mntns = uint32
type Pid = int
type ContainerId = string
type Profile = string

type MountNamespaceResolver struct {
	logger logr.Logger
	cache  map[Mntns]ContainerId
}

func NewMountNamespaceResolver(logger logr.Logger) *MountNamespaceResolver {
	return &MountNamespaceResolver{
		logger: logger,
		cache:  make(map[Mntns]ContainerId),
	}
}

func (b *MountNamespaceResolver) Get(mntns Mntns, pid Pid) (ContainerId, error) {
	// We cache mntns -> container id...
	containerId, ok := b.cache[mntns]
	if !ok {
		// ...but if we don't have it in cache, we use the pid to lookup
		// the container id via /proc/<pid>/cgroup.
		var err error
		// ContainerIDForPID wants a cache, but we don't want caching at this level.
		unused := ttlcache.New[string, string]()
		containerId, err = util.ContainerIDForPID(unused, pid)
		if err != nil {
			return "", fmt.Errorf("resolve profile to container id: %w", err)
		}
		b.cache[mntns] = containerId
	}
	return containerId, nil
}
