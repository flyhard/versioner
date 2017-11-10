package main

import (
	"os/exec"
	"log"
	"fmt"
)

func main() {
	cmd := exec.Command("git", "describe", "--tags", "--dirty")

	bytes, err := cmd.Output()
	if err != nil {
		log.Fatal("Failed to execute: ", err)
		return
	}

	fmt.Println(string(bytes))
}
