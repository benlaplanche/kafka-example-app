package kafka_example_app_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKafkaExampleApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KafkaExampleApp Suite")
}
