package main

import (
	"fmt"
	webserver "github.com/benlaplanche/kafka-example-app/webserver"
	"net/http"
)

func main() {
	fmt.Println("Server starting")
	http.ListenAndServe(":3000", webserver.Handlers())
}
