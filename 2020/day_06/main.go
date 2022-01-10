package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	groups, groupLength := parseInputs(getInputsByLine(os.Args[1]))

	solutionOne := 0
	solutionTwo := 0
	for i, v := range groups {
		solutionOne += len(v)
		for vi := range v {
			if v[vi] == groupLength[i] {
				solutionTwo++
			}
		}
	}

	fmt.Printf("Solution 1: %d.\n", solutionOne)
	fmt.Printf("Solution 2: %d.\n", solutionTwo)
}

func parseInputs(inp []string) ([]map[rune]int, []int) {
	var res []map[rune]int
	var groups []int

	grplen := 0
	group := make(map[rune]int)

	for _, v := range inp {
		if strings.TrimSpace(v) == "" {
			res = append(res, group)
			groups = append(groups, grplen)
			group = make(map[rune]int)
			grplen = 0
			continue
		}
		grplen++
		for _, r := range v {
			group[r] += 1
		}
	}
	res = append(res, group)
	groups = append(groups, grplen)

	return res, groups
}

func getInputsByLine(filename string) []string {
	inputFile, err := os.Open(filename)
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
