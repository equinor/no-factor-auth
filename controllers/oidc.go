package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// OidcConfig returns config for host
func OidcConfig(c echo.Context) error {
	oidc := openIDConfig{}
	oidc.JwksURI = AuthServer + "/discovery/v2.0/keys"
	oidc.Issuer = AuthServer + "/v2.0"
	oidc.AuthorizationEndpoint = AuthServer + "/oauth2/v2.0/authorize"
	oidc.TokenEndpoint = AuthServer + "/oauth2/v2.0/token"
	return c.JSON(http.StatusOK, oidc)
}

type openIDConfig struct {
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	JwksURI               string `json:"jwks_uri"`
}
