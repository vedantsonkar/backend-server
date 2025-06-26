package utils

import (
	"os"
	"strconv"
)

func GetenvOrDefault(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func GetenvInt(key string, fallback int) int {
	val := os.Getenv(key)
	if i, err := strconv.Atoi(val); err == nil {
		return i
	}
	return fallback
}

func GetenvFloat(key string, fallback float64) float64 {
	val := os.Getenv(key)
	if f, err := strconv.ParseFloat(val, 64); err == nil {
		return f
	}
	return fallback
}
