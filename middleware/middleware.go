package middleware

import (
	"context"
	
	"fmt"
	"keycloak-go/client"
	"net/http"
	"strings"
)

type keyCloakMiddleware struct {
	keycloak *client.Keycloak
}

func NewMiddleware(keycloak *client.Keycloak) *keyCloakMiddleware {
	return &keyCloakMiddleware{keycloak: keycloak}
}

func (auth *keyCloakMiddleware) ExtractBearerToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}

func (auth *keyCloakMiddleware) VerifyToken(next http.Handler) http.Handler {

	f := func(w http.ResponseWriter, r *http.Request) {
		// try to extract Authorization parameter from the HTTP header
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// extract Bearer token
		token = auth.ExtractBearerToken(token)

		if token == "" {
			http.Error(w, "Bearer Token missing", http.StatusUnauthorized)
			return
		}

		result, err := auth.keycloak.Gocloak.RetrospectToken(context.Background(), token, auth.keycloak.ClientId, auth.keycloak.ClientSecret, auth.keycloak.Realm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		// jwt, _, err := auth.keycloak.Gocloak.DecodeAccessToken(context.Background(), token, auth.keycloak.Realm)
		// if err != nil {
		// 	http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
		// 	return
		// }

		// jwtj, _ := json.Marshal(jwt)
		// fmt.Printf("token: %v\n", string(jwtj))

		if !*result.Active {
			http.Error(w, "Invalid or expired Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}
