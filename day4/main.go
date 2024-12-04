package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/jacobkuzmits/advent24/utils"
)

type vector struct {
	y int
	x int
}

var vectors = []vector{
	{y: 0, x: 1},   // right
	{y: 0, x: -1},  // left
	{y: 1, x: 0},   // down
	{y: -1, x: 0},  // up
	{y: 1, x: 1},   // down-right
	{y: 1, x: -1},  // down-left
	{y: -1, x: 1},  // up-right
	{y: -1, x: -1}, // up-left
}

func multiplyVector(y, x, scalar int) vector {
	return vector{
		y: y * scalar,
		x: x * scalar,
	}
}

func findXmas(lines []string, y int, x int) int {
	initialCharacter := lines[y][x]
	xmasFound := 0

	// check if starting character is x
	if strings.ToLower(string(initialCharacter)) == "x" {
		// fmt.Printf("found X at {%d,%d}\n", y, x)
		// scan all directions
		for _, direction := range vectors {
			// look for valid end position
			// continue if end position doesn't exist or isn't an "S"
			checkPos := multiplyVector(direction.y, direction.x, 3)
			targetY := y + checkPos.y
			targetX := x + checkPos.x
			if bytes, exists := safeGetChar(lines, targetY, targetX); exists {
				char := strings.ToLower(string(bytes))
				if char != "s" {
					continue
				}
				// fmt.Printf("found S at {%d,%d}\n", targetY, targetX)
				// character is an S. now we need to find the other characters
				// look for M. if it's not found, go to the next scan direction
				secondCharacterVector := multiplyVector(direction.y, direction.x, 1)
				secondCharacterY := y + secondCharacterVector.y
				secondCharacterX := x + secondCharacterVector.x
				secondCharacterBytes := lines[secondCharacterY][secondCharacterX]
				secondCharacter := strings.ToLower(string(secondCharacterBytes))
				if secondCharacter != "m" {
					continue
				}
				// fmt.Printf("found M at {%d,%d}\n", secondCharacterY, secondCharacterX)
				// M was found
				// look for A. if it's not found, go to the next scan direction
				thirdCharacterVector := multiplyVector(direction.y, direction.x, 2)
				thirdCharacterY := y + thirdCharacterVector.y
				thirdCharacterX := x + thirdCharacterVector.x
				thirdCharacterBytes := lines[thirdCharacterY][thirdCharacterX]
				thirdCharacter := strings.ToLower(string(thirdCharacterBytes))
				if thirdCharacter != "a" {
					continue
				}
				// fmt.Printf("found A at {%d,%d}\n", thirdCharacterY, thirdCharacterX)
				// all characters were found!
				xmasFound = xmasFound + 1

			}
		}

	}
	return xmasFound
}

func findMasX(lines []string, y int, x int) int {
	masXFound := 0
	initialBytes := lines[y][x]
	initialCharacter := strings.ToLower(string(initialBytes))
	if initialCharacter != "a" {
		return masXFound
	}
	surroundingChars := getSurroundingChars(lines, y, x)
	charString := strings.Join(surroundingChars, "")
	if len(charString) == 4 {
		validStrings := []string{
			"mmss",
			"msms",
			"ssmm",
			"smsm",
		}
		if slices.Contains(validStrings, charString) {
			masXFound = masXFound + 1
		}
	}

	return masXFound
}

func getCharFromByte(b byte) string {
	return strings.ToLower(string(b))
}

func getSurroundingChars(lines []string, y int, x int) (result []string) {
	var dirs = []vector{
		{y: -1, x: -1},
		{y: -1, x: 1},
		{y: 1, x: -1},
		{y: 1, x: 1},
	}
	for _, dir := range dirs {
		if b, exists := safeGetChar(lines, y+dir.y, x+dir.x); exists {
			char := getCharFromByte(b)
			result = append(result, char)
		}
	}
	return result
}

func safeGetChar(lines []string, y int, x int) (b byte, exists bool) {
	// Check if row exists
	if y < 0 || y >= len(lines) {
		return 0, false
	}

	// Check if column exists in that row
	if x < 0 || x >= len(lines[y]) {
		return 0, false
	}

	return lines[y][x], true
}

func partOne(filePath string) {
	// get lines as a slice of strings
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}
	totalXmas := 0
	for i, row := range lines {
		for j := range row {
			xmasFound := findXmas(lines, i, j)
			totalXmas = totalXmas + xmasFound
		}
	}
	fmt.Println(totalXmas)

}

func partTwo(filePath string) {
	// get lines as a slice of strings
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}

	totalMasX := 0
	for i, row := range lines {
		for j := range row {
			masXFound := findMasX(lines, i, j)
			totalMasX = totalMasX + masXFound
		}
	}
	fmt.Println(totalMasX)

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
