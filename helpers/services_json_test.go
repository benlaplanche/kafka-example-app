package services_json_test

import (
	"github.com/benlaplanche/kafka-example-app/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var _ = Describe("VCAP_SERVICES", func() {

	Describe("All tests", func() {

		It("cannot find the Kafka service", func() {
			file, _ := ioutil.ReadFile("../assets/vcap_services_false.json")
			kafka, zookeeper, err := services_json.Parse(string(file[:]))

			Expect(kafka.Port).To(Equal(0))
			Expect(len(kafka.Nodes)).To(Equal(0))

			Expect(zookeeper.Port).To(Equal(0))
			Expect(len(zookeeper.Nodes)).To(Equal(0))
			Expect(err).Should(HaveOccurred())
		})

		It("has the correct Kafka service properties", func() {
			file, _ := ioutil.ReadFile("../assets/vcap_services.json")
			kafka, zookeeper, err := services_json.Parse(string(file[:]))

			Expect(err).ShouldNot(HaveOccurred())

			Expect(kafka.Port).To(Equal(9092))
			Expect(kafka.Nodes).To(Equal([]string{"10.244.9.2", "10.244.9.6", "10.244.9.10", "10.244.9.14", "10.244.9.18"}))

			Expect(zookeeper.Port).To(Equal(2181))
			Expect(zookeeper.Nodes).To(Equal([]string{"10.244.9.22", "10.244.9.26", "10.244.9.30"}))

		})

	})

})
