package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
	"sort"
)

func main() {
	lines := aoch.GetInputAsLines()
	var fin, tmp []int
	for _, v := range lines {
		if v == "" {
			fin = append(fin, aoch.SumScore(tmp, func(elem int) int { return elem }))
			tmp = []int{}
			continue
		}
		tmp = append(tmp, aoch.ForceInt(v))
	}

	sort.Ints(fin)

	fmt.Println("PART ONE: ", fin[len(fin)-1])
	fmt.Println("PART TWO: ", fin[len(fin)-1]+fin[len(fin)-2]+fin[len(fin)-3])
}
