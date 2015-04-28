package main_test

import (
	api "github.com/benlaplanche/kafka-example-app/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
)

var _ = Describe("API", func() {

	Describe("/ root path tests", func() {

		makeRequest := func() *httptest.ResponseRecorder {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/", nil)

			api.RootPathHandler(recorder, request)

			return recorder
		}

		// setEnv := func() {
		// 	file, _ := ioutil.ReadFile("assets/vcap_services.json")
		// 	os.Setenv("VCAP_SERVICES", string(file[:]))
		// 	// val, _ := json.Marshal(file)
		// 	// os.Setenv("VCAP_SERVICES", val)

		// 	file, _ = ioutil.ReadFile("assets/vcap_application.json")
		// 	os.Setenv("VCAP_APPLICATION", string(file[:]))
		// 	// val, _ = json.Marshal(file)
		// 	// os.Setenv("VCAP_APPLICATION", val)
		// }

		validEnv := []string{
			`VCAP_APPLICATION={"instance_id":"451f045fd16427bb99c895a2649b7b2a","instance_index":0,"host":"0.0.0.0","port":61857,"started_at":"2013-08-12 00:05:29 +0000","started_at_timestamp":1376265929,"start":"2013-08-12 00:05:29 +0000","state_timestamp":1376265929,"limits":{"mem":512,"disk":1024,"fds":16384},"application_version":"c1063c1c-40b9-434e-a797-db240b587d32","application_name":"styx-james","application_uris":["styx-james.a1-app.cf-app.com"],"version":"c1063c1c-40b9-434e-a797-db240b587d32","name":"styx-james","uris":["styx-james.a1-app.cf-app.com"],"users":null}`,
			`HOME=/home/vcap/app`,
			`MEMORY_LIMIT=512m`,
			`PWD=/home/vcap`,
			`TMPDIR=/home/vcap/tmp`,
			`USER=vcap`,
			`VCAP_SERVICES={"p-kafka":[{"credentials":{"kafka":{"node_ips":["10.244.9.2","10.244.9.6","10.244.9.10","10.244.9.14","10.244.9.18"],"port":9092},"zookeeper":{"node_ips":["10.244.9.22","10.244.9.26","10.244.9.30"],"port":2181}},"label":"p-kafka","name":"kafka","plan":"shared","tags":["pivotal","kafka"]}]}`,
		}

		setEnv := func() {
			file, _ = ioutil.ReadFile(("assets/vcap_services.json"))
		}

		It("is successful", func() {
			setEnv()
			response := makeRequest()

			Expect(response.Body.String()).To(ContainSubstring("hello world"))
			Expect(response.Code).To(Equal(200))
		})

		It("is unable to find the kafka VCAP_SERVICES credentials", func() {
			response := makeRequest()

			Expect(response.Body.String()).To(ContainSubstring("Unable to find the Kafka credentials in VCAP_SERVICES"))
			Expect(response.Code).To(Equal(400))
		})

	})

})
