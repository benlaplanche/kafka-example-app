package main

import (
	"fmt"
	"github.com/benlaplanche/kafka-example-app/api"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Server starting")

	api := api.RouterHandler(api.Router())

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.ListenAndServe(":"+port, api)
}
