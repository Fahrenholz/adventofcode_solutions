package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	overlap, fullyContains := parseInputLines(getInputsByLine())

	fmt.Println("PART ONE: ", fullyContains)
	fmt.Println("PART TWO: ", overlap)

}

func parseInputLines(lines []string) (int, int) {
	var fullyContains, overlaps int

	for _, l := range lines {
		li := strings.Split(l, ",")
		li1 := strings.Split(li[0], "-")
		li2 := strings.Split(li[1], "-")
		one := [2]int{suppressStrconvErr(li1[0]), suppressStrconvErr(li1[1])}
		two := [2]int{suppressStrconvErr(li2[0]), suppressStrconvErr(li2[1])}

		if (one[0] <= two[0] && one[1] >= two[1]) || (one[0] >= two[0] && one[1] <= two[1]) {
			fullyContains++
		}

		if (one[0] >= two[0] && one[0] <= two[1]) || (one[0] <= two[0] && one[1] >= two[0]) || (two[0] > one[0] && two[0] < one[1]) || (two[0] < one[0] && two[1] > one[0]) {
			overlaps++
		}
	}

	return overlaps, fullyContains
}

func suppressStrconvErr(toConvert string) int {
	res, _ := strconv.Atoi(toConvert)

	return res
}

func getInputsByLine() []string {
	inputFile, err := os.Open(fmt.Sprintf("./%s.txt", os.Args[1]))
	if err != nil {
		fmt.Println("could not find file")
		os.Exit(1)
	}

	defer inputFile.Close()

	var inputs []string

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return inputs
}
