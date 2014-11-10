package main_test

import (
	. "github.com/innotech/hydra-worker-map-by-limit/vendors/github.com/onsi/ginkgo"
	. "github.com/innotech/hydra-worker-map-by-limit/vendors/github.com/onsi/gomega"

	"testing"
)

func TestHydraWorkerMapByLimit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HydraWorkerMapByLimit Suite")
}
