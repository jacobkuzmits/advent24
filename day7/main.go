package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jacobkuzmits/advent24/utils"
)

func search(nums []string, target int, curIndex int, curResult int, isPart2 bool) bool {
	if curIndex >= len(nums) {
		return curResult == target
	}
	currentNum, _ := strconv.Atoi(nums[curIndex])

	if search(nums, target, curIndex+1, curResult+currentNum, isPart2) {
		return true
	}

	if search(nums, target, curIndex+1, curResult*currentNum, isPart2) {
		return true
	}

	if isPart2 {
		concatResult, _ := strconv.Atoi(strconv.Itoa(curResult) + strconv.Itoa(currentNum))
		if search(nums, target, curIndex+1, concatResult, isPart2) {
			return true
		}
	}

	return false
}

func solve(filePath string) (part1Total, part2Total int) {
	utils.StreamFile(filePath, func(line string) {
		splitString := strings.Split(line, ":")
		expected, _ := strconv.Atoi(splitString[0])
		nums := strings.Split(strings.TrimLeft(splitString[1], " "), " ")
		firstNum, _ := strconv.Atoi(nums[0])
		if search(nums, expected, 1, firstNum, false) {
			part1Total += expected
		}
		if search(nums, expected, 1, firstNum, true) {
			part2Total += expected
		}
	})
	return part1Total, part2Total
}

func main() {
	t1, t2 := solve("testInput.txt")
	p1, p2 := solve("input.txt")
	fmt.Printf("Part 1 Test Solution: %d\n", t1)
	fmt.Printf("Part 1 Actual Solution: %d\n", p1)
	fmt.Printf("Part 2 Test Solution: %d\n", t2)
	fmt.Printf("Part 2 Actual Solution: %d\n", p2)
}
