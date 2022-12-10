package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
)

var crt [6][40]string

func main() {
	xVal := 1
	cycleNb := 1

	var sigStrength []int

	aoch.Walk(aoch.GetInputsAsLinesOfStringSlices(" "), func(el []string) {
		drawCrt(xVal, cycleNb)
		cycleNb++
		if cycleNb == 20 || cycleNb == 60 || cycleNb == 100 || cycleNb == 140 || cycleNb == 180 || cycleNb == 220 {
			sigStrength = append(sigStrength, xVal*cycleNb)
		}
		if el[0] == "addx" {
			drawCrt(xVal, cycleNb)
			xVal += aoch.ForceInt(el[1])
			cycleNb++
			if cycleNb == 20 || cycleNb == 60 || cycleNb == 100 || cycleNb == 140 || cycleNb == 180 || cycleNb == 220 {
				sigStrength = append(sigStrength, xVal*cycleNb)
			}
		}
	})

	fmt.Println("Part One: ", aoch.Sum(sigStrength))
	fmt.Println("Part Two")
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			if crt[i][j] == "#" {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func getCrtRow(cycleNb int) int {
	switch {
	case cycleNb <= 40:
		return 0
	case cycleNb > 40 && cycleNb <= 80:
		return 1
	case cycleNb > 80 && cycleNb <= 120:
		return 2
	case cycleNb > 120 && cycleNb <= 160:
		return 3
	case cycleNb > 160 && cycleNb <= 200:
		return 4
	case cycleNb > 200 && cycleNb <= 240:
		return 5
	}

	return 0
}

func drawCrt(xVal int, cycleNb int) {
	row := getCrtRow(cycleNb)
	pos := (cycleNb - 1) % 40
	if pos == xVal-1 || pos == xVal || pos == xVal+1 {
		crt[row][pos] = "#"
	}
}
