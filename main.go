package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the day number to run, e.g., 'go run . 1'")
	}

	day := os.Args[1]
	dayDir := fmt.Sprintf("day%s", day)
	dayMain := fmt.Sprintf("%s/main.go", dayDir)

	cmd := exec.Command("go", "run", dayMain)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to run day %s: %v", day, err)
	}
}
