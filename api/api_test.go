package api_test

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

	BeforeEach(func() {
		// os.Unsetenv("VCAP_SERVICES")
		os.Setenv("VCAP_SERVICES", "")
	})

	Describe("/ root path tests", func() {

		makeRequest := func() *httptest.ResponseRecorder {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/", nil)

			api.RootPathHandler(recorder, request)

			return recorder
		}

		setEnv := func() {
			file, _ := ioutil.ReadFile(("../assets/vcap_services.json"))
			os.Setenv("VCAP_SERVICES", string(file[:]))
			// fmt.Println(string(os.Getenv("VCAP_SERVICES"))[:])
		}

		It("is successful", func() {
			setEnv()
			response := makeRequest()

			Expect(response.Body.String()).To(ContainSubstring("Hello World!"))
			Expect(response.Code).To(Equal(200))
		})

		It("is unable to find VCAP_SERVICES credentials", func() {
			response := makeRequest()

			Expect(response.Body.String()).To(ContainSubstring("Unable to find VCAP_SERVICES"))
			Expect(response.Code).To(Equal(500))
		})

	})

})
