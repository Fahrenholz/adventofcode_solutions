package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
)

var (
	ptsOne     = map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
	ptsGameOne = map[string]map[string]int{"X": {"A": 3, "B": 0, "C": 6}, "Y": {"A": 6, "B": 3, "C": 0}, "Z": {"A": 0, "B": 6, "C": 3}}
	ptsTwo     = map[string]map[string]int{"X": {"A": 3, "B": 1, "C": 2}, "Y": {"A": 1, "B": 2, "C": 3}, "Z": {"A": 2, "B": 3, "C": 1}}
	ptsGameTwo = map[string]int{"X": 0, "Y": 3, "Z": 6}
)

func main() {
	fmt.Println("Part One: ", aoch.Sum(aoch.Map(aoch.GetInputsAsLinesOfStringSlices(" "), func(elem []string) int {
		return ptsOne[elem[1]] + ptsGameOne[elem[1]][elem[0]]
	})))
	fmt.Println("Part Two: ", aoch.Sum(aoch.Map(aoch.GetInputsAsLinesOfStringSlices(" "), func(elem []string) int {
		return ptsGameTwo[elem[1]] + ptsTwo[elem[1]][elem[0]]
	})))
}
