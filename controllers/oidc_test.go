package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestOidc(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	AuthServer = "http://example.com/common"
	c := e.NewContext(req, rec)
	var oidc openIDConfig
	if assert.NoError(t, OidcConfig(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		json.Unmarshal(rec.Body.Bytes(), &oidc)
		assert.Equal(t, AuthServer+"/discovery/v2.0/keys", oidc.JwksURI)
		assert.Equal(t, AuthServer+"/v2.0", oidc.Issuer)
		assert.Equal(t, AuthServer+"/oauth2/v2.0/authorize", oidc.AuthorizationEndpoint)
		assert.Equal(t, AuthServer+"/oauth2/v2.0/token", oidc.TokenEndpoint)
	}
}
