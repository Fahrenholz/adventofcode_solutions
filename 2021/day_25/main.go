package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	grid := parseGrid(getInputsByLine())
	i := 0
	moved := true
	for moved {
		var movedEast, movedSouth int
		grid, movedEast = weGoEast(grid)
		grid, movedSouth = weGoSouth(grid)
		moved = movedEast != 0 || movedSouth != 0
		i++
	}

	fmt.Printf("Solution 1: stopped moving after %d steps\n", i)
}

func cp(grid [][]string) [][]string {
	nGrid := make([][]string, len(grid))
	for i, _ := range grid {
		nGrid[i] = make([]string, len(grid[i]))
		for j, _ := range grid[i] {
			nGrid[i][j] = grid[i][j]
		}
	}

	return nGrid
}

func weGoEast(grid [][]string) ([][]string, int) {
	nGrid := cp(grid)
	moves := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			nextX := x + 1
			if nextX == len(grid[y]) {
				nextX = 0
			}
			if grid[y][x] == ">" && grid[y][nextX] == "." {
				moves++
				nGrid[y][x] = "."
				nGrid[y][nextX] = ">"
			}
		}
	}

	return nGrid, moves
}

func weGoSouth(grid [][]string) ([][]string, int) {
	nGrid := cp(grid)
	moves := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			nextY := y + 1
			if nextY == len(grid) {
				nextY = 0
			}
			if grid[y][x] == "v" && grid[nextY][x] == "." {
				moves++
				nGrid[y][x] = "."
				nGrid[nextY][x] = "v"
			}
		}
	}

	return nGrid, moves
}

func parseGrid(inp []string) [][]string {
	grid := make([][]string, len(inp))
	for i, v := range inp {
		grid[i] = strings.Split(v, "")
	}
	return grid
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
