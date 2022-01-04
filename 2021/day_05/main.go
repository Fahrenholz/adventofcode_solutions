package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type coords struct {
	x int
	y int
}

type vector struct {
	from coords
	to   coords
}

func main() {
	inputs := getInputsByLine()
	matrix := buildMatrix(getVectors(inputs))

	cnt := 0
	for _, v := range matrix {
		for _, vv := range v {
			if vv > 1 {
				cnt++
			}
		}
	}

	fmt.Printf("Solution 1: %d\n", cnt)
}

func buildMatrix(vectors []vector, maxX int, maxY int) [][]int {
	matrix := make([][]int, maxY+1)

	for i, _ := range matrix {
		matrix[i] = make([]int, maxX+1)
	}

	for _, v := range vectors {
		from, to := order(v)

		diffX := to.x - from.x
		diffY := to.y - from.y

		if diffX != 0 && diffY != 0 {
			for _, vv := range getDiagonaleCoordinates(from, to) {
				matrix[vv.y][vv.x] += 1
			}
			continue
		}

		if diffY != 0 {
			for i := 0; i <= diffY; i++ {
				matrix[from.y+i][from.x] += 1
			}

			continue
		}

		if diffX != 0 {
			for i := 0; i <= diffX; i++ {
				matrix[from.y][from.x+i] += 1
			}

			continue
		}
	}

	return matrix
}

func order(v vector) (coords, coords) {

	isDiag := isValidDiagonal(v.from.x, v.from.y, v.to.x, v.to.y)

	if (v.from.y > v.to.y || v.from.x > v.to.x) && !isDiag {
		return v.to, v.from
	}

	if isDiag && v.from.x > v.to.x {
		return v.to, v.from
	}

	return v.from, v.to
}

func getVectors(inputs []string) ([]vector, int, int) {
	rp := regexp.MustCompile(`([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)`)
	var result []vector
	var maxX, maxY int

	for _, v := range inputs {
		matches := rp.FindStringSubmatch(v)
		if len(matches) != 5 {
			panic("wrong match length")
		}

		xOne, _ := strconv.Atoi(matches[1])
		yOne, _ := strconv.Atoi(matches[2])
		xTwo, _ := strconv.Atoi(matches[3])
		yTwo, _ := strconv.Atoi(matches[4])

		if xOne != xTwo && yOne != yTwo && !isValidDiagonal(xOne, yOne, xTwo, yTwo) {
			//mixed
			continue
		}

		if xOne > maxX {
			maxX = xOne
		}

		if yOne > maxY {
			maxY = yOne
		}

		if xTwo > maxX {
			maxX = xTwo
		}

		if yTwo > maxY {
			maxY = yTwo
		}

		result = append(result, vector{from: coords{x: xOne, y: yOne}, to: coords{x: xTwo, y: yTwo}})
	}

	return result, maxX, maxY
}

func isValidDiagonal(xOne, yOne, xTwo, yTwo int) bool {
	return xOne-xTwo == yOne-yTwo || xOne-xTwo == (yOne-yTwo)*-1
}

func getDiagonaleCoordinates(from coords, to coords) []coords {
	var c []coords
	incrementFunc := func(i int) int {
		return i + 1
	}
	compareFunc := func(i int) bool {
		return i <= to.y
	}
	if from.y > to.y {
		incrementFunc = func(i int) int {
			return i - 1
		}

		compareFunc = func(i int) bool {
			return i >= to.y
		}
	}
	x := from.x
	for i := from.y; compareFunc(i); i = incrementFunc(i) {
		c = append(c, coords{x: x, y: i})
		x++
	}

	return c
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
