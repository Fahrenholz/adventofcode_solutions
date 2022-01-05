package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputs := getInputsByLine()
	solvePartOne(inputs)
	solvePartTwo(inputs)
}

func solvePartOne(inputs []string) {
	trees := followSlope(inputs, 3, 1)

	fmt.Printf("Solution 1: %d trees encountered.\n", trees)
}

func solvePartTwo(inputs []string) {
	slopes := [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	result := 1

	for _, slope := range slopes {
		trees := followSlope(inputs, slope[0], slope[1])
		result *= trees
	}

	fmt.Printf("Solution 2: product of encountered trees is %d.\n", result)
}

func followSlope(inputs []string, right, down int) int {
	x := 0
	y := 0
	trees := 0

	for y < len(inputs) {
		if inputs[y][x%len(inputs[y])] == '#' {
			trees++
		}

		x += right
		y += down
	}

	return trees
}

func getInputsByLine() []string {
	inputFile, err := os.Open("./inputs.txt")
	if err != nil {
		fmt.Println("could not find file")
		os.Exit(1)
	}

	defer inputFile.Close()

	var inputs []string

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return inputs
}
