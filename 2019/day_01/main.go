package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputs := getInputsByLine()

	sum := 0

	for _, v := range inputs {
		mass, _ := strconv.Atoi(v)
		sum += calculateModuleMass(mass)
	}

	fmt.Printf("Part 1: %d\n", sum)

	// Part 2
	sum = 0

	for _, v := range inputs {
		mass, _ := strconv.Atoi(v)
		localSum := 0
		localSum += calculateModuleMass(mass)
		current := localSum
		for current > 0 {
			current = calculateModuleMass(current)
			if current > 0 {
				localSum += current
			}
		}
		sum += localSum
	}

	fmt.Printf("Part 2: %d\n", sum)
}

func calculateModuleMass(mass int) int {
	return (mass / int(3)) - 2
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
