package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jacobkuzmits/advent24/utils"
)

func recursiveSearch(nums []string, target int, currentIndex int, currentResult int) bool {
	// base case
	if currentIndex >= len(nums) {
		return currentResult == target
	}
	currentNum, _ := strconv.Atoi(nums[currentIndex])

	// addition
	if recursiveSearch(nums, target, currentIndex+1, currentResult+currentNum) {
		return true
	}

	// multiplication
	if recursiveSearch(nums, target, currentIndex+1, currentResult*currentNum) {
		return true
	}

	return false
}

func recursiveSearchWithConcat(nums []string, target int, currentIndex int, currentResult int) bool {
	// base case
	if currentIndex >= len(nums) {
		return currentResult == target
	}
	currentNum, _ := strconv.Atoi(nums[currentIndex])

	// addition
	if recursiveSearchWithConcat(nums, target, currentIndex+1, currentResult+currentNum) {
		return true
	}

	// multiplication
	if recursiveSearchWithConcat(nums, target, currentIndex+1, currentResult*currentNum) {
		return true
	}

	// concatenation
	concatResult, _ := strconv.Atoi(strconv.Itoa(currentResult) + strconv.Itoa(currentNum))
	if recursiveSearchWithConcat(nums, target, currentIndex+1, concatResult) {
		return true
	}

	return false
}

func partOne(filePath string) {
	total := 0
	utils.StreamFile(filePath, func(line string) {
		splitString := strings.Split(line, ":")
		expected, _ := strconv.Atoi(splitString[0])
		nums := strings.Split(strings.TrimLeft(splitString[1], " "), " ")
		firstNum, _ := strconv.Atoi(nums[0])
		if recursiveSearch(nums, expected, 1, firstNum) {
			total += expected
		}
	})
	fmt.Println(total)
}

func partTwo(filePath string) {
	total := 0
	utils.StreamFile(filePath, func(line string) {
		splitString := strings.Split(line, ":")
		expected, _ := strconv.Atoi(splitString[0])
		nums := strings.Split(strings.TrimLeft(splitString[1], " "), " ")
		firstNum, _ := strconv.Atoi(nums[0])
		if recursiveSearchWithConcat(nums, expected, 1, firstNum) {
			total += expected
		}
	})
	fmt.Println(total)
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
