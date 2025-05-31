//go:build linux && !no_bpf && amd64

package bpfaudit

import _ "embed"

//go:embed bpf/audit.bpf.o.amd64
var AuditProgram []byte
