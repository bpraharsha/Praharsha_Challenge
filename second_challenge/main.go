package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func traverse(input map[string]interface{}) map[string]interface{} {
	transformed := make(map[string]interface{})

	for key, value := range input {
		key = strings.TrimSpace(key)
		if key == "" {
			continue
		}

		switch val := value.(type) {
		case map[string]interface{}:
			transformed[key] = traverseMap(val)
		case string:
			transformed[key] = traverseString(val)
		case float64:
			transformed[key] = val
		case bool:
			transformed[key] = val
		case nil:
			continue
		}
	}

	return transformed
}

func traverseMap(input map[string]interface{}) map[string]interface{} {
	transformed := make(map[string]interface{})

	for key, value := range input {
		key = strings.TrimSpace(key)
		if key == "" {
			continue
		}

		// Transform nested maps recursively
		switch val := value.(type) {
		case map[string]interface{}:
			transformed[key] = traverseMap(val)
		case []interface{}:
			transformed[key] = traverseList(val)
		case string:
			transformed[key] = traverseString(val)
		case float64:
			transformed[key] = val
		case bool:
			transformed[key] = val
		case nil:
			continue
		}
	}

	return transformed
}

func traverseList(input []interface{}) []interface{} {
	var transformed []interface{}

	for _, item := range input {
		switch val := item.(type) {
		case map[string]interface{}:
			transformed = append(transformed, traverseMap(val))
		case string:
			transformed = append(transformed, traverseString(val))
		case float64:
			transformed = append(transformed, val)
		case bool:
			transformed = append(transformed, val)
		case nil:
			continue
		}
	}

	return transformed
}

func traverseString(input string) interface{} {
	// interpret the string as a time
	t, err := time.Parse(time.RFC3339, input)
	if err == nil {
		// Upon successful interpretation, return the Unix Epoch time
		return t.Unix()
	}

	// interpret the string as a number
	num, err := strconv.ParseFloat(input, 64)
	if err == nil {
		return num
	}

	// interpret the string as a boolean value
	switch input {
	case "1", "t", "T", "TRUE", "true", "True":
		return true
	case "0", "f", "F", "FALSE", "false", "False":
		return false
	default:
		//If no transformation is applied, return the original string
		return strings.TrimSpace(input)
	}
}

func main() {
	// Read input JSON file
	file := "input.json"
	inputData, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading input JSON file:", err)
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal(inputData, &data); err != nil {
		fmt.Println("Error parsing input JSON:", err)
		return
	}
	transformed := traverse(data)

	transformedJSON, err := json.MarshalIndent(transformed, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling transformed JSON:", err)
		return
	}
	fmt.Println(string(transformedJSON))
}
