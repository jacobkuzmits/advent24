package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jacobkuzmits/advent24/utils"
)

var splitMemo = make(map[int][]int)

func splitStone(stone int) (newStones []int) {
	stoneString := strconv.Itoa(stone)
	if cache, exists := splitMemo[stone]; exists {
		return cache
	}

	if stone == 0 {
		newStones = append(newStones, 1)
	} else if len(stoneString)%2 == 0 {
		firstHalf := stoneString[:len(stoneString)/2]
		secondHalf := stoneString[len(stoneString)/2:]
		secondHalfInt, _ := strconv.Atoi(secondHalf)
		secondHalfStripped := strconv.Itoa(secondHalfInt)
		firstStone, _ := strconv.Atoi(firstHalf)
		secondStone, _ := strconv.Atoi(secondHalfStripped)
		newStones = append(newStones, firstStone, secondStone)
	} else {
		// stoneInt, _ := strconv.Atoi(stone)
		newStone := stone * 2024
		newStones = append(newStones, newStone)
	}

	splitMemo[stone] = newStones
	return newStones
}

func blink(stones []int) []int {
	newStones := make([]int, 0, len(stones)*2)
	for _, stone := range stones {
		result := splitStone(stone)
		newStones = append(newStones, result...)
	}
	return newStones
}

func blinkDict(dict map[int]int) map[int]int {
	newDict := map[int]int{}
	for stone, count := range dict {
		newStones := splitStone(stone)
		for _, newStone := range newStones {

			newDict[newStone] += count
		}
	}
	return newDict
}

func partOne(filePath string) {
	// get lines as a slice of strings
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}
	stoneStrings := strings.Split(lines[0], " ")
	stones := []int{}
	for _, stone := range stoneStrings {
		stoneInt, _ := strconv.Atoi(stone)
		stones = append(stones, stoneInt)
	}
	iterations := 25
	for i := 0; i < iterations; i++ {
		stones = blink(stones)
	}
	fmt.Println(len(stones))
}

func partTwo(filePath string) {
	// Process input using precomputed values
	lines, _ := utils.GetLines(filePath)
	stoneStrings := strings.Split(lines[0], " ")
	stones := []int{}
	for _, stone := range stoneStrings {
		stoneInt, _ := strconv.Atoi(stone)
		stones = append(stones, stoneInt)
	}

	// initialize stones
	stoneDict := map[int]int{}
	for _, stone := range stones {
		stoneDict[stone] += 1
	}

	iterations := 75
	for i := 0; i < iterations; i++ {
		stoneDict = blinkDict(stoneDict)
	}
	total := 0
	for _, count := range stoneDict {
		total += count
	}
	fmt.Println(total)
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
