package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
)

func main() {
	inputs := aoch.GetInputsAsLinesOfStringSlices("")[0]
	fmt.Println("Part One: ", searchBlock(inputs, 4))
	fmt.Println("Part Two: ", searchBlock(inputs, 14))
}

func searchBlock(inputs []string, blockLength int) int {
	var res int

	for i := 0; i < len(inputs)-blockLength-1; i++ {
		inps := make(map[string]bool)
		for j := 0; j < blockLength; j++ {
			inps[inputs[i+j]] = true
		}

		if len(inps) == blockLength {
			res = i + blockLength
			break
		}
	}

	return res
}
