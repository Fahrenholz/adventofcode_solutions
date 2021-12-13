package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type foldInstruction struct {
	direction string
	val       int
}

type coordList struct {
	values [][]int
	maxX   int
	maxY   int
}

func main() {
	grid, folds := getInputsByLine("./inputs.txt")
	//_, _ = getInputsByLine("./inputs.txt")

	for i, v := range folds {
		switch v.direction {
		case "x":
			grid = fold(grid, v.val, 1, 0)
		case "y":
			grid = fold(grid, v.val, 0, 1)
		}
		if i == 0 {
			fmt.Printf("Solution 1: %d dots\n", len(grid.values))
		}
	}
	printGrid(grid)
}

func printGrid(grid coordList) {
	fmt.Printf("\nSolution 2:\n\n")
	display := make([][]bool, grid.maxY+1)
	for ord := 0; ord <= grid.maxY; ord++ {
		display[ord] = make([]bool, grid.maxX+1)
	}

	for _, v := range grid.values {
		display[v[1]][v[0]] = true
	}

	for _, li := range display {
		for _, pos := range li {
			if pos {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Printf("\n")
	}
}

func fold(grid coordList, foldVal, untouched, operated int) coordList {
	var ngv [][]int
	exists := make(map[string]bool)
	maxX := 0
	maxY := 0

	for _, coords := range grid.values {
		newCoords := make([]int, 2)
		newCoords[untouched] = coords[untouched]
		if coords[operated] < foldVal {
			newCoords[operated] = coords[operated]
		} else {
			newCoords[operated] = foldVal + foldVal - coords[operated]
		}

		if _, ok := exists[fmt.Sprintf("%d#%d", newCoords[0], newCoords[1])]; !ok {
			ngv = append(ngv, newCoords)

			if maxX < newCoords[0] {
				maxX = newCoords[0]
			}

			if maxY < newCoords[1] {
				maxY = newCoords[1]
			}

			exists[fmt.Sprintf("%d#%d", newCoords[0], newCoords[1])] = true
		}
	}

	return coordList{values: ngv, maxX: maxX, maxY: maxY}
}

func getInputsByLine(file string) (coordList, []foldInstruction) {
	dotsR := regexp.MustCompile(`(\d+),(\d+)`)
	fiR := regexp.MustCompile(`fold along ([x|y])=(\d+)`)

	inputFile, err := os.Open(file)
	if err != nil {
		fmt.Println("could not find file")
		os.Exit(1)
	}

	defer inputFile.Close()

	var gridValues [][]int
	var fI []foldInstruction
	maxX := 0
	maxY := 0

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		cur := scanner.Text()
		if matches := dotsR.FindStringSubmatch(cur); len(matches) == 3 {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])

			gridValues = append(gridValues, []int{x, y})
			if maxX < x {
				maxX = x
			}
			if maxY < y {
				maxY = y
			}
		}
		if matches := fiR.FindStringSubmatch(cur); len(matches) == 3 {
			val, _ := strconv.Atoi(matches[2])
			fI = append(fI, foldInstruction{direction: matches[1], val: val})
		}
	}

	return coordList{values: gridValues, maxX: maxX, maxY: maxY}, fI
}
