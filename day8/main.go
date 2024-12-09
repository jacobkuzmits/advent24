package main

import (
	"fmt"
	"log"

	"github.com/jacobkuzmits/advent24/utils"
)

type coord struct {
	x int
	y int
}

type antennas map[rune][]coord

func partOne(filePath string) {
	// get lines
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}

	// initialize
	antennaMap := antennas{}
	gridWidth := len(lines[0])
	gridHeight := len(lines)
	antinodes := make([][]bool, gridHeight)
	for i := range antinodes {
		antinodes[i] = make([]bool, gridWidth)
	}

	// get hashmap of lists of antennas
	for rowIndex, line := range lines {
		for colIndex, char := range line {
			if char != '.' {
				if _, exists := antennaMap[char]; !exists {
					antennaMap[char] = []coord{}
				}
				antennaMap[char] = append(antennaMap[char], coord{x: colIndex, y: rowIndex})
			}
		}
	}

	// iterate over each character
	for _, antennas := range antennaMap {
		// iterate over each antenna
		for i := 0; i < len(antennas); i++ {
			// iterate over each permutation of antennas
			for j := 0; j < len(antennas); j++ {
				if i == j {
					continue
				}
				// get distance between antennas
				distX := antennas[i].x - antennas[j].x
				distY := antennas[i].y - antennas[j].y

				// get antinode positions
				antinode1 := coord{x: antennas[i].x + distX, y: antennas[i].y + distY}
				antinode2 := coord{x: antennas[j].x - distX, y: antennas[j].y - distY}

				// check if antinodes are in bounds
				inBounds1 := antinode1.x >= 0 && antinode1.x < gridWidth && antinode1.y >= 0 && antinode1.y < gridHeight
				inBounds2 := antinode2.x >= 0 && antinode2.x < gridWidth && antinode2.y >= 0 && antinode2.y < gridHeight

				// place valid antinodes
				if inBounds1 {
					antinodes[antinode1.y][antinode1.x] = true
				}
				if inBounds2 {
					antinodes[antinode2.y][antinode2.x] = true
				}
			}
		}
	}

	// count all antinodes
	antinodeCount := 0
	for _, row := range antinodes {
		for _, col := range row {
			if col {
				antinodeCount++
			}
		}
	}
	fmt.Println(antinodeCount)
}

func partTwo(filePath string) {
	// get lines
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}

	// initialize
	antennaMap := antennas{}
	gridWidth := len(lines[0])
	gridHeight := len(lines)
	antinodes := make([][]bool, gridHeight)
	for i := range antinodes {
		antinodes[i] = make([]bool, gridWidth)
	}

	// get hashmap of lists of antennas
	for rowIndex, line := range lines {
		for colIndex, char := range line {
			if char != '.' {
				antinodes[rowIndex][colIndex] = true
				if _, exists := antennaMap[char]; !exists {
					antennaMap[char] = []coord{}
				}
				antennaMap[char] = append(antennaMap[char], coord{x: colIndex, y: rowIndex})
			}
		}
	}

	// iterate over each character
	for _, antennas := range antennaMap {
		// iterate over each antenna
		for i := 0; i < len(antennas); i++ {
			// iterate over each permutation of antennas
			for j := 0; j < len(antennas); j++ {
				if i == j {
					continue
				}
				// get distance between antennas
				distX := antennas[i].x - antennas[j].x
				distY := antennas[i].y - antennas[j].y

				// extend antinodes in one direction
				initX := antennas[i].x
				initY := antennas[i].y
				for {
					newAntinode := coord{x: initX + distX, y: initY + distY}
					inBounds := newAntinode.x >= 0 && newAntinode.x < gridWidth && newAntinode.y >= 0 && newAntinode.y < gridHeight
					if !inBounds {
						break
					}
					antinodes[newAntinode.y][newAntinode.x] = true
					initX = newAntinode.x
					initY = newAntinode.y
				}

				// extend the other direction
				initX = antennas[j].x
				initY = antennas[j].y
				for {
					newAntinode := coord{x: initX - distX, y: initY - distY}
					inBounds := newAntinode.x >= 0 && newAntinode.x < gridWidth && newAntinode.y >= 0 && newAntinode.y < gridHeight
					if !inBounds {
						break
					}
					antinodes[newAntinode.y][newAntinode.x] = true
					initX = newAntinode.x
					initY = newAntinode.y
				}
			}
		}
	}

	// count antinodes
	antinodeCount := 0
	for _, row := range antinodes {
		for _, col := range row {
			if col {
				antinodeCount++
			}
		}
	}
	fmt.Println(antinodeCount)
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
