package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs := getInputsByLine()

	var asteroids [][2]int

	for y, _ := range inputs {
		for x, _ := range inputs[y] {
			if inputs[y][x] == "#" {
				asteroids = append(asteroids, [2]int{x, y})
			}
		}
	}

	fmt.Println(asteroids)
}

func getInputsByLine() [][]string {
	inputFile, err := os.Open(fmt.Sprintf("./%s.txt", os.Args[1]))
	if err != nil {
		fmt.Println("could not find file")
		os.Exit(1)
	}

	defer inputFile.Close()

	var inputs [][]string

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, strings.Split(line, ""))
	}

	return inputs
}
