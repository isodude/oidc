package op

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/zitadel/logging"
	httphelper "github.com/isodude/oidc/v2/pkg/http"
	"github.com/isodude/oidc/v2/pkg/oidc"
)

type Introspector interface {
	Decoder() httphelper.Decoder
	Crypto() Crypto
	Storage() Storage
	AccessTokenVerifier(context.Context) AccessTokenVerifier
}

type IntrospectorJWTProfile interface {
	Introspector
	JWTProfileVerifier(context.Context) JWTProfileVerifier
}

func introspectionHandler(introspector Introspector) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Introspect(w, r, introspector)
	}
}

func Introspect(w http.ResponseWriter, r *http.Request, introspector Introspector) {
	response := oidc.NewIntrospectionResponse()
	token, clientID, err := ParseTokenIntrospectionRequest(r, introspector)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
                logging.Debugf("introspect error: %v", err)

		return
	}
	tokenID, subject, ok := getTokenIDAndSubject(r.Context(), introspector, token)
	if !ok {
		httphelper.MarshalJSON(w, response)
		return
	}
	err = introspector.Storage().SetIntrospectionFromToken(r.Context(), response, tokenID, subject, clientID)
	if err != nil {
		httphelper.MarshalJSON(w, response)
                logging.Debugf("introspect error: %v", err)
		return
	}
	response.SetActive(true)
	httphelper.MarshalJSON(w, response)
}

func ParseTokenIntrospectionRequest(r *http.Request, introspector Introspector) (token, clientID string, err error) {
	err = r.ParseForm()
	if err != nil {
		return "", "", errors.New("unable to parse request")
	}
	req := new(struct {
		oidc.IntrospectionRequest
		oidc.ClientAssertionParams
	})
	err = introspector.Decoder().Decode(req, r.Form)
	if err != nil {
		return "", "", errors.New("unable to parse request")
	}
	if introspectorJWTProfile, ok := introspector.(IntrospectorJWTProfile); ok && req.ClientAssertion != "" {
		profile, err := VerifyJWTAssertion(r.Context(), req.ClientAssertion, introspectorJWTProfile.JWTProfileVerifier(r.Context()))
		if err == nil {
			return req.Token, profile.Issuer, nil
		}
	}
	clientID, clientSecret, ok := r.BasicAuth()
	if ok {
		clientID, err = url.QueryUnescape(clientID)
		if err != nil {
			return "", "", errors.New("invalid basic auth header")
		}
		clientSecret, err = url.QueryUnescape(clientSecret)
		if err != nil {
			return "", "", errors.New("invalid basic auth header")
		}
		if err := introspector.Storage().AuthorizeClientIDSecret(r.Context(), clientID, clientSecret); err != nil {
			return "", "", err
		}
		return req.Token, clientID, nil
	}
	return "", "", errors.New("invalid authorization")
}
