//go:build linux && !no_bpf && arm64

package bpfaudit

import _ "embed"

//go:embed bpf/audit.bpf.o.arm64
var AuditProgram []byte
