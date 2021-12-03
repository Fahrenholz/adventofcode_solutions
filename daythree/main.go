package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type bitCount struct {
	zeros int
	ones  int
}

func main() {
	inputs := getInputsByLine()

	cnt := getCnt(inputs)

	var gammaRate, epsilonRate string

	for _, v := range cnt {
		switch {
		case v.zeros < v.ones:
			gammaRate = fmt.Sprintf("%s1", gammaRate)
			epsilonRate = fmt.Sprintf("%s0", epsilonRate)
		default:
			gammaRate = fmt.Sprintf("%s0", gammaRate)
			epsilonRate = fmt.Sprintf("%s1", epsilonRate)
		}
	}

	gRate, _ := strconv.ParseInt(gammaRate, 2, 32)
	eRate, _ := strconv.ParseInt(epsilonRate, 2, 32)

	fmt.Printf("Gamma rate: %d, epsilon rate : %d, power consumption: %d\n", gRate, eRate, gRate*eRate)

	oxyRate, _ := strconv.ParseInt(recurseOxy(inputs, 0), 2, 64)
	co2Rate, _ := strconv.ParseInt(recurseCo2(inputs, 0), 2, 64)
	fmt.Printf("Oxygen generator rating: %d, CO2 scrubber rating: %d, life support rating :%d\n", oxyRate, co2Rate, oxyRate*co2Rate)
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

func getCnt(inputs []string) []bitCount {
	cnt := make([]bitCount, 12)
	for i, _ := range cnt {
		cnt[i] = bitCount{}
	}

	for _, inp := range inputs {
		for i, v := range inp {

			val, _ := strconv.ParseInt(string(rune(v)), 10, 64)
			switch {
			case val == 0:
				cnt[i].zeros++
			case val == 1:
				cnt[i].ones++
			}
		}
	}

	return cnt
}

func recurseOxy(inputs []string, idx int) string {
	if len(inputs) == 1 {
		return inputs[0]
	}

	if len(inputs) == 0 || idx == 12 {
		fmt.Println(len(inputs), idx)
		panic("no inputs")
	}

	cnt := getCnt(inputs)

	var tmp []string

	for _, v := range inputs {
		if rune(v[idx]) == '0' && cnt[idx].zeros > cnt[idx].ones {
			tmp = append(tmp, v)
		}
		if rune(v[idx]) == '1' && cnt[idx].ones >= cnt[idx].zeros {
			tmp = append(tmp, v)
		}
	}

	return recurseOxy(tmp, idx+1)
}

func recurseCo2(inputs []string, idx int) string {
	if len(inputs) == 1 {
		return inputs[0]
	}

	if len(inputs) == 0 || idx == 12 {
		panic("no inputs")
	}

	cnt := getCnt(inputs)

	var tmp []string

	for _, v := range inputs {
		if rune(v[idx]) == '0' && cnt[idx].zeros <= cnt[idx].ones {
			tmp = append(tmp, v)
		}
		if rune(v[idx]) == '1' && cnt[idx].ones < cnt[idx].zeros {
			tmp = append(tmp, v)
		}
	}

	return recurseCo2(tmp, idx+1)
}
