package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jacobkuzmits/advent24/utils"
)

func sumList(nums []int) (r int) {
	for _, num := range nums {
		r += num
	}
	return r
}

type regionNode struct {
	y           int
	x           int
	beenVisited bool
	beenCounted bool
	plant       rune
	fences      []fence
	regionID    int
}

func (n regionNode) hasFence(f fence) bool {
	for _, fence := range n.fences {
		if fence == f {
			return true
		}
	}
	return false
}

type fence int

const (
	UP fence = iota
	RIGHT
	DOWN
	LEFT
)

type region struct {
	nodes     []regionNode
	perimeter int
	sides     int
	plant     rune
	ID        int
}

func (r region) fenceCost(part int) (area int) {
	if part == 1 {
		area = len(r.nodes) * r.perimeter
	} else if part == 2 {
		area = len(r.nodes) * r.sides
	}
	return area
}

type farmGrid [][]regionNode

func createFarm(lines []string) (farm farmGrid) {
	for row, line := range lines {
		newRow := []regionNode{}
		for col, plant := range line {
			node := regionNode{
				y:           row,
				x:           col,
				plant:       plant,
				beenVisited: false,
				fences:      []fence{},
				regionID:    0,
			}
			newRow = append(newRow, node)
		}
		farm = append(farm, newRow)
	}
	return farm
}

func markFences(node *regionNode, farm farmGrid) {
	if node.beenVisited {
		return
	}
	// up
	if node.y == 0 {
		node.fences = append(node.fences, UP)
	} else if node.plant != farm[node.y-1][node.x].plant {
		node.fences = append(node.fences, UP)
	}
	// right
	if node.x >= len(farm[0])-1 {
		node.fences = append(node.fences, RIGHT)
	} else if node.plant != farm[node.y][node.x+1].plant {
		node.fences = append(node.fences, RIGHT)
	}
	// down
	if node.y >= len(farm)-1 {
		node.fences = append(node.fences, DOWN)
	} else if node.plant != farm[node.y+1][node.x].plant {
		node.fences = append(node.fences, DOWN)
	}
	// left
	if node.x == 0 {
		node.fences = append(node.fences, LEFT)
	} else if node.plant != farm[node.y][node.x-1].plant {
		node.fences = append(node.fences, LEFT)
	}
}

func walkRegion(node *regionNode, farm farmGrid, region *region) {
	if node.beenVisited {
		return
	}
	// add the node to the region
	node.beenVisited = true
	if node.regionID == 0 {
		node.regionID = region.ID
	}
	region.nodes = append(region.nodes, *node)
	region.perimeter += len(node.fences)

	// walk each direction if it is part of the region
	if !node.hasFence(UP) {
		upNode := &farm[node.y-1][node.x]
		if node.plant == upNode.plant {
			walkRegion(upNode, farm, region)
		}
	}
	if !node.hasFence(DOWN) {
		downNode := &farm[node.y+1][node.x]
		if node.plant == downNode.plant {
			walkRegion(downNode, farm, region)
		}
	}
	if !node.hasFence(RIGHT) {
		rightNode := &farm[node.y][node.x+1]
		if node.plant == rightNode.plant {
			walkRegion(rightNode, farm, region)
		}
	}
	if !node.hasFence(LEFT) {
		leftNode := &farm[node.y][node.x-1]
		if node.plant == leftNode.plant {
			walkRegion(leftNode, farm, region)
		}
	}
}

func countSides(node *regionNode, farm farmGrid, region *region) {
	if node.beenCounted {
		return
	}
	node.beenCounted = true

	// count the corners
	var north, ne, east, se, south, sw, west, nw int
	// look north
	if node.y > 0 {
		north = farm[node.y-1][node.x].regionID
		// look nw
		if node.x > 0 {
			nw = farm[node.y-1][node.x-1].regionID
		}
		// look ne
		if node.x < len(farm[node.y])-1 {
			ne = farm[node.y-1][node.x+1].regionID
		}
	}
	// look west
	if node.x > 0 {
		west = farm[node.y][node.x-1].regionID
	}
	// look east
	if node.x < len(farm[node.y])-1 {
		east = farm[node.y][node.x+1].regionID
	}
	// look south
	if node.y < len(farm)-1 {
		south = farm[node.y+1][node.x].regionID
		// look sw
		if node.x > 0 {
			sw = farm[node.y+1][node.x-1].regionID
		}
		// look se
		if node.x < len(farm[node.y])-1 {
			se = farm[node.y+1][node.x+1].regionID
		}
	}

	// check for corner ne
	if node.regionID != ne {
		if node.regionID == north && node.regionID == east {
			region.sides += 1
		}
		if node.regionID != north && node.regionID != east {
			region.sides += 1
		}
	}
	if node.regionID == ne && node.regionID != north && node.regionID != east {
		region.sides += 1
	}

	// check for corner se
	if node.regionID != se {
		if node.regionID == south && node.regionID == east {
			region.sides += 1
		}
		if node.regionID != south && node.regionID != east {
			region.sides += 1
		}
	}
	if node.regionID == se && node.regionID != south && node.regionID != east {
		region.sides += 1
	}

	// check for corner sw
	if node.regionID != sw {
		if node.regionID == south && node.regionID == west {
			region.sides += 1
		}
		if node.regionID != south && node.regionID != west {
			region.sides += 1
		}
	}
	if node.regionID == sw && node.regionID != south && node.regionID != west {
		region.sides += 1
	}

	// check for corner nw
	if node.regionID != nw {
		if node.regionID == north && node.regionID == west {
			region.sides += 1
		}
		if node.regionID != north && node.regionID != west {
			region.sides += 1
		}
	}
	if node.regionID == nw && node.regionID != north && node.regionID != west {
		region.sides += 1
	}

	// walk each direction if it is part of the region
	if !node.hasFence(UP) {
		upNode := &farm[node.y-1][node.x]
		if node.plant == upNode.plant {
			countSides(upNode, farm, region)
		}
	}
	if !node.hasFence(DOWN) {
		downNode := &farm[node.y+1][node.x]
		if node.plant == downNode.plant {
			countSides(downNode, farm, region)
		}
	}
	if !node.hasFence(RIGHT) {
		rightNode := &farm[node.y][node.x+1]
		if node.plant == rightNode.plant {
			countSides(rightNode, farm, region)
		}
	}
	if !node.hasFence(LEFT) {
		leftNode := &farm[node.y][node.x-1]
		if node.plant == leftNode.plant {
			countSides(leftNode, farm, region)
		}
	}
}

func partOne(filePath string) {
	// get lines as a slice of strings
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}

	farm := createFarm(lines)

	// mark fences at farm edge and region edge boundaries
	for i := range farm {
		for j := range farm[i] {
			markFences(&farm[i][j], farm)
		}
	}

	// find all regions
	regions := []region{}
	for i := range farm {
		for j := range farm[i] {
			newRegion := region{}
			walkRegion(&farm[i][j], farm, &newRegion)
			if len(newRegion.nodes) > 0 {
				regions = append(regions, newRegion)
			}
		}
	}

	totalCost := 0
	for _, region := range regions {
		totalCost += region.fenceCost(1)
	}
	fmt.Println(totalCost)
}

func partTwo(filePath string) {
	// get lines as a slice of strings
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}

	farm := createFarm(lines)

	// mark fences at farm edge and region edge boundaries
	for i := range farm {
		for j := range farm[i] {
			markFences(&farm[i][j], farm)
		}
	}

	// find all regions
	regionID := 1
	regions := []region{}
	for i := range farm {
		for j := range farm[i] {
			newRegion := region{}
			newRegion.ID = regionID
			walkRegion(&farm[i][j], farm, &newRegion)
			countSides(&farm[i][j], farm, &newRegion)
			if len(newRegion.nodes) > 0 {
				regions = append(regions, newRegion)
				regionID++
			}

		}
	}

	totalCost := 0
	for _, region := range regions {
		totalCost += region.fenceCost(2)
	}
	fmt.Println(totalCost)
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
