package main

import (
	"time"

	"github.com/go-logr/logr"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/cli"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/daemon/bpfaudit"
)

func main() {
	logger := logr.New(&cli.LogSink{})

	audit := bpfaudit.New(logger)
	if err := audit.Load(); err != nil {
		panic(err)
	}

	for {
		val := audit.GetViolationCount("")
		logger.Info("violations", "violations", val)
		time.Sleep(1 * time.Second)
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
