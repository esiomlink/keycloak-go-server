package client

import (
	"os"

	"github.com/Nerzal/gocloak/v11"
)

type Keycloak struct {
	Gocloak      gocloak.GoCloak // keycloak client
	ClientId     string          // clientId specified in Keycloak
	ClientSecret string          // client secret specified in Keycloak
	Realm        string          // realm specified in Keycloak
}

func NewKeycloak() *Keycloak {
	return &Keycloak{
		Gocloak:      gocloak.NewClient(os.Getenv("KEYCLOAK_NODE_BASE_URL")),
		ClientId:     os.Getenv("KEYCLOAK_NODE_CLIENT_ID"),
		ClientSecret: os.Getenv("KEYCLOAK_NODE_CLIENT_SECRET"),
		Realm:       os.Getenv("KEYCLOAK_NODE_REALM_NAME"),
	}
}
/* 
export KEYCLOAK_NODE_BASE_URL="http://localhost:8086"
export KEYCLOAK_NODE_REALM_NAME="medium"
export KEYCLOAK_NODE_CLIENT_SECRET="pZtkW487SYPsb9BW01Te0JUl8YpXlqCv"
export KEYCLOAK_NODE_CLIENT_ID="my-go-service"
*/


