package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type passwordPolicy struct {
	lower  int
	upper  int
	letter rune
}

type passwordLine struct {
	policy   passwordPolicy
	password string
}

func (l passwordLine) isValidSledgeRentalPassword() bool {
	occurences := 0
	for _, v := range l.password {
		if v == l.policy.letter {
			occurences++
		}
	}

	return occurences >= l.policy.lower && occurences <= l.policy.upper
}

func (l passwordLine) isValidTobboganRentalPassword() bool {
	return (rune(l.password[l.policy.lower-1]) == l.policy.letter && rune(l.password[l.policy.upper-1]) != l.policy.letter) || (rune(l.password[l.policy.lower-1]) != l.policy.letter && rune(l.password[l.policy.upper-1]) == l.policy.letter)
}

func main() {
	inputs := parseInputs(getInputsByLine())
	validSledgeRentalPasswords := 0
	validTobboganRentalPasswords := 0

	for _, v := range inputs {
		if v.isValidSledgeRentalPassword() {
			validSledgeRentalPasswords++
		}
		if v.isValidTobboganRentalPassword() {
			validTobboganRentalPasswords++
		}
	}

	fmt.Printf("Solution 1: %d valid passwords.\n", validSledgeRentalPasswords)
	fmt.Printf("Solution 2: %d valid passwords.\n", validTobboganRentalPasswords)
}

func parseInputs(inp []string) []passwordLine {
	re := regexp.MustCompile(`(\d+)-(\d+) ([a-zA-Z]): (\w+)`)
	var res []passwordLine

	for _, li := range inp {
		matches := re.FindStringSubmatch(li)
		lo, _ := strconv.Atoi(matches[1])
		up, _ := strconv.Atoi(matches[2])
		res = append(res, passwordLine{policy: passwordPolicy{lower: lo, upper: up, letter: rune(matches[3][0])}, password: matches[4]})
	}

	return res
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
