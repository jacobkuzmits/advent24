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
	visitedFrom    map[direction]bool
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

type labMap [][]*mapNode

type coord struct {
	x int
	y int
}

var walkVectors = map[string]coord{
	"up":    coord{x: 0, y: -1},
	"right": coord{x: 1, y: 0},
	"down":  coord{x: 0, y: 1},
	"left":  coord{x: -1, y: 0},
}

func parseMap(rows []string) (lm labMap) {
	for rowIndex, row := range rows {
		rowNodes := []*mapNode{}
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

			node := &mapNode{
				isObstacle:     isObstacle,
				beenVisited:    beenVisited,
				isEdge:         isEdge,
				guardPresent:   guardPresent,
				guardDirection: guardDirection,
				visitedFrom:    make(map[direction]bool),
			}
			rowNodes = append(rowNodes, node)
		}
		lm = append(lm, rowNodes)
	}
	return lm
}

func showMap(lm *labMap) {
	for _, row := range *lm {
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

func countVisitedNodes(lm labMap) (r int) {
	for _, row := range lm {
		for _, col := range row {
			if col.beenVisited {
				r += 1
			}
		}
	}
	return r
}

func getVisitedNodes(lm labMap) (nodes []coord) {
	for rowIndex, row := range lm {
		for colIndex, col := range row {
			if col.beenVisited {
				visitedCoord := coord{
					x: colIndex,
					y: rowIndex,
				}
				nodes = append(nodes, visitedCoord)
			}
		}
	}
	return nodes
}

func findGuard(lm labMap) (node *mapNode, pos coord) {
	for rowIndex, row := range lm {
		for colIndex, col := range row {
			if col.guardPresent {
				node = col
				pos = coord{
					x: colIndex,
					y: rowIndex,
				}
			}
		}
	}
	return node, pos
}

func walkGuard(guardPos *coord, startNode *mapNode, lm *labMap, offOfMap *bool) (*mapNode, bool) {
	startNode.beenVisited = true
	startNode.visitedFrom[startNode.guardDirection] = true

	// check if guard is about to walk off of the edge
	if startNode.isEdge {
		*offOfMap = true
		return startNode, true
	}

	// get vector to move based on guard's direction
	vector := walkVectors[startNode.guardDirection.String()]
	targetCoord := coord{
		x: guardPos.x + vector.x,
		y: guardPos.y + vector.y,
	}

	// check if target position is off the map
	if targetCoord.y < 0 || targetCoord.y >= len(*lm) || targetCoord.x < 0 || targetCoord.x >= len((*lm)[0]) {
		*offOfMap = true
		return startNode, true
	}

	// get the target node
	targetNode := (*lm)[targetCoord.y][targetCoord.x]

	// obstacle in the way, turn 90 degrees right
	if targetNode.isObstacle {
		startNode.guardDirection = (startNode.guardDirection + 1) % 4
		return startNode, false
	}

	// no obstacle, walk forwards
	if !targetNode.isObstacle {
		// move the guard position
		guardPos.x = targetCoord.x
		guardPos.y = targetCoord.y
		startNode.guardPresent = false
		targetNode.guardPresent = true
		targetNode.guardDirection = startNode.guardDirection
		startNode.guardDirection = NONE
	}

	return targetNode, false
}

func partOne(filePath string) {
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}

	labMap := parseMap(lines)

	guardNode, guardPos := findGuard(labMap)

	offOfMap := false
	startNode := guardNode
	for !offOfMap {
		startNode, offOfMap = walkGuard(&guardPos, startNode, &labMap, &offOfMap)
	}
	visitedNodeCount := countVisitedNodes(labMap)
	fmt.Printf("Visited %d nodes\n", visitedNodeCount)
}

func partTwo(filePath string) {
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}
	labMap := parseMap(lines)

	guardNode, guardPos := findGuard(labMap)

	offOfMap := false
	startNode := guardNode
	for !offOfMap {
		startNode, offOfMap = walkGuard(&guardPos, startNode, &labMap, &offOfMap)
	}
	visitedNodes := getVisitedNodes(labMap)
	loopPositions := 0
	for _, coord := range visitedNodes {
		newMap := parseMap(lines)
		newGuardNode, newGuardPos := findGuard(newMap)
		newMap[coord.y][coord.x].isObstacle = true
		offOfGrid := false
		startingNode := newGuardNode
		for !offOfGrid {
			if startingNode.visitedFrom[startingNode.guardDirection] {
				loopPositions += 1
				offOfGrid = true
				continue
			}
			startingNode, offOfGrid = walkGuard(&newGuardPos, startingNode, &newMap, &offOfGrid)
		}
	}
	fmt.Printf("%d possible loops\n", loopPositions)
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
