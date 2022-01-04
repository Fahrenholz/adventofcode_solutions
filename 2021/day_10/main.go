package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	complements      = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>', ')': '(', ']': '[', '}': '{', '>': '<'}
	syntaxScoreTable = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	complScoreTable  = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
)

func main() {
	inputs := getInputsByLine()
	syntaxScore := 0
	var complScores []int

	for _, v := range inputs {
		sc, compl := evalLines(v)
		syntaxScore += sc
		if compl != "" {
			complScores = append(complScores, calcCompletionScore(compl))
		}
	}

	complScore := findMiddleScore(complScores)

	fmt.Printf("Total syntax score: %d\n", syntaxScore)
	fmt.Printf("Total completion score: %d\n", complScore)

}

func findMiddleScore(sc []int) int {
	sort.Ints(sc)
	middle := len(sc) / 2

	return sc[middle]
}

func evalLines(inp string) (int, string) {
	var stack []rune
	var last rune

	for _, v := range inp {

		if v == '(' || v == '[' || v == '{' || v == '<' {
			stack = append(stack, v)
			continue
		}

		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if complements[v] != last {
			return syntaxScoreTable[v], ""
		}
	}

	compl := ""

	for _, stv := range stack {
		compl = fmt.Sprintf("%s%s", string(complements[stv]), compl)
	}

	return 0, compl
}

func calcCompletionScore(str string) int {
	sc := 0
	for _, v := range str {
		sc = (sc * 5) + complScoreTable[v]
	}

	return sc
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
