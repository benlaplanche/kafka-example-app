package main

import (
	"fmt"
	"github.com/benlaplanche/kafka-example-app/api"
	"net/http"
)

func main() {
	fmt.Println("Server starting")

	api := api.RouterHandler(api.Router())

	http.ListenAndServe(":3000", api)
}
