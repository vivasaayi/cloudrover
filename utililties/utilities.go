package utililties

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func GetStringEnvVar(key string, defaultValue string, required bool) string {
	value := os.Getenv(key)

	if required && value == "" {
		panic(fmt.Sprintf("Expected value for env var %s. But got empty.", key))
	}

	if len(value) == 0 {
		return defaultValue
	}

	return value
}

func GetIntEnvVar(key string, defaultValue int, required bool) int {
	value := os.Getenv(key)

	if required && value == "" {
		panic(fmt.Sprintf("Expected int value for env var %s. But got empty.", key))
	}

	if len(value) == 0 {
		return defaultValue
	}

	parsedVal, err := strconv.Atoi(value)

	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("Unable to parse env var %s as int.", key))
	}

	return parsedVal
}

func ConvertObjectToJson(data interface{}) (string, error) {
	jsonBody, err := json.MarshalIndent(data, "", "  ")

	jsonStr := string(jsonBody)

	if err != nil {
		return "", err
	}

	return jsonStr, nil
}
