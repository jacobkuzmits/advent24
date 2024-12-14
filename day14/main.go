package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jacobkuzmits/advent24/utils"
)

func findEndPos(posX, posY, vecX, vecY, gridX, gridY, steps int) (endX, endY int) {
	totalX := vecX * steps
	totalY := vecY * steps
	modX := totalX % gridX
	modY := totalY % gridY
	endX = (posX + modX) % gridX
	endY = (posY + modY) % gridY
	if endX < 0 {
		endX = gridX + endX
	}
	if endY < 0 {
		endY = gridY + endY
	}

	return endX, endY
}

type coord struct {
	x, y int
}

func partOne(filePath string, gridX int, gridY int, steps int) {
	robotList := map[coord]int{}

	utils.StreamFile(filePath, func(line string) {
		var posX, posY, vecX, vecY int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &posX, &posY, &vecX, &vecY)
		endX, endY := findEndPos(posX, posY, vecX, vecY, gridX, gridY, steps)
		robotList[coord{x: endX, y: endY}] += 1
	})

	// for i := 0; i < gridY; i++ {
	// 	rowString := ""
	// 	for j := 0; j < gridX; j++ {
	// 		if robotCount, exists := robotList[coord{x: j, y: i}]; exists {
	// 			rowString += strconv.Itoa(robotCount)
	// 		} else {
	// 			rowString += "."
	// 		}
	// 	}
	// 	fmt.Println(rowString)
	// }
	nw, ne, se, sw := 0, 0, 0, 0
	for robot, count := range robotList {
		midX := gridX / 2
		midY := gridY / 2
		if robot.x < midX && robot.y < midY {
			nw += count
		}
		if robot.x > midX && robot.y < midY {
			ne += count
		}
		if robot.x > midX && robot.y > midY {
			se += count
		}
		if robot.x < midX && robot.y > midY {
			sw += count
		}
	}
	fmt.Println(nw * ne * se * sw)

}

func partTwo(filePath string, gridX int, gridY int, steps int) string {

	outputString := ""
	for step := 33; step < steps; step += 103 {

		robotList := map[coord]int{}
		outputString += fmt.Sprintf("Step: %d\n", step)

		utils.StreamFile(filePath, func(line string) {
			var posX, posY, vecX, vecY int
			fmt.Sscanf(line, "p=%d,%d v=%d,%d", &posX, &posY, &vecX, &vecY)
			endX, endY := findEndPos(posX, posY, vecX, vecY, gridX, gridY, step)
			robotList[coord{x: endX, y: endY}] += 1
		})

		for i := 0; i < gridY; i++ {
			rowString := ""
			for j := 0; j < gridX; j++ {
				if _, exists := robotList[coord{x: j, y: i}]; exists {
					rowString += "#"
				} else {
					rowString += "."
				}
			}
			outputString += rowString + "\n"
		}

	}
	return outputString
}

func main() {
	fmt.Println("\nPart 1 Test Solution")
	start := time.Now()
	partOne("testInput.txt", 7, 11, 100)
	fmt.Printf("Test execution time: %v\n", time.Since(start))

	fmt.Println("\nPart 1 Actual Solution")
	start = time.Now()
	partOne("input.txt", 101, 103, 100)
	fmt.Printf("Actual execution time: %v\n", time.Since(start))

	// fmt.Println("\nPart 2 Test Solution")
	// start = time.Now()
	// partTwo("testInput.txt")
	// fmt.Printf("Test execution time: %v\n", time.Since(start))

	fmt.Println("\nPart 2 Actual Solution")
	start = time.Now()
	output := partTwo("input.txt", 101, 103, 20000)
	fmt.Printf("Actual execution time: %v\n", time.Since(start))
	file, err := os.Create("partTwoOutput.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
