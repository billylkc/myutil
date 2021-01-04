package myutil

import (
	"encoding/json"
	"fmt"
)

// Hello to return Hello
func Hello() {
	fmt.Println("Hello")
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
