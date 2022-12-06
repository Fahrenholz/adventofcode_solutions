package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
)

func main() {
	inputs := aoch.GetInputsAsLinesOfStringSlices("")[0]
	for i := 0; i < len(inputs)-3; i++ {
		inps := make(map[string]bool)
		inps[inputs[i]] = true
		inps[inputs[i+1]] = true
		inps[inputs[i+2]] = true
		inps[inputs[i+3]] = true

		if len(inps) == 4 {
			fmt.Println("Part One: ", i+4)
			break
		}
	}

	for i := 0; i < len(inputs)-13; i++ {
		inps := make(map[string]bool)
		inps[inputs[i]] = true
		inps[inputs[i+1]] = true
		inps[inputs[i+2]] = true
		inps[inputs[i+3]] = true
		inps[inputs[i+4]] = true
		inps[inputs[i+5]] = true
		inps[inputs[i+6]] = true
		inps[inputs[i+7]] = true
		inps[inputs[i+8]] = true
		inps[inputs[i+9]] = true
		inps[inputs[i+10]] = true
		inps[inputs[i+11]] = true
		inps[inputs[i+12]] = true
		inps[inputs[i+13]] = true

		if len(inps) == 14 {
			fmt.Println("Part Two: ", i+14)
			break
		}
	}
}
