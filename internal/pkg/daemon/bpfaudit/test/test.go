package main

import (
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/cli"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/daemon/bpfaudit"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/daemon/bpfrecorder"
)

func main() {
	logger := logr.New(&cli.LogSink{})

	b := bpfrecorder.New("", logger, false, false)
	mntns, err := b.FindProcMountNamespace(1)
	if err != nil {
		panic(err)
	}

	audit := bpfaudit.New(logger)
	if err := audit.Load(); err != nil {
		panic(err)
	}

	for {
		val := audit.GetViolationCount(mntns)
		fmt.Printf("Violations for namespace %d: %d\n", mntns, val)
		val = audit.GetViolationCount(42)
		fmt.Printf("Violations for namespace %d: %d\n", 42, val)
		time.Sleep(1 * time.Second)
	}
}
