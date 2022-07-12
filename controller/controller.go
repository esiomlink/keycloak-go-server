package controller

import (
	"context"
	"encoding/json"
	"io/ioutil"
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
type getUserBody struct {
	UserId string `json:"userId"`
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

func GetAccesToken(r *http.Request) string {
	accesToken := r.Header.Get("Authorization")
	accesToken = strings.Replace(accesToken, "Bearer ", "", 1)
	return accesToken
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

func (c *Controller) GetUserById(w http.ResponseWriter, r *http.Request) {
		payload := getUserBody{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
		}

		err = json.Unmarshal(body, &payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
		}


	rs, err := c.keycloak.Gocloak.GetUserByID(
		context.Background(),
		c.keycloak.AccesToken,
		c.keycloak.Realm,
		payload.UserId,
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

func (c *Controller) GetUsers(w http.ResponseWriter, r *http.Request) {

	rs, err := c.keycloak.Gocloak.GetUsers(
		context.Background(),
		c.keycloak.AccesToken,
		c.keycloak.Realm,
		gocloak.GetUsersParams{},
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
