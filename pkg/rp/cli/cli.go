package cli

import (
	"context"
	"net/http"

	"github.com/caos/oidc/pkg/oidc"
	"github.com/caos/oidc/pkg/rp"
	"github.com/caos/oidc/pkg/utils"
)

const (
	loginPath = "/login"
)

func CodeFlow(relayingParty rp.RelayingParty, callbackPath, port string, stateProvider func() string) *oidc.Tokens {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var token *oidc.Tokens
	callback := func(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens, state string) {
		token = tokens
		msg := "<p><strong>Success!</strong></p>"
		msg = msg + "<p>You are authenticated and can now return to the CLI.</p>"
		w.Write([]byte(msg))
	}
	http.Handle(loginPath, rp.AuthURLHandler(stateProvider, relayingParty))
	http.Handle(callbackPath, rp.CodeExchangeHandler(callback, relayingParty))

	utils.StartServer(ctx, port)

	utils.OpenBrowser("http://localhost:" + port + loginPath)

	return token
}
