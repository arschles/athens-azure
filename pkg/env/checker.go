package env

import (
	"fmt"
	"os"
)

func Check(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("Missing environment variable %s", key)
	}
	return val, nil
}
