package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
	"strings"
)

func main() {
	overlap, fullyContains := parseInputLines(aoch.GetInputAsLines())

	fmt.Println("PART ONE: ", fullyContains)
	fmt.Println("PART TWO: ", overlap)

}

func parseInputLines(lines []string) (int, int) {
	var fullyContains, overlaps int

	for _, l := range lines {
		li := strings.Split(l, ",")
		li1 := strings.Split(li[0], "-")
		li2 := strings.Split(li[1], "-")
		one := [2]int{aoch.ForceInt(li1[0]), aoch.ForceInt(li1[1])}
		two := [2]int{aoch.ForceInt(li2[0]), aoch.ForceInt(li2[1])}

		if (one[0] <= two[0] && one[1] >= two[1]) || (one[0] >= two[0] && one[1] <= two[1]) {
			fullyContains++
		}

		if (one[0] >= two[0] && one[0] <= two[1]) || (one[0] <= two[0] && one[1] >= two[0]) || (two[0] > one[0] && two[0] < one[1]) || (two[0] < one[0] && two[1] > one[0]) {
			overlaps++
		}
	}

	return overlaps, fullyContains
}
