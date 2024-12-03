package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/jacobkuzmits/advent24/utils"
)

func partOne(filePath string) {
	// create regex to match mul(x,y)
	// use capture groups on x and y
	regexPattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(regexPattern)

	// initialization
	result := 0

	utils.StreamFile(filePath, func(line string) {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			// make sure we got the capture groups
			if len(match) >= 3 {
				first, err := strconv.Atoi(match[1])
				if err != nil {
					log.Fatalf("error parsing int from string: %v", err)
				}

				second, err := strconv.Atoi(match[2])
				if err != nil {
					log.Fatalf("error parsing int from string: %v", err)
				}

				multResult := first * second

				result += multResult
			}
		}
	})
	fmt.Println(result)
}

func partTwo(filePath string) {
	// create regex to match all of:
	// mul(x,y)
	// do()
	// don't()
	regexPattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(regexPattern)

	// initialization
	result := 0
	enabled := true

	utils.StreamFile(filePath, func(line string) {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			command := match[0][0:3]

			if command == "don" {
				enabled = false
			}

			if command == "do(" {
				enabled = true
			}

			if command == "mul" && enabled {
				// make sure capture groups worked
				if len(match) >= 3 {
					first, err := strconv.Atoi(match[1])
					if err != nil {
						log.Fatalf("error parsing int from string: %v", err)
					}
					second, err := strconv.Atoi(match[2])
					if err != nil {
						log.Fatalf("error parsing int from string: %v", err)
					}
					multResult := first * second
					result += multResult
				}
			}
		}

	})
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
