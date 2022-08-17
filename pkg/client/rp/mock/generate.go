package mock

//go:generate mockgen -package mock -destination ./verifier.mock.go github.com/isodude/oidc/pkg/rp Verifier
