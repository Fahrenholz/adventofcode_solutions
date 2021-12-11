package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type octopus struct {
	energyLevel int
	highlighted bool
}

func main() {
	var stepFlashes int
	var allFlashing bool

	grid := parseInputs(getInputsByLine())
	totalFlashes := 0

	for i := 1; i <= 1000; i++ {
		stepFlashes, allFlashing = step(grid)
		totalFlashes += stepFlashes
		if i == 100 {
			fmt.Printf("Total Flashes: %d\n", totalFlashes)
		}
		if allFlashing {
			fmt.Printf("All flashing first in round: %d\n", i)
			break
		}
	}

}

func step(grid [][]octopus) (int, bool) {
	var res int

	for ord, li := range grid {
		for abs, _ := range li {
			grid[ord][abs].energyLevel += 1
		}
	}

	for ord, li := range grid {
		for abs, _ := range li {
			if grid[ord][abs].energyLevel > 9 && !grid[ord][abs].highlighted {
				res += flash(grid, ord, abs)
			}
		}
	}

	totalCount := 0
	hightlightedCount := 0
	for ord, li := range grid {
		for abs, _ := range li {
			totalCount++
			if grid[ord][abs].highlighted {
				hightlightedCount++
				grid[ord][abs].energyLevel = 0
				grid[ord][abs].highlighted = false
			}
		}
	}

	return res, totalCount == hightlightedCount
}

func getFlashed(grid [][]octopus, ord, abs int) int {
	if grid[ord][abs].highlighted {
		return 0
	}

	grid[ord][abs].energyLevel++
	if grid[ord][abs].energyLevel <= 9 {
		return 0
	}

	return flash(grid, ord, abs)
}

func flash(grid [][]octopus, ord, abs int) int {
	grid[ord][abs].highlighted = true
	subsequentFlashes := 0

	for lOrd := ord - 1; lOrd <= ord+1; lOrd++ {
		for lAbs := abs - 1; lAbs <= abs+1; lAbs++ {
			if lOrd < 0 || lOrd >= len(grid) || lAbs < 0 || lAbs >= len(grid[lOrd]) || (ord == lOrd && abs == lAbs) {
				continue
			}
			subsequentFlashes += getFlashed(grid, lOrd, lAbs)
		}
	}

	return subsequentFlashes + 1
}

func parseInputs(inp []string) [][]octopus {
	res := make([][]octopus, len(inp))

	for ord, vv := range inp {
		li := strings.Split(vv, "")
		res[ord] = make([]octopus, len(li))
		for abs, v := range li {
			el, _ := strconv.Atoi(v)
			res[ord][abs] = octopus{energyLevel: el}
		}
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
