package aochelper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInputsAsLinesOfIntSlices(delimiter string) [][]int {
	inputFile := GetFileHandle()
	defer inputFile.Close()

	var inputs [][]int
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), delimiter)
		line := make([]int, len(tmp))

		for i, v := range tmp {
			line[i] = ForceInt(v)
		}

		inputs = append(inputs, line)
	}

	return inputs
}

func GetInputAsLinesOfInts() []int {
	inputFile := GetFileHandle()
	defer inputFile.Close()

	var inputs []int

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		inputs = append(inputs, ForceInt(scanner.Text()))
	}

	return inputs
}

func GetInputsAsLinesOfCharSlices(delimiter string) [][]string {
	inputFile := GetFileHandle()
	defer inputFile.Close()

	var inputs [][]string
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		inputs = append(inputs, strings.Split(scanner.Text(), delimiter))
	}

	return inputs
}

func GetInputAsLines() []string {
	inputFile := GetFileHandle()
	defer inputFile.Close()

	var inputs []string

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return inputs
}

func GetFileHandle() *os.File {
	fileName := "inputs"
	if len(os.Args) >= 2 {
		fileName = os.Args[1]
	}

	inputFile, err := os.Open(fmt.Sprintf("./%s.txt", fileName))
	if err != nil {
		fmt.Println("could not find file")
		os.Exit(1)
	}

	return inputFile
}
