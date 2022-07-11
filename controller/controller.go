package controller

import (
	"context"
	"encoding/json"
	"keycloak-go/client"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v11"
)

type UsersParams struct {
	email     string
	first     int
	firstName string
	lastName  string
	search    string
	username  string
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
}

type Controller struct {
	keycloak *client.Keycloak
}

func NewController(keycloak *client.Keycloak) *Controller {
	return &Controller{
		keycloak: keycloak,
	}
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {

	rq := &loginRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(rq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jwt, err := c.keycloak.Gocloak.Login(context.Background(),
		c.keycloak.ClientId,
		c.keycloak.ClientSecret,
		c.keycloak.Realm,
		rq.Username,
		rq.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	rs := &loginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	}

	rsJs, _ := json.Marshal(rs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rsJs)
}

func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	accesToken := r.Header.Get("Authorization")
	accesToken = strings.Replace(accesToken, "Bearer ", "", 1)

	//	token = auth.extractBearerToken(token)
	rs, err := c.keycloak.Gocloak.GetUserInfo(
		context.Background(),
		accesToken,
		c.keycloak.Realm,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	rsJs, _ := json.Marshal(rs)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rsJs)
}

type user struct {
	email string
}

func (c *Controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	accesToken := r.Header.Get("Authorization")
	accesToken = strings.Replace(accesToken, "Bearer ", "", 1)
	params := gocloak.GetUsersParams{}

	rs, err := c.keycloak.Gocloak.GetUsers(
		context.Background(),
		accesToken,
		c.keycloak.Realm,
		params,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	rsJs, _ := json.Marshal(rs)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rsJs)
}
