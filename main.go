package main

import (
	"fmt"
	kafka "github.com/benlaplanche/kafka-example-app"
	"net/http"
)

func main() {
	fmt.Println("Server starting")
	http.ListenAndServe(":3000", kafka.Handlers())
}
