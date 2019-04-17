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

func CheckOrArg(key string, lst []string, idx int) (string, error) {
	return CheckOr(key, func(key string) (string, error) {
		if idx+1 > len(lst) {
			return "", fmt.Errorf(
				"no %s env var set, and argument %d was missing",
				key,
				idx+1,
			)
		}
		return lst[idx], nil
	})
}
