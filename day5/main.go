package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/jacobkuzmits/advent24/utils"
)

func parseLines(lines []string) (reqs []string, orders [][]int) {

	seenBlankLine := false

	for _, line := range lines {
		if line == "" {
			seenBlankLine = true
			continue
		}

		if !seenBlankLine {
			reqs = append(reqs, line)
		}
		if seenBlankLine {
			numLine := []int{}
			split := strings.Split(line, ",")
			for _, numString := range split {
				num, _ := strconv.Atoi(numString)
				numLine = append(numLine, num)
			}
			orders = append(orders, numLine)

		}
	}
	return reqs, orders
}

func getReqsMaps(reqs []string) (map[int][]int, map[int][]int) {
	comesAfter := make(map[int][]int)
	goesBefore := make(map[int][]int)
	for _, req := range reqs {
		var first, second int
		fmt.Sscanf(req, "%d|%d", &first, &second)
		comesAfter[second] = append(comesAfter[second], first)
		goesBefore[first] = append(goesBefore[first], second)
	}

	return comesAfter, goesBefore
}

func contains(slice []int, item int) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

func partOne(filePath string) {
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}

	reqs, orders := parseLines(lines)
	comesAfter, goesBefore := getReqsMaps(reqs)
	validOrders := [][]int{}
outer:
	for _, order := range orders {
		for index, num := range order {
			prevNums := order[:index]
			nextNums := order[index+1:]
			for _, prev := range prevNums {
				if exists := contains(comesAfter[num], prev); !exists {
					continue outer
				}
			}
			for _, next := range nextNums {
				if exists := contains(goesBefore[num], next); !exists {
					continue outer
				}
			}
		}
		validOrders = append(validOrders, order)
	}
	result := 0
	for _, order := range validOrders {
		result += order[int(len(order)/2)]
	}
	fmt.Println(result)
}

func partTwo(filePath string) {
	input, _ := os.ReadFile(filePath)
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	cmp := func(a, b string) int {
		for _, s := range strings.Split(split[0], "\n") {
			if s := strings.Split(s, "|"); s[0] == a && s[1] == b {
				return -1
			}
		}
		return 0
	}

	run := func(sorted bool) (r int) {
		for _, s := range strings.Split(split[1], "\n") {
			if s := strings.Split(s, ","); slices.IsSortedFunc(s, cmp) == sorted {
				slices.SortFunc(s, cmp)
				n, _ := strconv.Atoi(s[len(s)/2])
				r += n
			}
		}
		return r
	}

	fmt.Println(run(false))
}

func main() {
	fmt.Println("\nPart 1 Test Solution")
	partOne("testInput.txt")

	fmt.Println("\nPart 1 Actual Solution")
	partOne("input.txt")

	fmt.Println("\nPart 2 Test Solution")
	partTwo("day5/testInput.txt")

	fmt.Println("\nPart 2 Actual Solution")
	partTwo("day5/input.txt")
}
