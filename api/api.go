package api

import (
	"fmt"
	"github.com/benlaplanche/kafka-example-app/helpers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", RootPathHandler).Methods("GET")

	return router
}

func RouterHandler(router *mux.Router) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		router.ServeHTTP(res, req)
	}
}

func RootPathHandler(rw http.ResponseWriter, r *http.Request) {
	vcap_services := os.Getenv("VCAP_SERVICES")

	if vcap_services == "" {
		msg := "Unable to find VCAP_SERVICES"
		http.Error(rw, msg, http.StatusInternalServerError)
		return
	}

	kafka, _, err := services_json.Parse(string(vcap_services[:]))

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(kafka.Nodes) == 0 {
		msg := "Unable to find the Kafka credentials in VCAP_SERVICES"
		http.Error(rw, msg, http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintln(rw, "hello world")
		return
	}

	return
}
