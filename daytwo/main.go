package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    inputs := getInputsByLine()
	commands := splitInputs(inputs)

	var hPos, depth, aim int

	for _, v := range commands {
		switch v.direction {
		case "forward":
			hPos += v.by
			depth += v.by*aim
		case "down":
			aim += v.by
		case "up":
			aim -= v.by
		}
	}

	fmt.Printf("Horizontal position: %d\nDepth: %d\nMultiply: %d\n", hPos, depth, hPos*depth)
}

type command struct {
	direction string
	by int
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


func splitInputs(inputs []string) []command {
	result := make([]command, len(inputs))

	for i, v := range inputs {
		vSplit := strings.Split(v, " ")

		b, err := strconv.Atoi(vSplit[1])
		if err != nil {
			panic(err)
		}

		result[i] = command{direction: vSplit[0], by: b}
	}

	return result
}