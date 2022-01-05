package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passportData struct {
	BirthYear      string `json:"byr"`
	IssueYear      string `json:"iyr"`
	ExpirationYear string `json:"eyr"`
	Height         string `json:"hgt"`
	HairColor      string `json:"hcl"`
	EyeColor       string `json:"ecl"`
	PassportID     string `json:"pid"`
	CountryID      string `json:"cid"`
}

func (p passportData) isValidBirthYear() bool {
	y, err := strconv.Atoi(p.BirthYear)

	return err == nil && 1920 <= y && y <= 2002
}

func (p passportData) isValidIssueYear() bool {
	y, err := strconv.Atoi(p.IssueYear)

	return err == nil && 2010 <= y && y <= 2020
}

func (p passportData) isValidExpirationYear() bool {
	y, err := strconv.Atoi(p.ExpirationYear)

	return err == nil && 2020 <= y && y <= 2030
}

func (p passportData) isValidHeight() bool {

	if len(p.Height) < 4 {
		return false
	}

	unit := p.Height[len(p.Height)-2:]
	measure, err := strconv.Atoi(p.Height[0 : len(p.Height)-2])
	if err != nil {
		return false
	}

	switch unit {
	case "cm":
		return 150 <= measure && measure <= 193
	case "in":
		return 59 <= measure && measure <= 76
	default:
		return false
	}
}

func (p passportData) isValidHairColor() bool {
	re := regexp.MustCompile(`#[0-9a-f]{6}`)
	return re.MatchString(p.HairColor)
}

func (p passportData) isValidEyeColor() bool {
	vals := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}

	_, res := vals[p.EyeColor]

	return res
}

func (p passportData) isValidPassportID() bool {
	_, err := strconv.Atoi(p.PassportID)

	return err == nil && len(p.PassportID) == 9
}

func main() {
	inputs := parseInputs(getInputsByLine(os.Args[1]))
	solvePartOne(inputs)
	solvePartTwo(inputs)
}

func parseInputs(inputs []string) []passportData {
	passport := make(map[string]string)
	var result []map[string]string
	for _, v := range inputs {
		if strings.TrimSpace(v) == "" {
			result = append(result, passport)
			passport = make(map[string]string)
			continue
		}

		vals := strings.Split(v, " ")
		for _, val := range vals {
			information := strings.Split(val, ":")
			passport[information[0]] = information[1]
		}
	}

	result = append(result, passport)

	res, _ := json.Marshal(result)

	var marshalledRes []passportData

	_ = json.Unmarshal(res, &marshalledRes)

	return marshalledRes
}

func solvePartOne(inp []passportData) {
	valid := 0
	for _, set := range inp {
		if set.BirthYear != "" && set.IssueYear != "" && set.ExpirationYear != "" && set.Height != "" && set.HairColor != "" && set.EyeColor != "" && set.PassportID != "" {
			valid++
		}
	}

	fmt.Printf("Solution 1: %d valid passports out of %d.\n", valid, len(inp))
}

func solvePartTwo(inp []passportData) {
	valid := 0
	for _, set := range inp {
		if set.isValidBirthYear() && set.isValidIssueYear() && set.isValidExpirationYear() && set.isValidHeight() && set.isValidHairColor() && set.isValidEyeColor() && set.isValidPassportID() {
			valid++
		}
	}

	fmt.Printf("Solution 2: %d valid passports out of %d\n", valid, len(inp))
}

func getInputsByLine(file string) []string {
	inputFile, err := os.Open(file)
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
