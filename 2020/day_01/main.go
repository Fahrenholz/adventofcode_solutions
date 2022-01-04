package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputs := getInputsByLine()
	solvePartOne(inputs)
	solvePartTwo(inputs)
}

func solvePartOne(inputs []int) {
	for _, v := range inputs {
		for _, p := range inputs {
			if v+p == 2020 {
				fmt.Printf("Solution 1: %d + %d = 2020, Product: %d\n", v, p, v*p)
				return
			}
		}
	}
}

func solvePartTwo(inputs []int) {
	for _, v := range inputs {
		for _, p := range inputs {
			for _, x := range inputs {
				if v+p+x == 2020 {
					fmt.Printf("Solution 2: %d + %d + %d = 2020, Product: %d\n", v, p, x, v*p*x)
					return
				}
			}
		}
	}
}

func getInputsByLine() []int {
	inputFile, err := os.Open("./inputs.txt")
	if err != nil {
		fmt.Println("could not find file")
		os.Exit(1)
	}

	defer inputFile.Close()

	var inputs []int

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		t, _ := strconv.Atoi(scanner.Text())
		inputs = append(inputs, t)
	}

	return inputs
}
