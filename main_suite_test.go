package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKafkaExampleApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kafka Example App Suite")
}
