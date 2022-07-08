package client

import "github.com/Nerzal/gocloak/v11"

type Keycloak struct {
	Gocloak      gocloak.GoCloak // keycloak client
	ClientId     string          // clientId specified in Keycloak
	ClientSecret string          // client secret specified in Keycloak
	Realm        string          // realm specified in Keycloak
}

func NewKeycloak() *Keycloak {
	return &Keycloak{
		Gocloak:      gocloak.NewClient("http://localhost:8086"),
		ClientId:     "my-go-service",
		ClientSecret: "pZtkW487SYPsb9BW01Te0JUl8YpXlqCv",
		Realm:        "medium",
	}
}
