package main

import (
	"fmt"
	"log"

	"github.com/jacobkuzmits/advent24/utils"
)

type mapNode struct {
	isObstacle     bool
	beenVisited    bool
	isEdge         bool
	guardPresent   bool
	guardDirection direction
}

type direction int

const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
	NONE
)

func (d direction) String() string {
	return [...]string{"up", "right", "down", "left", "none"}[d]
}

type labMap [][]mapNode

type coord struct {
	x int
	y int
}

func parseMap(rows []string) (lm labMap) {
	for rowIndex, row := range rows {
		rowNodes := []mapNode{}
		for colIndex, col := range row {
			isObstacle := string(col) == "#"
			beenVisited := false
			isEdge := (rowIndex == 0) || (rowIndex == len(rows)-1) || (colIndex == 0) || (colIndex == len(row)-1)
			guardFacingUp := string(col) == "^"
			guardFacingRight := string(col) == ">"
			guardFacingDown := string(col) == "v"
			guardFacingLeft := string(col) == "<"
			guardPresent := guardFacingUp || guardFacingRight || guardFacingDown || guardFacingLeft
			var guardDirection direction
			if guardFacingUp {
				guardDirection = UP
			} else if guardFacingRight {
				guardDirection = RIGHT
			} else if guardFacingDown {
				guardDirection = DOWN
			} else if guardFacingLeft {
				guardDirection = LEFT
			} else {
				guardDirection = NONE
			}

			node := mapNode{
				isObstacle:     isObstacle,
				beenVisited:    beenVisited,
				isEdge:         isEdge,
				guardPresent:   guardPresent,
				guardDirection: guardDirection,
			}
			rowNodes = append(rowNodes, node)
		}
		lm = append(lm, rowNodes)
	}
	return lm
}

func showMap(lm labMap) {
	for _, row := range lm {
		rowString := ""
		for _, col := range row {
			if col.isObstacle {
				rowString += "#"
			} else if col.isEdge {
				rowString += "-"
			} else if col.beenVisited {
				rowString += "*"
			} else if col.guardDirection == UP {
				rowString += "^"
			} else if col.guardDirection == RIGHT {
				rowString += ">"
			} else if col.guardDirection == DOWN {
				rowString += "v"
			} else if col.guardDirection == LEFT {
				rowString += "<"
			} else {
				rowString += " "
			}

		}
		fmt.Println(rowString)
	}
}

// func findGuard()

func partOne(filePath string) {
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}
	labMap := parseMap(lines)
	showMap(labMap)
}

func partTwo(filePath string) {
	// get lines as a slice of strings
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}
	fmt.Println(lines[0])

	utils.StreamFile(filePath, func(line string) {
		// do something with each line
	})
}

func main() {
	fmt.Println("\nPart 1 Test Solution")
	partOne("testInput.txt")

	// fmt.Println("\nPart 1 Actual Solution")
	// partOne("input.txt")

	// fmt.Println("\nPart 2 Test Solution")
	// partTwo("testInput.txt")

	// fmt.Println("\nPart 2 Actual Solution")
	// partTwo("input.txt")
}
