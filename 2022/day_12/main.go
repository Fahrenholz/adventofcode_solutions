package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
)

const starting = "S"
const finishLine = "E"

func main() {

	inputs := aoch.GetInputsAsLinesOfStringSlices("")

	var start, finish [2]int
	var possibleStarts [][2]int

	for i, _ := range inputs {
		for j, _ := range inputs[i] {
			if inputs[i][j] == starting {
				start[0] = j
				start[1] = i
			}
			if inputs[i][j] == finishLine {
				finish[0] = j
				finish[1] = i
			}

			if inputs[i][j] == starting || inputs[i][j] == "a" {
				possibleStarts = append(possibleStarts, [2]int{j, i})
			}
		}
	}

	inputs[start[1]][start[0]] = "a"
	inputs[finish[1]][finish[0]] = "z"

	fmt.Println("Part one: ", bft(inputs, start, finish))
	fmt.Println("Part two: ", bft(inputs, aoch.MinScore(possibleStarts, func(st [2]int) int {
		res := bft(inputs, st, finish)
		if res == 0 {
			return 1000
		}
		return res
	}), finish))

}

func bft(inputs [][]string, start, finish [2]int) int {
	shortestDistance := make([][]int, len(inputs))
	for i, _ := range inputs {
		shortestDistance[i] = make([]int, len(inputs[i]))
	}

	visited := make(map[[2]int]bool)
	modifiers := []int{-1, 1}

	queue := [][3]int{{start[0], start[1], 0}}

	for len(queue) > 0 {
		elWithSteps := queue[0]
		queue[0] = [3]int{}
		queue = queue[1:]
		el := [2]int{elWithSteps[0], elWithSteps[1]}

		if _, ok := visited[el]; ok {
			continue
		}

		shortestDistance[el[1]][el[0]] = elWithSteps[2]
		visited[el] = true
		if el == finish {
			continue
		}

		for _, modX := range modifiers {
			if el[0]+modX < 0 || el[0]+modX > len(inputs[0])-1 {
				continue
			}
			nE := [2]int{el[0] + modX, el[1]}

			delta := int(inputs[nE[1]][nE[0]][0]) - int(inputs[el[1]][el[0]][0])
			if delta <= 1 && !visited[nE] {
				queue = append(queue, [3]int{nE[0], nE[1], elWithSteps[2] + 1})
			}
		}

		for _, modY := range modifiers {
			if el[1]+modY < 0 || el[1]+modY > len(inputs)-1 {
				continue
			}
			nE := [2]int{el[0], el[1] + modY}

			delta := int(inputs[nE[1]][nE[0]][0]) - int(inputs[el[1]][el[0]][0])
			if delta <= 1 && !visited[nE] {
				queue = append(queue, [3]int{nE[0], nE[1], elWithSteps[2] + 1})
			}
		}
	}

	return shortestDistance[finish[1]][finish[0]]
}
