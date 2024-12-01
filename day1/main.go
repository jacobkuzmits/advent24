package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func StreamFile(filePath string, callback func(string)) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		callback(sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
}

func partOne() {
	left := []int{}
	right := []int{}
	StreamFile("input.txt", func(line string) {
		var val1, val2 int
		_, err := fmt.Sscanf(line, "%d   %d", &val1, &val2)
		if err != nil {
			log.Fatal("error reading line: ", err)
		}
		left = append(left, val1)
		right = append(right, val2)
	})

	sort.Ints(left)
	sort.Ints(right)
	distance := 0
	for i, leftVal := range left {
		rightVal := right[i]
		distance += int(math.Abs(float64(leftVal) - float64(rightVal)))
	}
	fmt.Println("Part 1 answer: ", distance)
}

func partTwo() {
	left := []int{}
	right := []int{}
	StreamFile("input.txt", func(line string) {
		var val1, val2 int
		_, err := fmt.Sscanf(line, "%d   %d", &val1, &val2)
		if err != nil {
			log.Fatal("error reading line: ", err)
		}
		left = append(left, val1)
		right = append(right, val2)
	})

	occurences := make(map[int]int)

	for _, val := range right {
		occurences[val] += 1
	}

	similarity := 0

	for _, val := range left {
		occ, exists := occurences[val]
		if exists != true {
			continue
		}
		similarity += occ * val
	}

	fmt.Println("Part 2 answer: ", similarity)

}

func main() {
	partOne()
	partTwo()
}
