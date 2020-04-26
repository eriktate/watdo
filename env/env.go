package env

import (
	"os"
	"strconv"
)

func GetString(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return def
}

func GetInt(key string, def int) int {
	val, err := strconv.ParseInt(os.Getenv(key), 10, 32)
	if err != nil {
		return def
	}

	return int(val)
}

func GetUint(key string, def uint) uint {
	val, err := strconv.ParseUint(os.Getenv(key), 10, 32)
	if err != nil {
		return def
	}

	return uint(val)
}
