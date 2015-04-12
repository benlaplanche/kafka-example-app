package main_test

import (
	api "github.com/benlaplanche/kafka-example-app/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("API", func() {

	Describe("/ root path tests", func() {

		makeRequest := func() *httptest.ResponseRecorder {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/", nil)

			api.RootPathHandler(recorder, request)

			return recorder
		}

		It("works", func() {
			response := makeRequest()

			Expect(response.Body.String()).To(ContainSubstring("hello world"))
			Expect(response.Code).To(Equal(200))
		})

	})

})
