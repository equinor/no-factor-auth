package controllers

import (
	"log"
	"net/http"
	"net/url"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var (
	// AuthServer is the url where this server is hosted, including tenant
	AuthServer string
)

func newToken(claims map[string]interface{}) (string, error) {
	defaultClaims := jwt.MapClaims{
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(1 * time.Hour).Unix(),
	}

	for key, value := range claims {
		defaultClaims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, defaultClaims)

	token.Header = map[string]interface{}{
		"typ": "JWT",
		"alg": jwt.SigningMethodRS256.Name,
		"kid": "1",
	}

	tokenString, err := token.SignedString(privateKey())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

type signInReq struct {
	ClientID     string `json:"client_id" query:"client_id"`
	Tenant       string `json:"tenant" query:"tenant"`
	ResponseType string `json:"response_type" query:"response_type"`
	RedirectURI  string `json:"redirect_uri" query:"redirect_uri"`
	State        string `json:"state" query:"state"`
}

// Authorize provides id_token and access_token to anyone who asks
func Authorize(c echo.Context) error {
	r := new(signInReq)
	err := c.Bind(r)
	if err != nil {
		return c.String(400, "bad")
	}

	claims := make(map[string]interface{})
	claims["iss"] = AuthServer + "/v2.0"
	claims["aud"] = r.ClientID
	claims["sub"] = r.ClientID

	tokenString, err := newToken(claims)
	if err != nil {
		return err
	}

	params := url.Values{}
	params.Set("id_token", tokenString)
	params.Set("access_token", tokenString)
	params.Set("state", r.State)
	log.Println(tokenString)

	return c.Redirect(http.StatusFound, r.RedirectURI+"#"+params.Encode())
}
