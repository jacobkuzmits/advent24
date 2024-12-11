package main

import (
	"fmt"
	"log"

	"github.com/jacobkuzmits/advent24/utils"
)

type coord struct {
	y, x int
}

func gridToInts(grid []string) (newGrid [][]int) {
	for _, row := range grid {
		newRow := []int{}
		for _, val := range row {
			newRow = append(newRow, int(val)-'0')
		}
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}

func getTrailheads(grid [][]int) (t []coord) {
	for rowIndex, row := range grid {
		for colIndex, height := range row {
			if height == 0 {
				t = append(t, coord{y: rowIndex, x: colIndex})
			}
		}
	}
	return t
}

func walk(grid [][]int, curPos coord, prevHeight int, visited map[coord]bool, peaks map[coord]bool) int {
	if curPos.y < 0 || curPos.y >= len(grid) || curPos.x < 0 || curPos.x >= len(grid[0]) {
		return 0
	}

	currentHeight := grid[curPos.y][curPos.x]

	if currentHeight != prevHeight+1 && prevHeight != -1 {
		return 0
	}

	if visited[curPos] {
		return 0
	}

	if currentHeight == 9 {
		if peaks[curPos] {
			return 0
		}
		peaks[curPos] = true
		return 1
	}

	visited[curPos] = true

	totalPaths := 0
	directions := []coord{
		{y: curPos.y - 1, x: curPos.x},
		{y: curPos.y + 1, x: curPos.x},
		{y: curPos.y, x: curPos.x - 1},
		{y: curPos.y, x: curPos.x + 1},
	}

	for _, dir := range directions {
		totalPaths += walk(grid, dir, currentHeight, visited, peaks)
	}

	visited[curPos] = false

	return totalPaths
}

func walk2(grid [][]int, curPos coord, prevHeight int, visited map[coord]bool) int {
	if curPos.y < 0 || curPos.y >= len(grid) || curPos.x < 0 || curPos.x >= len(grid[0]) {
		return 0
	}

	currentHeight := grid[curPos.y][curPos.x]

	if currentHeight != prevHeight+1 && prevHeight != -1 {
		return 0
	}

	if visited[curPos] {
		return 0
	}

	if currentHeight == 9 {
		return 1
	}

	visited[curPos] = true

	totalPaths := 0
	directions := []coord{
		{y: curPos.y - 1, x: curPos.x},
		{y: curPos.y + 1, x: curPos.x},
		{y: curPos.y, x: curPos.x - 1},
		{y: curPos.y, x: curPos.x + 1},
	}

	for _, dir := range directions {
		totalPaths += walk2(grid, dir, currentHeight, visited)
	}

	visited[curPos] = false

	return totalPaths
}

func partOne(filePath string) {
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}
	grid := gridToInts(lines)
	trailheads := getTrailheads(grid)

	totalValidPaths := 0
	for _, start := range trailheads {

		peaks := make(map[coord]bool)
		visited := make(map[coord]bool)
		totalValidPaths += walk(grid, start, -1, visited, peaks)
	}

	fmt.Println(totalValidPaths)
}

func partTwo(filePath string) {
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}
	grid := gridToInts(lines)
	trailheads := getTrailheads(grid)

	totalValidPaths := 0
	for _, start := range trailheads {

		visited := make(map[coord]bool)
		totalValidPaths += walk2(grid, start, -1, visited)
	}

	fmt.Println(totalValidPaths)
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
