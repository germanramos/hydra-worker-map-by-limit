package main

import (
	"os"
	"strconv"

	worker "github.com/innotech/hydra-worker-map-by-limit/vendors/github.com/innotech/hydra-worker-lib"
)

func main() {
	// New Worker connected to Hydra Load Balancer
	mapByLimitWorker := worker.NewWorker(os.Args)
	fn := func(instances []interface{}, args map[string]interface{}) []interface{} {
		limitAttr := args["limitAttr"].(string)
		limitValue, _ := strconv.ParseFloat(args["limitValue"].(string), 64)
		mapSort := args["mapSort"].(string)

		mappedInstances := make([][]map[string]interface{}, 2)
		for _, i := range instances {
			instance := i.(map[string]interface{})
			if interfaceValue, ok := instance["Info"].(map[string]interface{})[limitAttr]; ok {
				value, _ := strconv.ParseFloat(interfaceValue.(string), 64)
				if value < limitValue {
					mappedInstances[0] = append(mappedInstances[0], instance)
				} else {
					mappedInstances[1] = append(mappedInstances[1], instance)
				}
			} else {
				mappedInstances[1] = append(mappedInstances[1], instance)
			}
		}

		computedInstances := make([]interface{}, 2)
		if mapSort == "reverse" {
			computedInstances[0] = mappedInstances[1]
			computedInstances[1] = mappedInstances[0]
		} else {
			computedInstances[0] = mappedInstances[0]
			computedInstances[1] = mappedInstances[1]
		}

		return computedInstances
	}
	mapByLimitWorker.Run(fn)
}
