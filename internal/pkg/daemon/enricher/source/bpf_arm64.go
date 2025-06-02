//go:build linux && !no_bpf && arm64

package source

import _ "embed"

//go:embed bpf/enricher.bpf.o.arm64
var AuditProgram []byte
