package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jacobkuzmits/advent24/utils"
)

func createFs(s string) (fs []int64) {
	for index, char := range s {
		blockSize, _ := strconv.Atoi(string(char))
		id := -1
		if index%2 == 0 {
			id = index / 2
		}
		for i := 0; i < blockSize; i++ {
			fs = append(fs, int64(id))
		}
	}
	return fs
}

func compactFs(n []int64) []int64 {
	start := 0
	end := len(n) - 1
	for start < end+1 {
		if n[start] >= 0 {
			start += 1
			continue
		}
		if n[end] >= 0 {
			n[start] = n[end]

			end += -1
			continue
		}

		end += -1
	}
	return n[:end+1]
}

func compactFs2(n []int64) []int64 {
	end := len(n) - 1

	for end >= 0 {
		if n[end] < 0 {
			end--
			continue
		}

		fileValue := n[end]
		fileSize := 1
		for end-1 >= 0 && n[end-1] == fileValue {
			fileSize++
			end--
		}

		currentPos := 0
		moved := false
		for currentPos < end {
			if n[currentPos] >= 0 {
				currentPos++
				continue
			}

			emptyStart := currentPos
			emptySize := 0
			for currentPos+emptySize < len(n) && n[currentPos+emptySize] < 0 {
				emptySize++
			}

			if fileSize <= emptySize {
				for i := 0; i < fileSize; i++ {
					n[emptyStart+i] = fileValue
				}
				for i := 0; i < fileSize; i++ {
					n[end+i] = -1
				}
				moved = true
				break
			}
			currentPos += emptySize
		}

		if !moved {
			end--
		}
	}

	end = len(n) - 1
	for end >= 0 && n[end] < 0 {
		end--
	}
	return n[:end+1]
}

func calcChecksum(nums []int64) (r int64) {
	for i, n := range nums {
		r += int64(i) * n
	}
	return r
}

func calcChecksum2(nums []int64) (r int64) {
	for i, n := range nums {
		if n >= 0 {
			r += int64(i) * n
		}
	}
	return r
}

func partOne(filePath string) {
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}
	fs := createFs(lines[0])
	fs = compactFs(fs)
	result := calcChecksum(fs)
	fmt.Println(result)
}

func partTwo(filePath string) {
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}
	fs := createFs(lines[0])
	fs = compactFs2(fs)
	result := calcChecksum2(fs)
	fmt.Println(result)
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
