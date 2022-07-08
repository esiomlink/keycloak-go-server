package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"keycloak-go/client"
	"net/http"
	"strings"
	"time"
)

type doc struct {
	Id   string    `json:"id"`
	Num  string    `json:"num"`
	Date time.Time `json:"date"`
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


func (c *Controller)  GetUser(w http.ResponseWriter, r *http.Request) {
	accesToken := r.Header.Get("Authorization")
	accesToken= strings.Replace(accesToken, "Bearer ", "", 1)

	//	token = auth.extractBearerToken(token)
fmt.Printf("token=>>>>>>>>> %v\n",accesToken)
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

func (c *Controller) GetDocs(w http.ResponseWriter, r *http.Request) {

	rs := []*doc{
		{
			Id:   "1",
			Num:  "ABC-123",
			Date: time.Now().UTC(),
		},
		{
			Id:   "2",
			Num:  "ABC-456",
			Date: time.Now().UTC(),
		},
	}

	rsJs, _ := json.Marshal(rs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rsJs)

}
