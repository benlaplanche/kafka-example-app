package kafka_example_app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Handlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", RootHandler).Methods("GET")

	return r
}

func RootHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "hello world")
	rw.WriteHeader(http.StatusOK)
}
