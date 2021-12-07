package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputs := getInputsByLine()
	positions := getPositions(inputs[0])
	sort.Ints(positions)

	// Solution 1
	median := positions[len(positions)/2]
	totalFuel := 0

	for _, v := range positions {
		totalFuel += int(math.Abs(float64(median - v)))
	}

	fmt.Printf("Solution 1 : Position: %d, Fuel: %d\n", median, totalFuel)

	// Solution 2
	bestPos := 0
	lowestFuel := 999999999999999999

	for pos := positions[0]; pos < positions[len(positions)-1]; pos++ {
		posFuel := 0
		for _, v := range positions {
			steps := int(math.Abs(float64(pos - v)))
			posFuel += calcFuel(steps)
		}

		if posFuel < lowestFuel {
			bestPos = pos
			lowestFuel = posFuel
		}
	}

	fmt.Printf("Solution 1 : Position: %d, Fuel: %d\n", bestPos, lowestFuel)
}

func calcFuel(steps int) int {
	return steps * (steps + 1) / 2
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

func getPositions(inputLine string) []int {
	var res []int

	pos := strings.Split(inputLine, ",")

	for _, v := range pos {
		vv, _ := strconv.Atoi(v)
		res = append(res, vv)
	}

	return res
}
