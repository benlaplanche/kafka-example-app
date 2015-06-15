package main

import (
	"fmt"
	"github.com/benlaplanche/kafka-example-app/api"
	"github.com/benlaplanche/kafka-example-app/helpers"
	"github.com/pivotal-golang/lager"
	"net/http"
	"os"
)

func main() {
	logger := lager.NewLogger("kafka-broker")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	logger.RegisterSink(lager.NewWriterSink(os.Stderr, lager.ERROR))

	fmt.Println("Server starting")

	vcap_services := fetch_vcap_services()
	kafka_config, _, err := services_json.Parse(vcap_services)

	if err != nil {
		logger.Fatal("Parsing VCAP_SERVICES", err)
	}

	api := api.RouterHandler(api.Router(kafka_config))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Port = " + port)

	logger.Fatal("http-listen", http.ListenAndServe(":"+port, api))
}

func fetch_vcap_services() string {
	vcap_services := os.Getenv("VCAP_SERVICES")

	if vcap_services == "" {
		msg := "Unable to find VCAP_SERVICES"
		panic(msg)
	}

	return vcap_services
}
