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
