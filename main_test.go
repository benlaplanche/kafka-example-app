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

		setEnv := func() {
			file, _ := ioutil.ReadFile("assets/vcap_services.json")
			os.Setenv("VCAP_SERVICES", string(file[:]))

			file, _ = ioutil.ReadFile("assets/vcap_application.json")
			os.Setenv("VCAP_APPLICATION", string(file[:]))
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
