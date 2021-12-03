package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	inputs := getInputsByLine()

	var outputs []string
	outputs = append(outputs, "(N/A - no previous measurement)")
	countIncreases := 0

	for i := 1; i+2 < len(inputs); i++ {
		current := calcSlidingWindow(inputs, i, i+3)
		previous := calcSlidingWindow(inputs, i-1, i+2)

		switch {
		case current < previous:
			outputs = append(outputs, "(decrease)")
		case current > previous:
			outputs = append(outputs, "(increase)")
			countIncreases++
		default:
			outputs = append(outputs, "(N/A - equal)")
		}
	}

	fmt.Printf("Increases: %d\n", countIncreases)
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

func calcSlidingWindow(inputs []string, lower int, upper int) int {
	window := inputs[lower:upper]
	sum := 0

	for _, v := range window {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		sum += val
	}

	return sum
}
