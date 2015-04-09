package kafka_example_app_test

import (
	kafka "github.com/benlaplanche/kafka-example-app"

	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Http", func() {

	var server *httptest.Server
	var reader io.Reader

	BeforeEach(func() {

		server = httptest.NewServer(kafka.Handlers())
	})

	Describe("/ tests", func() {
		rootUrl := fmt.Sprintf("%s/", server.URL)
		request, err := http.NewRequest("GET", rootUrl, reader)

		response, err := http.DefaultClient.Do(request)

		Expect(response).To(Equal("hello world"))
		Expect(err).To(Equal(nil))
	})

})
