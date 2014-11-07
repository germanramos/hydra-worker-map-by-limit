package coverage_fixture_test

import (
	. "github.com/innotech/hydra-worker-map-by-limit/vendors/github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCoverageFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CoverageFixture Suite")
}
