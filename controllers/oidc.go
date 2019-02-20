package controllers
import (
	"net/http"

	"github.com/equinor/no-factor-auth/oidc"

	"github.com/labstack/echo/v4"
)

// OidcConfig returns config for host
func OidcConfig(c echo.Context) error {

	hostURL := c.Request().Host
	oidc := oidc.Default()
	oidc.JwksURI = hostURL + "/discovery/keys"
	oidc.Issuer = hostURL
	oidc.AuthorizationEndpoint = hostURL + "/oauth2/authorize"
	return c.JSON(http.StatusOK, oidc)
}
