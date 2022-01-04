package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	template, rules := parseInputs(getInputsByLine("./inputs.txt"))
	pairs := templateToPairs(template)

	var max, min int

	for i := 1; i <= 40; i++ {
		newPairs := make(map[string]int)
		for p := range pairs {
			nc := rules[p]
			newPairs[fmt.Sprintf("%c%s", rune(p[0]), nc)] += pairs[p]
			newPairs[fmt.Sprintf("%s%c", nc, rune(p[1]))] += pairs[p]
		}
		pairs = newPairs
		max, min = countPairs(pairs, rune(template[len(template)-1]))
		if i == 10 {
			fmt.Printf("Solution 1 : %d - %d = %d\n", max, min, max-min)
		}
		if i == 40 {
			fmt.Printf("Solution 2 : %d - %d = %d\n", max, min, max-min)
		}
	}
}

func countPairs(pairs map[string]int, last rune) (int, int) {
	cnt := make(map[rune]int)

	for p := range pairs {
		cnt[rune(p[0])] += pairs[p]
	}
	cnt[last] += 1

	max := 0
	min := -1
	for v := range cnt {
		if cnt[v] > max {
			max = cnt[v]
		}
		if cnt[v] < min || min == -1 {
			min = cnt[v]
		}
	}

	return max, min
}

func templateToPairs(tmpl string) map[string]int {
	pairs := make(map[string]int)
	for i := 0; i < len(tmpl)-1; i++ {
		if _, ok := pairs[fmt.Sprintf("%c%c", tmpl[i], tmpl[i+1])]; !ok {
			pairs[fmt.Sprintf("%c%c", tmpl[i], tmpl[i+1])] = 0
		}

		pairs[fmt.Sprintf("%c%c", tmpl[i], tmpl[i+1])]++
	}
	return pairs
}

func parseInputs(inputs []string) (string, map[string]string) {
	template := inputs[0]
	pairRules := make(map[string]string)
	for i := 2; i < len(inputs); i++ {
		rule := strings.Split(inputs[i], " -> ")
		pairRules[rule[0]] = rule[1]
	}

	return template, pairRules
}

func getInputsByLine(fileN string) []string {
	inputFile, err := os.Open(fileN)
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
