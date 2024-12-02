package main

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/jacobkuzmits/advent24/utils"
)

func partOne(filePath string) {
	// initialize two slices to hold the values from the file
	left := []int{}
	right := []int{}

	// stream the file and parse the values into the slices
	utils.StreamFile(filePath, func(line string) {
		var leftVal, rightVal int
		_, err := fmt.Sscanf(line, "%d   %d", &leftVal, &rightVal)
		if err != nil {
			log.Fatal("error reading line: ", err)
		}
		left = append(left, leftVal)
		right = append(right, rightVal)
	})

	// sort the slices for easy comparison
	sort.Ints(left)
	sort.Ints(right)

	// solve the problem!
	distance := 0
	for i, leftVal := range left {
		rightVal := right[i]
		distance += int(math.Abs(float64(leftVal - rightVal)))
	}
	fmt.Println(distance)
}

func partTwo(filePath string) {
	// initialize two slices to hold the values from the file
	left := []int{}
	right := []int{}

	// stream the file and parse the values into the slices
	utils.StreamFile(filePath, func(line string) {
		var val1, val2 int
		_, err := fmt.Sscanf(line, "%d   %d", &val1, &val2)
		if err != nil {
			log.Fatal("error reading line: ", err)
		}
		left = append(left, val1)
		right = append(right, val2)
	})

	// create a map to hold the frequencies of each value in the right slice
	frequencies := make(map[int]int)

	// count the frequencies of each value in the right slice
	for _, val := range right {
		frequencies[val] += 1
	}

	// solve the problem!
	similarityScore := 0
	for _, val := range left {
		if freq, exists := frequencies[val]; exists {
			similarityScore += freq * val
		}
	}

	fmt.Println(similarityScore)
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
