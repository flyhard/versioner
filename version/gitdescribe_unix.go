package version

import (
	"log"
	"os/exec"
	"strings"
)

// +build linux darwin
func GetVersion() string {
	cmd := exec.Command("git", "describe", "--tags", "--dirty")
	cmd.Dir = "."

	bytes, err := cmd.Output()
	if err != nil {
		log.Fatal("Failed to execute: ", string(bytes), err)
		return "Not specified"
	}
	version := strings.Trim(string(bytes), "\n ")
	return version
}
func GetRevision() string {
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	cmd.Dir = "."

	bytes, err := cmd.Output()
	if err != nil {
		log.Fatal("Failed to execute: ", string(bytes), err)
		return "Not specified"
	}
	version := strings.Trim(string(bytes), "\n ")
	return version
}
