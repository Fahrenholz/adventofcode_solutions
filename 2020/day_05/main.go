package main

import (
	"bufio"
	"fmt"
	"os"
)

type boardingPass struct {
	row       string
	column    string
	rowInt    int
	columnInt int
	id        int
}

func (p boardingPass) calculateAttributes() boardingPass {
	startRow := 0
	endRow := 127
	for _, v := range p.row {
		if v == 'F' {
			endRow = endRow - (((endRow + 1) - startRow) / 2)
		}
		if v == 'B' {
			startRow = startRow + (((endRow + 1) - startRow) / 2)
		}
	}

	p.rowInt = endRow

	startCol := 0
	endCol := 7

	for _, v := range p.column {
		if v == 'L' {
			endCol = endCol - (((endCol + 1) - startCol) / 2)
		}
		if v == 'R' {
			startCol = startCol + (((endCol + 1) - startCol) / 2)
		}
	}

	p.columnInt = endCol
	p.id = (p.rowInt * 8) + p.columnInt

	return p
}

func main() {
	inputs := getInputsByLine()
	var passes []boardingPass
	max := 0
	ids := make(map[int]bool)
	for _, v := range inputs {
		p := v.calculateAttributes()
		passes = append(passes, p)
		if p.id > max {
			max = p.id
		}
		ids[p.id] = true
	}

	fmt.Printf("Solution 1: Max seat-id is %d.\n", max)

	for i := 1; i < max; i++ {
		_, okLower := ids[i-1]
		_, okUpper := ids[i+1]
		_, ok := ids[i]
		if okLower && okUpper && !ok {
			fmt.Printf("Solution 2: Seat-ID %d.\n", i)
			break
		}
	}
}

func getInputsByLine() []boardingPass {
	inputFile, err := os.Open("./inputs.txt")
	if err != nil {
		fmt.Println("could not find file")
		os.Exit(1)
	}

	defer inputFile.Close()

	var inputs []boardingPass

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		passData := scanner.Text()
		inputs = append(inputs, boardingPass{
			row:    passData[0:7],
			column: passData[7:],
		})
	}

	return inputs
}
