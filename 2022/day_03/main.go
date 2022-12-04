package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputs, lines := getInputsByLine()
	prioSum := 0
	for _, v := range inputs {
		commonRunes := make(map[rune]bool)

		for i := 0; i < len(v[0]); i++ {
			for _, r := range v[1] {
				if r == rune(v[0][i]) {
					commonRunes[r] = true
				}
			}
			for _, r := range v[0] {
				if r == rune(v[1][i]) {
					commonRunes[r] = true
				}
			}
		}

		for c, _ := range commonRunes {
			prioSum += getPrio(c)
		}
	}

	prioSumP2 := 0
	for i := 0; i < len(lines)/3; i++ {
		firstElemList := getCommonalities(lines[i*3], lines[i*3+1])
		var cs []rune
		for c, _ := range firstElemList {
			cs = append(cs, c)
		}
		secondElemList := getCommonalities(string(cs), lines[i*3+2])
		for c, _ := range secondElemList {
			prioSumP2 += getPrio(c)
		}
	}

	fmt.Println("PART ONE : ", prioSum)
	fmt.Println("PART TWO : ", prioSumP2)
}

func getCommonalities(s1 string, s2 string) map[rune]bool {
	commonRunes := make(map[rune]bool)

	for i := 0; i < len(s1); i++ {
		for _, r := range s2 {
			if r == rune(s1[i]) {
				commonRunes[r] = true
			}
		}
	}

	for i := 0; i < len(s2); i++ {
		for _, r := range s1 {
			if r == rune(s2[i]) {
				commonRunes[r] = true
			}
		}
	}

	return commonRunes
}

func getPrio(char rune) int {
	if int(char) >= 65 && int(char) <= 90 {
		//A-Z
		return int(char) - 38
	}

	return int(char) - 96
}

func getInputsByLine() ([][2]string, []string) {
	inputFile, err := os.Open("./inputs.txt")
	if err != nil {
		fmt.Println("could not find file")
		os.Exit(1)
	}

	defer inputFile.Close()

	var inputs [][2]string
	var inputLines []string

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		input := scanner.Text()
		inputLines = append(inputLines, input)
		inputs = append(inputs, [2]string{input[:(len(input) / 2)], input[len(input)/2:]})
	}

	return inputs, inputLines
}
