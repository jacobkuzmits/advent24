package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jacobkuzmits/advent24/utils"
)

func partOne(filePath string) {
	lines, _ := utils.GetLines(filePath)

	safe := 0

	for _, line := range lines {
		report := convertStrToInts(line)
		if isSafeAsc(report) || isSafeDesc(report) {
			safe++
		}
	}

	fmt.Println(safe)

}

func partTwo(filePath string) {
	lines, _ := utils.GetLines(filePath)

	safe := 0

outer:
	for _, line := range lines {
		nums := convertStrToInts(line)
		for i := range len(nums) {
			temp := slices.Clone(nums)
			report := append(temp[:i], temp[i+1:]...)
			if isSafeAsc(report) || isSafeDesc(report) {
				safe++
				continue outer
			}
		}
	}
	fmt.Println(safe)
}

func isSafeAsc(nums []int) bool {
	for i := range len(nums) - 1 {
		if !(nums[i+1]-nums[i] > 0 && nums[i+1]-nums[i] <= 3) {
			return false
		}
	}
	return true
}

func isSafeDesc(nums []int) bool {
	slices.Reverse(nums)
	return isSafeAsc(nums)
}

func convertStrToInts(s string) []int {
	result := []int{}

	intStrings := strings.Fields(s)
	for _, str := range intStrings {
		num, _ := strconv.Atoi(str)
		result = append(result, num)
	}
	return result
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
