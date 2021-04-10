package myutil

import (
	"encoding/json"
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
