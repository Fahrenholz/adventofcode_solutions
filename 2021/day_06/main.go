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
	fishes := createFishes(inputs)

	fmt.Printf("initial state: total of %d fishes\n", len(fishes))

	for i := 1; i <= 80; i++ {
		fishes = day(fishes)
	}

	fmt.Printf("After 80 days: %d fishes\n", sum(fishes))

	for i := 1; i <= 176; i++ {
		fishes = day(fishes)
	}

	fmt.Printf("After 256 days: %d fishes\n", sum(fishes))
}

func sum(fishes []uint64) (sum uint64) {
	for _, v := range fishes {
		sum += v
	}

	return
}

func day(fishes []uint64) []uint64 {
	res := make([]uint64, 9)

	for i, v := range fishes {
		switch i {
		case 0:
			res[6] += v
			res[8] += v
		default:
			res[i-1] += v
		}
	}

	return res
}

func createFishes(inputs []string) []uint64 {
	stringFishes := strings.Split(inputs[0], ",")

	res := make([]uint64, 9)

	for _, v := range stringFishes {
		fAge, _ := strconv.Atoi(v)
		res[fAge] += 1
	}

	return res
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
