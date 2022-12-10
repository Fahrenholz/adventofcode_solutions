package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
)

func main() {
	inputs := aoch.GetInputsAsLinesOfIntSlices("")
	flipped := aoch.Flip(inputs)
	var visible int
	var scenicScores []int

	for row, _ := range inputs {
		for col, _ := range inputs[row] {
			switch {
			case row == 0 || col == 0:
				visible++
			case row == len(inputs[0])-1 || col == len(inputs)-1:
				visible++
			case aoch.MaxIdx(flipped[col][:row+1], false) == row || aoch.MaxIdx(flipped[col][row:], true)+row == row:
				visible++
			case aoch.MaxIdx(inputs[row][:col+1], false) == col || aoch.MaxIdx(inputs[row][col:], true)+col == col:
				visible++
			}

			scenicScores = append(scenicScores, getScenicScore(inputs, flipped, row, col))
		}
	}

	fmt.Println("Part One: ", visible)
	fmt.Println("Part Two: ", aoch.Max(scenicScores))

}

func getScenicScore(inputs, flipped [][]int, row, col int) int {
	toLeft := 0
	if col != 0 {
		toLeft = LastVisibleIdxTree(aoch.Reverse(inputs[row][:col]), inputs[row][col]) + 1
	}
	toRight := 0
	if col != len(inputs[row])-1 {
		toRight = LastVisibleIdxTree(inputs[row][col+1:], inputs[row][col]) + 1
	}
	up := 0
	if row != 0 {
		up = LastVisibleIdxTree(aoch.Reverse(flipped[col][:row]), inputs[row][col]) + 1
	}
	down := 0
	if row != len(flipped[col])-1 {
		down = LastVisibleIdxTree(flipped[col][row+1:], inputs[row][col]) + 1
	}

	return toLeft * toRight * up * down
}

func LastVisibleIdxTree(vals []int, treeHeight int) int {
	var max, maxIdx int

	for i, _ := range vals {
		if vals[i] > max {
			max = vals[i]
			maxIdx = i
		}
		if treeHeight > max && maxIdx != i {
			max = vals[i]
			maxIdx = i
		}
		if max >= treeHeight {
			break
		}
	}

	return maxIdx
}
