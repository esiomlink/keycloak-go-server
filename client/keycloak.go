package client

import (
	"context"
	"fmt"
	"os"

	"github.com/Nerzal/gocloak/v11"
)

type Keycloak struct {
	Gocloak      gocloak.GoCloak // keycloak client
	ClientId     string          // clientId specified in Keycloak
	ClientSecret string          // client secret specified in Keycloak
	Realm        string          // realm specified in Keycloak
	AccesToken   string
}

func NewKeycloak() *Keycloak {
	gocloak := gocloak.NewClient(os.Getenv("KEYCLOAK_NODE_BASE_URL"))
	clientID := os.Getenv("KEYCLOAK_NODE_CLIENT_ID")
	clientSecret := os.Getenv("KEYCLOAK_NODE_CLIENT_SECRET")
	realm := os.Getenv("KEYCLOAK_NODE_REALM_NAME")
	ctx := context.Background()

	token, err := gocloak.LoginClient(ctx, clientID, clientSecret, realm)
	if err != nil {
		panic("Login failed:" + err.Error())
	}

	rptResult, err := gocloak.RetrospectToken(ctx, token.AccessToken, clientID, clientSecret, realm)
	if err != nil {
		panic("Inspection failed:" + err.Error())
	}

	permissions := rptResult

	fmt.Printf("AccessToken => %v/n", token.AccessToken)
	fmt.Printf("permissions => %v/n", permissions)

	return &Keycloak{
		Gocloak:      gocloak,
		ClientId:     clientID,
		ClientSecret: clientSecret,
		Realm:        realm,
		AccesToken:   token.AccessToken,
	}
}

/*
export KEYCLOAK_NODE_BASE_URL="http://localhost:8086/auth"
export KEYCLOAK_NODE_REALM_NAME="medium"
export KEYCLOAK_NODE_CLIENT_SECRET="pZtkW487SYPsb9BW01Te0JUl8YpXlqCv"
export KEYCLOAK_NODE_CLIENT_ID="my-go-service"
*/
