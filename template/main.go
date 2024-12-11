package main

import (
	"fmt"
	"log"
	"time"

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
	start := time.Now()
	partOne("testInput.txt")
	fmt.Printf("Test execution time: %v\n", time.Since(start))

	fmt.Println("\nPart 1 Actual Solution")
	start = time.Now()
	partOne("input.txt")
	fmt.Printf("Actual execution time: %v\n", time.Since(start))

	fmt.Println("\nPart 2 Test Solution")
	start = time.Now()
	partTwo("testInput.txt")
	fmt.Printf("Test execution time: %v\n", time.Since(start))

	fmt.Println("\nPart 2 Actual Solution")
	start = time.Now()
	partTwo("input.txt")
	fmt.Printf("Actual execution time: %v\n", time.Since(start))
}
