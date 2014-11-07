package hydra_worker_map_by_limit

import (
	"errors"
	"log"
	"os"
	"strconv"

	worker "github.com/innotech/hydra-worker-map-by-limit/vendors/github.com/innotech/hydra-worker-lib"
)

const (
	// Worker
	LimitAttrKey        string = "limitAttr"
	LimitValueKey       string = "limitValue"
	MapSortKey          string = "mapSort"
	ReverseMapSortValue string = "reverse"
	// Instance
	InstanceInfoKey string = "Info"
)

func main() {
	// New Worker connected to Hydra Load Balancer
	mapByLimitWorker := worker.NewWorker(os.Args)
	mapByLimitWorker.Run(MapByLimit)
}

func MapByLimit(instances []interface{}, workerArgs map[string]interface{}) (finalInstances []interface{}) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error: ", r)
		}
	}()
	finalInstances = instances

	limitAttr, err := obtainLimitAttr(workerArgs)
	if err != nil {
		log.Println(err.Error())
		return instances
	}
	limitValue, err := obtainLimitValue(workerArgs)
	if err != nil {
		log.Println(err.Error())
		return instances
	}
	mapSort, err := obtainMapSort(workerArgs)
	if err != nil {
		log.Println(err.Error())
		return instances
	}

	mappedInstances := make([][]map[string]interface{}, 2)
	for _, i := range instances {
		instance := i.(map[string]interface{})
		if interfaceValue, ok := instance[InstanceInfoKey].(map[string]interface{})[limitAttr]; ok {
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
	if mapSort == ReverseMapSortValue {
		computedInstances[0] = mappedInstances[1]
		computedInstances[1] = mappedInstances[0]
	} else {
		computedInstances[0] = mappedInstances[0]
		computedInstances[1] = mappedInstances[1]
	}

	finalInstances = computedInstances
	return
}

func obtainLimitAttr(workerArgs map[string]interface{}) (string, error) {
	if val, ok := workerArgs[LimitAttrKey]; ok && val != "" {
		return val.(string), nil
	}
	return "", errors.New("Invalid limitAttrKey")
}

func obtainLimitValue(workerArgs map[string]interface{}) (float64, error) {
	if val, ok := workerArgs[LimitValueKey]; ok && val != "" {
		return strconv.ParseFloat(val.(string), 64)
	}
	return 0, errors.New("Invalid limitValueKey")
}

func obtainMapSort(workerArgs map[string]interface{}) (string, error) {
	if val, ok := workerArgs[MapSortKey]; ok && val != "" {
		return val.(string), nil
	}
	return "", errors.New("Invalid mapSort")
}
