package gbytes_test

import (
	. "github.com/innotech/hydra-worker-map-by-limit/vendors/github.com/onsi/ginkgo"
	. "github.com/innotech/hydra-worker-map-by-limit/vendors/github.com/onsi/gomega"

	"testing"
)

func TestGbytes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gbytes Suite")
}
