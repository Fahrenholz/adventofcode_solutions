package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
	"math"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	maxY, maxX, minX, caveMap := newCaveMap()
	unitsOfSand := 0
	for {
		newSand := [2]int{500, 0}
		unitsOfSand++
		for {
			if newSand[1] == maxY+1 {
				caveMap[[2]int{newSand[0], newSand[1]}] = 'o'
				break
			}
			if r, ok := caveMap[[2]int{newSand[0], newSand[1] + 1}]; ok && (r == '#' || r == 'o') {
				if rr, ok2 := caveMap[[2]int{newSand[0] - 1, newSand[1] + 1}]; ok2 && (rr == '#' || rr == 'o') {
					if rrr, ok3 := caveMap[[2]int{newSand[0] + 1, newSand[1] + 1}]; ok3 && (rrr == '#' || rrr == 'o') {
						caveMap[[2]int{newSand[0], newSand[1]}] = 'o'
						break
					} else {
						newSand[0]++
						newSand[1]++
					}
				} else {
					newSand[0]--
					newSand[1]++
				}
			} else {
				newSand[1]++
			}
		}
		if r, ok := caveMap[[2]int{500, 0}]; ok && r == 'o' {
			break
		}
	}

	drawP2(caveMap, minX, maxX, maxY)
	fmt.Println("Part two: ", unitsOfSand)
}

func partOne() {
	maxY, maxX, minX, caveMap := newCaveMap()
	fallsThrough := false
	unitsOfSand := 0
	for !fallsThrough {
		newSand := [2]int{500, 0}
		unitsOfSand++
		for {
			if newSand[1] > maxY {
				fallsThrough = true
				break
			}
			if r, ok := caveMap[[2]int{newSand[0], newSand[1] + 1}]; ok && (r == '#' || r == 'o') {
				if rr, ok2 := caveMap[[2]int{newSand[0] - 1, newSand[1] + 1}]; ok2 && (rr == '#' || rr == 'o') {
					if rrr, ok3 := caveMap[[2]int{newSand[0] + 1, newSand[1] + 1}]; ok3 && (rrr == '#' || rrr == 'o') {
						caveMap[[2]int{newSand[0], newSand[1]}] = 'o'
						break
					} else {
						newSand[0]++
						newSand[1]++
					}
				} else {
					newSand[0]--
					newSand[1]++
				}
			} else {
				newSand[1]++
			}
		}
	}

	drawP1(caveMap, minX, maxX, maxY)
	fmt.Println("Part one: ", unitsOfSand-1)
}

func newCaveMap() (int, int, int, map[[2]int]rune) {
	maxY := 0
	maxX := 0
	minX := math.MaxInt

	caveMap := make(map[[2]int]rune)

	aoch.Walk(aoch.GetInputAsLines(), func(line string) {
		var points [][2]int
		pointsStr := strings.Split(line, " -> ")
		for _, ptS := range pointsStr {
			pt := strings.Split(ptS, ",")
			if maxY < aoch.ForceInt(pt[1]) {
				maxY = aoch.ForceInt(pt[1])
			}
			if maxX < aoch.ForceInt(pt[0]) {
				maxX = aoch.ForceInt(pt[0])
			}
			if minX > aoch.ForceInt(pt[0]) {
				minX = aoch.ForceInt(pt[0])
			}
			points = append(points, [2]int{aoch.ForceInt(pt[0]), aoch.ForceInt(pt[1])})
		}

		for i := 0; i < len(points)-1; i++ {
			if points[i][0] != points[i+1][0] {
				//line on x
				for x := aoch.Min([]int{points[i][0], points[i+1][0]}); x <= aoch.Max([]int{points[i][0], points[i+1][0]}); x++ {
					caveMap[[2]int{x, points[i][1]}] = '#'
				}
			} else {
				for y := aoch.Min([]int{points[i][1], points[i+1][1]}); y <= aoch.Max([]int{points[i][1], points[i+1][1]}); y++ {
					caveMap[[2]int{points[i][0], y}] = '#'
				}
			}
		}
	})
	return maxY, maxX, minX, caveMap
}

func drawP1(caveMap map[[2]int]rune, minX, maxX, maxY int) {
	for y := 0; y <= maxY; y++ {
		fmt.Print("|")
		for x := minX - 1; x <= maxX+1; x++ {
			if r, ok := caveMap[[2]int{x, y}]; ok {
				switch r {
				case '#':
					fmt.Print("█")
				case 'o':
					fmt.Print("o")
				default:
					fmt.Print(" ")
				}
			} else {

				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}
}

func drawP2(caveMap map[[2]int]rune, minX, maxX, maxY int) {
	for y := 0; y <= maxY+1; y++ {
		fmt.Print("|")
		for x := minX - 1; x <= maxX+1; x++ {
			if r, ok := caveMap[[2]int{x, y}]; ok {
				switch r {
				case '#':
					fmt.Print("█")
				case 'o':
					fmt.Print("o")
				default:
					fmt.Print(" ")
				}
			} else {

				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}
	for x := minX - 1; x <= maxX+1; x++ {
		fmt.Print("█")
	}
	fmt.Println()
}
