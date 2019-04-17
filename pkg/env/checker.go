package env

import (
	"fmt"
	"os"
)

type Defaulter func(string) (string, error)

func Check(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("Missing environment variable %s", key)
	}
	return val, nil
}

func CheckOr(key string, fn Defaulter) (string, error) {
	ret, err := Check(key)
	if err == nil {
		return ret, nil
	}
	ret, err = fn(key)
	if err != nil {
		return "", err
	}
	return ret, nil
}
