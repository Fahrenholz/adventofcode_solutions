package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type entry struct {
	uniqueSignalPatterns []string
	digits               []string
}

var uniqueNumberSegmentMap = map[int]int{2: 1, 4: 4, 3: 7, 7: 8}

func main() {
	inputs := parseInputs(getInputsByLine())
	numberEasyDigits := 0
	for _, v := range inputs {
		for _, vv := range v.digits {
			numberEasyDigits += determineEasyDigits(vv)
		}
	}

	fmt.Printf("Solution 1 : %d digits in (1,4,7,8)\n", numberEasyDigits)

	sumDigits := 0

	for _, v := range inputs {
		sumDigits += decodeEntry(v)
	}

	fmt.Printf("Solution 2 : %d total\n", sumDigits)
}

func decodeEntry(inp entry) int {
	allValues := append(inp.uniqueSignalPatterns, inp.digits...)

	signalPatternVal := make([]string, len(allValues))
	cablePatternVal := make(map[int]string)

	//decode easy inputs
	for i, v := range allValues {
		if sv, ok := uniqueNumberSegmentMap[len(v)]; ok {
			cablePatternVal[sv] = v
			signalPatternVal[i] = fmt.Sprintf("%d", sv)
		}
	}

	for i, v := range allValues {
		if len(v) == 6 { //0, 6 or 9
			if !contains(v, cablePatternVal[1]) {
				signalPatternVal[i] = "6"
				cablePatternVal[6] = v

				continue
			}

			//0 or 9
			if !contains(v, cablePatternVal[4]) { //0
				signalPatternVal[i] = "0"
				cablePatternVal[0] = v

				continue
			}

			signalPatternVal[i] = "9"
			cablePatternVal[9] = v

			continue
		}
	}

	for i, v := range allValues {
		if len(v) == 5 { //2 or 5

			if contains(cablePatternVal[6], v) {
				signalPatternVal[i] = "5"
				cablePatternVal[5] = v

				continue
			}

			if contains(v, cablePatternVal[1]) {
				signalPatternVal[i] = "3"
				cablePatternVal[3] = v

				continue
			}

			signalPatternVal[i] = "2"
			cablePatternVal[2] = v
		}
	}

	digits, _ := strconv.Atoi(strings.Join(signalPatternVal[len(allValues)-4:len(allValues)], ""))

	return digits
}

func contains(str, sub string) bool {
	for _, v := range sub {
		if !strings.ContainsRune(str, v) {
			return false
		}
	}

	return true
}

func determineEasyDigits(str string) int {
	if _, ok := uniqueNumberSegmentMap[len(str)]; ok {
		return 1
	}

	return 0
}

func parseInputs(inputs []string) []entry {
	var res []entry
	for _, v := range inputs {
		split := strings.Split(v, " | ")
		res = append(res, entry{
			uniqueSignalPatterns: strings.Split(split[0], " "),
			digits:               strings.Split(split[1], " "),
		})
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
