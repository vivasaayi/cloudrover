package utililties

import (
	"fmt"
	"os"
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
