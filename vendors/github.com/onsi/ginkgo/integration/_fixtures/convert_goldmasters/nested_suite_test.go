package nested_test

import (
	. "github.com/innotech/hydra-worker-map-by-limit/vendors/github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNested(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nested Suite")
}