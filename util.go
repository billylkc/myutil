package myutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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

// Shell calls the shell command
func Shell(command string) (error, string, string) {
	const ShellToUse = "bash"
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
