#!/usr/bin/env bash

# watchexec -e c -r make internal/pkg/daemon/bpfaudit/bpf/audit.bpf.o.amd64
# sudo watchexec -e go,amd64 -r internal/pkg/daemon/bpfaudit/test/test.sh
# watchexec --delay-run 2s -w internal/pkg/daemon/bpfaudit/test head -1 CONTRIBUTING.md
# sudo cat /sys/kernel/debug/tracing/trace_pipe

pushd $(dirname "$0") || exit

CC=gcc CGO_CFLAGS="-I /usr/include/bpf" CGO_LDFLAGS="-lbpf" go build test.go
sudo ./test

popd || exit
