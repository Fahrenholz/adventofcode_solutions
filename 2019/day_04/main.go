package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputs := strings.Split(getInputsByLine()[0], "-")
	lower, _ := strconv.Atoi(inputs[0])
	upper, _ := strconv.Atoi(inputs[1])

	pp, pp2 := getPossiblePasswords(lower, upper)

	fmt.Printf("Part 1: %d\n", len(pp))
	fmt.Printf("Part 2: %d\n", len(pp2))

}

func getPossiblePasswords(lower int, upper int) ([]int, []int) {
	var passwords []int
	var passwordsPartTwo []int

	cnt := 0
	for curr := lower; curr <= upper; curr++ {
		mapparts := make(map[int]int)
		strCurr := strconv.Itoa(curr)
		var currSl [6]int
		for i, v := range strCurr {
			part, _ := strconv.Atoi(string(v))
			mapparts[part] += 1
			currSl[i] = part
		}

		mapOccurences := make(map[int]int)
		for _, v := range mapparts {
			mapOccurences[v]++
		}

		if sort.SliceIsSorted(currSl, func(i, j int) bool {
			return currSl[i] < currSl[j]
		}) {
			cnt++
			pt1 := false
			pt2 := false
			for i := range mapOccurences {
				if i >= 2 {
					pt1 = true
				}
				if i == 2 {
					pt2 = true
				}
			}

			fmt.Println(mapparts, mapOccurences, currSl, pt1, pt2)

			if pt1 {
				passwords = append(passwords, curr)
			}

			if pt2 {
				passwordsPartTwo = append(passwordsPartTwo, curr)
			}

		}

	}

	fmt.Println(cnt)

	return passwords, passwordsPartTwo
}

func getInputsByLine() []string {
	inputFile, err := os.Open("./inputs.txt")
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
