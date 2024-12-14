package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/jacobkuzmits/advent24/utils"
)

type button struct {
	cost int
	x    int
	y    int
}

type prize struct {
	x int
	y int
}

type game struct {
	buttons   []button
	prize     prize
	solutions []gameSolution
}

type gameSolution struct {
	buttonPresses map[button]int
	cost          int
}

func parseGames(lines []string) (games []game) {
	for i := 0; i < len(lines); i += 4 {
		buttonAStr := lines[i]
		buttonBStr := lines[i+1]
		prizeStr := lines[i+2]
		var buttonAX, buttonAY, buttonBX, buttonBY, prizeX, prizeY int
		fmt.Sscanf(buttonAStr, "Button A: X+%d, Y+%d", &buttonAX, &buttonAY)
		fmt.Sscanf(buttonBStr, "Button B: X+%d, Y+%d", &buttonBX, &buttonBY)
		fmt.Sscanf(prizeStr, "Prize: X=%d, Y=%d", &prizeX, &prizeY)
		buttonA := button{
			cost: 3,
			x:    buttonAX,
			y:    buttonAY,
		}
		buttonB := button{
			cost: 1,
			x:    buttonBX,
			y:    buttonBY,
		}
		prize := prize{
			x: prizeX,
			y: prizeY,
		}
		game := game{
			buttons:   []button{buttonA, buttonB},
			prize:     prize,
			solutions: []gameSolution{},
		}
		games = append(games, game)
	}
	return games
}

func (g game) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\nGame Prize: X=%d Y=%d\n", g.prize.x, g.prize.y))
	sb.WriteString(fmt.Sprintf("Button A: X=%d Y=%d\n", g.buttons[0].x, g.buttons[1].y))
	sb.WriteString(fmt.Sprintf("Button B: X=%d Y=%d", g.buttons[1].x, g.buttons[1].y))
	return sb.String()
}

func (g *game) findSolutions() {
	maxAX := g.prize.x / g.buttons[0].x
	maxAY := g.prize.y / g.buttons[0].y
	maxA := int(math.Min(float64(maxAX), float64(maxAY)))
	maxBX := g.prize.x / g.buttons[1].x
	maxBY := g.prize.y / g.buttons[1].y
	maxB := int(math.Min(float64(maxBX), float64(maxBY)))

outer:
	for a := 0; a <= maxA; a++ {
		for b := 0; b <= maxB; b++ {
			posX := a*g.buttons[0].x + b*g.buttons[1].x
			posY := a*g.buttons[0].y + b*g.buttons[1].y
			if posX > g.prize.x || posY > g.prize.y {
				continue outer
			}
			if posX == g.prize.x && posY == g.prize.y {
				newSolution := gameSolution{
					buttonPresses: map[button]int{
						g.buttons[0]: a,
						g.buttons[1]: b,
					},
					cost: a*3 + b,
				}
				g.solutions = append(g.solutions, newSolution)
			}
		}
	}

}

func (g *game) modifyPrize() {
	g.prize.x = g.prize.x + 10000000000000
	g.prize.y = g.prize.y + 10000000000000
}

func partOne(filePath string) {
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}

	games := parseGames(lines)
	totalCost := 0
	for _, game := range games {
		(&game).findSolutions()
		if len(game.solutions) > 0 {
			lowestCost := math.MaxInt
			for _, solution := range game.solutions {
				if solution.cost < lowestCost {
					lowestCost = solution.cost
				}
			}
			totalCost += lowestCost
		}
	}
	fmt.Println(totalCost)
}

func partTwo(filePath string) {
	lines, err := utils.GetLines(filePath)
	if err != nil {
		log.Fatalf("utils.GetLines() error: %v", err)
	}

	games := parseGames(lines)
	totalCost := int64(0)
	for _, game := range games {
		(&game).modifyPrize()
		x1, y1 := int64(game.buttons[0].x), int64(game.buttons[0].y)
		x2, y2 := int64(game.buttons[1].x), int64(game.buttons[1].y)
		z1, z2 := int64(game.prize.x), int64(game.prize.y)

		a, b, err := solveLinearEquations(x1, y1, x2, y2, z1, z2)
		if err != nil {
			// fmt.Println("No solution found for game:", game)
			continue
		}

		// Calculate the cost
		cost := a*int64(game.buttons[0].cost) + b*int64(game.buttons[1].cost)
		totalCost += cost
	}
	fmt.Println(totalCost)
}

func solveLinearEquations(x1, y1, x2, y2, z1, z2 int64) (int64, int64, error) {
	// Calculate the determinant
	det := x1*y2 - y1*x2

	// Check if the determinant is zero
	if det == 0 {
		return 0, 0, fmt.Errorf("no unique solution")
	}

	// Calculate the numerators
	i := y2*z1 - x2*z2
	j := -y1*z1 + x1*z2

	// Check if the numerators are divisible by the determinant
	if i%det != 0 || j%det != 0 {
		return 0, 0, fmt.Errorf("no integer solution")
	}

	// Calculate the values of a and b
	a := i / det
	b := j / det

	return a, b, nil
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
