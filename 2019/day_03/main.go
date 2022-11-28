package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := getInputsByLine(os.Args[1])

	coordsC1 := getCoordinates(strings.Split(inputs[0], ","))
	coordsC2 := getCoordinates(strings.Split(inputs[1], ","))

	var intersections [][3]int

	for i := 1; i < len(coordsC1); i++ {
		for j := 1; j < len(coordsC2); j++ {
			if stableX(coordsC1[i], coordsC1[i-1]) && stableY(coordsC2[j], coordsC2[j-1]) && crossesX([2][3]int{coordsC2[j-1], coordsC2[j]}, coordsC1[i][0]) && crossesY([2][3]int{coordsC1[i-1], coordsC1[i]}, coordsC2[j][1]) {
				steps := 0
				for _, v := range coordsC1[:i] {
					steps += v[2]
				}
				for _, v := range coordsC2[:j] {
					steps += v[2]
				}

				steps += abs(coordsC1[i][0]-coordsC2[j-1][0]) + abs(coordsC2[j][1]-coordsC1[i-1][1])

				if !(coordsC1[i][0] == 0 && coordsC2[j][1] == 0) {
					intersections = append(intersections, [3]int{coordsC1[i][0], coordsC2[j][1], steps})
				}
			} else if stableX(coordsC2[j], coordsC2[j-1]) && stableY(coordsC1[i], coordsC1[i-1]) && crossesX([2][3]int{coordsC1[i-1], coordsC1[i]}, coordsC2[j][0]) && crossesY([2][3]int{coordsC2[j-1], coordsC2[j]}, coordsC1[i][1]) {
				steps := 0
				for _, v := range coordsC1[:i] {
					steps += v[2]
				}

				for _, v := range coordsC2[:j] {
					steps += v[2]
				}

				steps += abs(coordsC2[j][0]-coordsC1[i-1][0]) + abs(coordsC1[i][1]-coordsC2[j-1][1])

				if !(coordsC2[j][0] == 0 && coordsC1[i][1] == 0) {
					intersections = append(intersections, [3]int{coordsC2[j][0], coordsC1[i][1], steps})
				}
			}
		}
	}

	min := 0
	minSteps := 0
	initialized := false
	for _, v := range intersections {
		if !initialized {
			min = getManhattan(v)
			minSteps = v[2]
			initialized = true
			continue
		}

		mh := getManhattan(v)
		if min > mh {
			min = mh
		}

		if minSteps > v[2] {
			minSteps = v[2]
		}
	}

	fmt.Printf("Part 1: %d\n", min)
	fmt.Printf("Part 2: %d\n", minSteps)
}

func stableX(pointA, pointB [3]int) bool {
	return pointA[0] == pointB[0]
}

func stableY(pointA, pointB [3]int) bool {
	return pointA[1] == pointB[1]
}

func crossesX(vector [2][3]int, x int) bool {
	return (vector[0][0] <= x && vector[1][0] >= x) || (vector[1][0] <= x && vector[0][0] >= x)
}

func crossesY(vector [2][3]int, y int) bool {
	return (vector[0][1] <= y && vector[1][1] >= y) || (vector[1][1] <= y && vector[0][1] >= y)
}

func getManhattan(point [3]int) int {
	return abs(0-point[0]) + abs(0-point[1])
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func getCoordinates(steps []string) [][3]int {
	var coordinates [][3]int
	x := 0
	y := 0

	coordinates = append(coordinates, [3]int{0, 0, 0})

	for _, step := range steps {
		direction := step[0:1]
		change, _ := strconv.Atoi(step[1:])
		switch direction {
		case "L":
			x += (change * -1)
		case "R":
			x += change
		case "D":
			y += (change * -1)
		case "U":
			y += change
		}

		coordinates = append(coordinates, [3]int{x, y, change})
	}

	return coordinates
}

func getInputsByLine(filename string) []string {
	inputFile, err := os.Open(fmt.Sprintf("./%s.txt", filename))
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
