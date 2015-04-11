package main

import (
	"fmt"
	"github.com/benlaplanche/cf-kafka-example-app/api"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Server starting")

	api := api.New()

	http.ListenAndServe(":3000", api)
}
