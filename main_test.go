package hydra_worker_map_by_limit_test

import (
	. "github.com/innotech/hydra-worker-map-by-limit"

	. "github.com/innotech/hydra-worker-map-by-limit/vendors/github.com/onsi/ginkgo"
	. "github.com/innotech/hydra-worker-map-by-limit/vendors/github.com/onsi/gomega"
)

var _ = Describe("MapByLimit", func() {
	var (
		inputInstances []interface{}
		workerArgs     map[string]interface{}
	)

	BeforeEach(func() {
		inputInstances = []interface{}{
			map[string]interface{}{
				"Info": map[string]interface{}{
					"cpuLoad": "40",
				},
			},
			map[string]interface{}{
				"Info": map[string]interface{}{
					"cpuLoad": "95",
				},
			},
			map[string]interface{}{
				"Info": map[string]interface{}{
					"cpuLoad": "15",
				},
			},
		}
		workerArgs = map[string]interface{}{
			LimitAttrKey:  "cpuLoad",
			LimitValueKey: "90",
			MapSortKey:    "normal",
		}
	})

	Context("when doesn't exist limit attribute key", func() {
		It("should return input instances", func() {
			delete(workerArgs, LimitAttrKey)
			outputInstances := MapByLimit(inputInstances, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when invalid limit attribute key", func() {
		It("should return input instances", func() {
			workerArgs[LimitAttrKey] = ""
			outputInstances := MapByLimit(inputInstances, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when doesn't exist limit value key", func() {
		It("should return input instances", func() {
			delete(workerArgs, LimitValueKey)
			outputInstances := MapByLimit(inputInstances, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when invalid limit value key", func() {
		It("should return input instances", func() {
			workerArgs[LimitValueKey] = ""
			outputInstances := MapByLimit(inputInstances, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when doesn't exist map sort key", func() {
		It("should return input instances", func() {
			delete(workerArgs, MapSortKey)
			outputInstances := MapByLimit(inputInstances, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when invalid map sort key", func() {
		It("should return input instances", func() {
			workerArgs[MapSortKey] = ""
			outputInstances := MapByLimit(inputInstances, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when limit attribute exists in all instances", func() {
		Context("when map sort attribute is equal to reverse", func() {
			It("should map instances correctly", func() {
				workerArgs[MapSortKey] = "reverse"
				outputInstances := MapByLimit(inputInstances, workerArgs)
				expectedInstances := []interface{}{
					[]map[string]interface{}{
						map[string]interface{}{
							"Info": map[string]interface{}{
								"cpuLoad": "95",
							},
						},
					},
					[]map[string]interface{}{
						map[string]interface{}{
							"Info": map[string]interface{}{
								"cpuLoad": "40",
							},
						},
						map[string]interface{}{
							"Info": map[string]interface{}{
								"cpuLoad": "15",
							},
						},
					},
				}
				Expect(outputInstances).To(HaveLen(2), "Must contain two elements")
				Expect(outputInstances).To(Equal(expectedInstances), "Must return the expected slice of instances")
			})
		})
		Context("when map sort attribute is not equal to reverse", func() {
			It("should map instances correctly", func() {
				workerArgs[MapSortKey] = "direct"
				outputInstances := MapByLimit(inputInstances, workerArgs)
				expectedInstances := []interface{}{
					[]map[string]interface{}{
						map[string]interface{}{
							"Info": map[string]interface{}{
								"cpuLoad": "40",
							},
						},
						map[string]interface{}{
							"Info": map[string]interface{}{
								"cpuLoad": "15",
							},
						},
					},
					[]map[string]interface{}{
						map[string]interface{}{
							"Info": map[string]interface{}{
								"cpuLoad": "95",
							},
						},
					},
				}
				Expect(outputInstances).To(HaveLen(2), "Must contain two elements")
				Expect(outputInstances).To(Equal(expectedInstances), "Must return the expected slice of instances")
			})
		})
	})
	Context("when limit attribute doesn't exist in some instances", func() {
		It("should map instances correctly", func() {
			delete(inputInstances[0].(map[string]interface{})["Info"].(map[string]interface{}), "cpuLoad")
			workerArgs[MapSortKey] = "direct"
			outputInstances := MapByLimit(inputInstances, workerArgs)
			expectedInstances := []interface{}{
				[]map[string]interface{}{
					map[string]interface{}{
						"Info": map[string]interface{}{
							"cpuLoad": "15",
						},
					},
				},
				[]map[string]interface{}{
					map[string]interface{}{
						"Info": map[string]interface{}{},
					},
					map[string]interface{}{
						"Info": map[string]interface{}{
							"cpuLoad": "95",
						},
					},
				},
			}
			Expect(outputInstances).To(HaveLen(2), "Must contain two elements")
			Expect(outputInstances).To(Equal(expectedInstances), "Must return the expected slice of instances")
		})
	})
})
