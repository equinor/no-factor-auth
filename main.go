package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/equinor/no-factor-auth/controllers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	Version = ""
	v       = flag.Bool("version", false, "Display version")
)

func version() {
	fmt.Println("Version:", Version)
	fmt.Println("Go Version:", runtime.Version())
	os.Exit(0)
}

func main() {

	flag.Parse()

	if *v {
		version()
	}

	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file", err)
	}

	var claims map[string]interface{}
	rawClaims := os.Getenv("TOKEN_ENDPOINT_CLAIMS")
	fmt.Println(rawClaims)
	if len(rawClaims) > 0 {
		err = json.Unmarshal([]byte(rawClaims), &claims)
		if err != nil {
			log.Fatalf(" %v", err)
		}
	}

	controllers.AuthServer = os.Getenv("AUTH_SERVER")
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/common/v2.0/.well-known/openid-configuration", controllers.OidcConfig)
	e.GET("/common/discovery/v2.0/keys", controllers.Jwks)
	e.GET("/common/oauth2/v2.0/authorize", controllers.Authorize)
	// ec := `{"aud":"https://storage.azure.com","iss":"https://sts.windows.net/common"}`
	e.POST("/common/oauth2/v2.0/token", controllers.Token(claims))

	e.Logger.Fatal(e.Start("0.0.0.0:8089"))
}
