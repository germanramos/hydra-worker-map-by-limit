package failing_before_suite_test

import (
	. "github.com/innotech/hydra-worker-pilot-client/vendors/github.com/onsi/ginkgo"
)

var _ = Describe("FailingBeforeSuite", func() {
	It("should never run", func() {
		println("NEVER SEE THIS")
	})

	It("should never run", func() {
		println("NEVER SEE THIS")
	})
})
