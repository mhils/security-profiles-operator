package main

import (
	"github.com/go-logr/logr"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/cli"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/daemon/enricher/source"
)

func main() {
	logger := logr.New(&cli.LogSink{})

	bpf := source.NewBpfSource(logger)
	log, err := bpf.StartTail()
	if err != nil {
		panic(err)
	}

	for line := range log {
		logger.Info("violation", "line", line)
	}
}

/*
func getCgroupID() bpfaudit.CgroupId {
	content, _ := os.ReadFile("/proc/self/cgroup")
	re := regexp.MustCompile(`(?m)^0::(.*)$`)

	matches := re.FindStringSubmatch(string(content))
	cgroupPath := ""
	if len(matches) > 1 {
		cgroupPath = matches[1]
	}

	var stat syscall.Stat_t
	_ = syscall.Stat(filepath.Join("/sys/fs/cgroup", cgroupPath), &stat)
	return stat.Ino
}
*/
