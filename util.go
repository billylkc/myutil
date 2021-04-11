package myutil

import (
	"encoding/json"
	"fmt"
	"os"
)

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// GetEnv gets the environment variable with the given name
func GetEnv(name string) (string, error) {
	v := os.Getenv(name)
	if v == "" {
		return "", fmt.Errorf("missing required environment variable " + name)
	}
	return v, nil
}
