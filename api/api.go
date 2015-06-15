package api

import (
	"fmt"
	"github.com/benlaplanche/kafka-example-app/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

type kafka_example_app struct {
	Config services_json.Result
}

func (app *kafka_example_app) Router(config services_json.Result) *mux.Router {
	router := mux.NewRouter()
	app.Config = config

	router.HandleFunc("/", app.RootPathHandler).Methods("GET")

	return router
}

func (app *kafka_example_app) RouterHandler(router *mux.Router) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		router.ServeHTTP(res, req)
	}
}

func (app *kafka_example_app) RootPathHandler(rw http.ResponseWriter, r *http.Request) {
	// vcap_services := os.Getenv("VCAP_SERVICES")

	// if vcap_services == "" {
	// 	msg := "Unable to find VCAP_SERVICES"
	// 	http.Error(rw, msg, http.StatusInternalServerError)
	// 	return
	// }

	// kafka, _, err := services_json.Parse(string(vcap_services[:]))

	// if err != nil {
	// 	http.Error(rw, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	if len(app.Config.Nodes) == 0 {
		msg := "Unable to find the Kafka credentials in VCAP_SERVICES"
		http.Error(rw, msg, http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintln(rw, "Hello World!")
		return
	}

	return
}
