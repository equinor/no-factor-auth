package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TokenOKResponse ok type
type TokenOKResponse struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    string `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// TokenErrorResponse error type
type TokenErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorCodes       []int  `json:"error_codes"`
	Timestamp        string `json:"timestamp"`
	TraceID          string `json:"trace_id"`
	CorrelationID    string `json:"correlation_id"`
}

type reqParams struct {
	GrantType    string `json:"grant_type" form:"grant_type" query:"grant_type"`
	ClientID     string `json:"client_id" form:"client_id" query:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`
	ExtraClaims  string `json:"extra_claims" form:"extra_claims"`
	Assertion    string `json:"assertion" form:"assertion" query:"assertion"`
}

// Token provides id_token and access_token to anyone who asks
func Token(claimsCustom map[string]interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("claims: %v", claimsCustom)
		p := new(reqParams)
		err := c.Bind(p)
		if err != nil {
			log.Fatalf("binding: %v", err)
			return err
		}
		log.Printf("p: %+v", p)

		claims := make(map[string]interface{})
		claims["iss"] = AuthServer + "/v2.0"
		claims["sub"] = "sub"

		for key, val := range claimsCustom {
			claims[key] = val
		}

		a, err := newToken(claims)
		if err != nil {
			return err
		}
		log.Printf("token: %v", a)

		return c.JSON(http.StatusOK, TokenOKResponse{AccessToken: a})
	}
}
