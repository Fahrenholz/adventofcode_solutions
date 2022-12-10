package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
)

func main() {

	fmt.Println("Part 1: ", solve(2))
	fmt.Println("Part 1: ", solve(10))
}

func solve(knotNb int) int {
	positions := make(map[[2]int]bool)
	rope := make([][2]int, knotNb)
	cycle := 0

	aoch.Walk(aoch.GetInputsAsLinesOfStringSlices(" "), func(el []string) {
		for i := 1; i <= aoch.ForceInt(el[1]); i++ {
			switch el[0] {
			case "R": //x++
				rope[0][0]++
			case "L": //x--
				rope[0][0]--
			case "U": //y++
				rope[0][1]++
			case "D": //y--
				rope[0][1]--
			}

			for j := 1; j < len(rope); j++ {
				deltaX := rope[j-1][0] - rope[j][0]
				deltaY := rope[j-1][1] - rope[j][1]

				switch {
				case abs(deltaY) <= 1 && abs(deltaX) <= 1: //nothing to do
				case deltaY == 0 && deltaX == 2:
					rope[j][0]++
				case deltaY == 0 && deltaX == -2:
					rope[j][0]--
				case deltaX == 0 && deltaY == 2:
					rope[j][1]++
				case deltaX == 0 && deltaY == -2:
					rope[j][1]--
				default:
					if deltaX < 0 {
						rope[j][0]--
					} else {
						rope[j][0]++
					}
					if deltaY < 0 {
						rope[j][1]--
					} else {
						rope[j][1]++
					}
				}
			}
			positions[rope[len(rope)-1]] = true
		}

		cycle += aoch.ForceInt(el[1])
	})

	return len(positions)
}

func abs(my int) int {
	if my < 0 {
		return my * -1
	}
	return my
}
