package main

import (
	"fmt"
	"keycloak-go/client"
	"keycloak-go/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	s := http.NewServer("localhost", "8081", client.NewKeycloak())
	fmt.Println("server go")
	s.Listen()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

}