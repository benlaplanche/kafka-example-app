package api

import (
	"fmt"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/gorilla/mux"
	"net/http"
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
	appEnv, err := cfenv.Current()
	if err != nil {
		fmt.Fprintln(rw, err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	services := appEnv.Services
	fmt.Println("printing services")
	fmt.Println(services)
	_, error := services.WithName("p-kafka")

	if error != nil {
		fmt.Fprintln(rw, "Unable to find the Kafka credentials in VCAP_SERVICES")
		rw.WriteHeader(http.StatusBadRequest)
		return
	} else {
		fmt.Fprintln(rw, "hello world")
		rw.WriteHeader(http.StatusOK)
		return
	}

	return
}
