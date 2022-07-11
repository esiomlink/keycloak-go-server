package http

import (
	"fmt"
	"keycloak-go/client"
	"keycloak-go/controller"
	"keycloak-go/middleware"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

type httpServer struct {
	server       *http.Server
}

func NewServer(host, port string, keycloak *client.Keycloak) *httpServer {

	// create a root router
	router := mux.NewRouter()

	// add a subrouter based on matcher func
	// note, routers are processed one by one in order, so that if one of the routing matches other won't be processed
	noAuthRouter := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return r.Header.Get("Authorization") == ""
	}).Subrouter()

	// add one more subrouter for the authenticated service methods
	authRouter := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return true
	}).Subrouter()

	// instantiate a new controller which is supposed to serve our routes
	controller := controller.NewController(keycloak)

	// map url routes to controller's methods
	noAuthRouter.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		controller.Login(writer, request)
	}).Methods("POST")

	authRouter.HandleFunc("/users", func(writer http.ResponseWriter, request *http.Request) {
		controller.GetUsers(writer, request)
	}).Methods("GET")

	authRouter.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
	 	controller.GetUser(writer, request)
	 }).Methods("GET")

	 //apply middleware
	mdw := middleware.NewMiddleware(keycloak)
	authRouter.Use(mdw.VerifyToken)

	// create a server object
	s := &httpServer{
		server: &http.Server{
			Addr:         fmt.Sprintf("%s:%s", host, port),
			Handler:      router,
			WriteTimeout: time.Hour,
			ReadTimeout:  time.Hour,
		},
	}

	return s
}

func (s *httpServer) Listen() error {
	return s.server.ListenAndServe()
}
