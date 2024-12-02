# Advent of Code 2024 Solutions

This repository contains my solutions for [Advent of Code 2024](https://adventofcode.com/2024).

## Project Structure

- template/ - Base template for daily solutions
- utils/ - Common utilities for file reading and parsing
- dayN/ - Individual day solutions

## Usage

### Creating a New Day's Solution

Copy the template directory to create a new day. Replace N with the day number (1-25):

```
cp -r ./template/ ./dayN/
```

### Running Solutions

To run a specific day's solution, where N is the day number:

```
go run . N
```

Examples:

```
go run . 1    # runs day 1's solution
go run . 2    # runs day 2's solution
```

### Input Files

Each day's directory should contain:

- input.txt - Your puzzle input
- testInput.txt - The example input from the puzzle description

## Implementation

Each day's solution contains:

- main.go - Solution code with partOne() and partTwo()
- input.txt - Puzzle input
- testInput.txt - Example input for testing
