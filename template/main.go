package main

import (
	"fmt"
	"log"

	"github.com/jacobkuzmits/advent24/utils"
)

func partOne(filePath string) {
	// get lines as a slice of strings
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}
	fmt.Println(lines[0])

	utils.StreamFile(filePath, func(line string) {
		// do something with each line
	})
}

func partTwo(filePath string) {
	// get lines as a slice of strings
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}
	fmt.Println(lines[0])

	utils.StreamFile(filePath, func(line string) {
		// do something with each line
	})
}

func main() {
	fmt.Println("\nPart 1 Test Solution")
	partOne("testInput.txt")

	fmt.Println("\nPart 1 Actual Solution")
	partOne("input.txt")

	fmt.Println("\nPart 2 Test Solution")
	partTwo("testInput.txt")

	fmt.Println("\nPart 2 Actual Solution")
	partTwo("input.txt")
}
