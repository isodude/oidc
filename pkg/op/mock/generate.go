package mock

//go:generate mockgen -package mock -destination ./storage.mock.go github.com/isodude/oidc/pkg/op Storage
//go:generate mockgen -package mock -destination ./authorizer.mock.go github.com/isodude/oidc/pkg/op Authorizer
//go:generate mockgen -package mock -destination ./client.mock.go github.com/isodude/oidc/pkg/op Client
//go:generate mockgen -package mock -destination ./configuration.mock.go github.com/isodude/oidc/pkg/op Configuration
//go:generate mockgen -package mock -destination ./discovery.mock.go github.com/isodude/oidc/pkg/op DiscoverStorage
//go:generate mockgen -package mock -destination ./signer.mock.go github.com/isodude/oidc/pkg/op SigningKey,Key
//go:generate mockgen -package mock -destination ./key.mock.go github.com/isodude/oidc/pkg/op KeyProvider
