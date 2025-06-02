package source

import "sigs.k8s.io/security-profiles-operator/internal/pkg/daemon/enricher/types"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate -header ../../../../../hack/boilerplate/boilerplate.generatego.txt
//counterfeiter:generate . AuditLineSource
type AuditLineSource interface {
	StartTail() (chan *types.AuditLine, error)
	TailErr() error
}
