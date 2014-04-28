package main

import (
	"encoding/json"
	"os"

	worker "github.com/innotech/hydra-worker-pong/vendors/github.com/innotech/hydra-worker-lib"
)

func main() {
	if len(os.Args) < 3 {
		panic("Invalid number of arguments, you need to add at least the arguments for the server address and the service name")
	}
	serverAddr := os.Args[1]  // e.g. "tcp://localhost:5555"
	serviceName := os.Args[2] // e.g. map-by-limit
	verbose := len(os.Args) >= 4 && os.Args[3] == "-v"

	// New Worker connected to Hydra Load Balancer
	mapByLimitWorker := worker.NewWorker(serverAddr, serviceName, verbose)
	fn := func(instances []map[string]interface{}, args map[string]string) []interface{} {
		limitAttr := args["limitAttr"]
		limitValue := args["limitValue"]
		mapSort := args["mapSort"]

		mappedInstances := make([]map[string]interface{}, 2)
		for _, instance := range instances {
			if val, ok := mappedInstances[instance[limitAttr]]; ok {
				if val < limitValue {
					mappedInstances[0] = append(mappedInstances[0], instance)
				} else {
					mappedInstances[1] = append(mappedInstances[1], instance)
				}
			} else {
				// TODO:
			}
		}

		computedInstances := make([]interface{}, 2)
		if mapSort == "reverse" {
			computedInstances[0] = mappedInstances[1]
			computedInstances[1] = mappedInstances[0]
		} else {
			computedInstances = mappedInstances
		}
		return computedInstances
	}
	mapByLimitWorker.Run(fn)
}
