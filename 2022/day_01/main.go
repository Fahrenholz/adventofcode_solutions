package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	totalCalories int
	caloryList    []int
}

func main() {
	inputs := getInputsByElf()
	sort.Slice(inputs, func(i, j int) bool {
		return inputs[i].totalCalories < inputs[j].totalCalories
	})

	fmt.Println("PART ONE: ", inputs[len(inputs)-1].totalCalories)
	fmt.Println("PART TWO: ", inputs[len(inputs)-1].totalCalories+inputs[len(inputs)-2].totalCalories+inputs[len(inputs)-3].totalCalories)
}

func getInputsByElf() []elf {
	inputFile, err := os.Open("./inputs.txt")
	if err != nil {
		fmt.Println("could not find file")
		os.Exit(1)
	}

	defer inputFile.Close()

	var inputs []elf

	scanner := bufio.NewScanner(inputFile)

	tmp := elf{}
	for scanner.Scan() {
		txt := scanner.Text()

		if txt == "" {
			inputs = append(inputs, tmp)
			tmp = elf{}
		} else {
			nb, _ := strconv.Atoi(txt)
			tmp.totalCalories += nb
			tmp.caloryList = append(tmp.caloryList, nb)
		}
	}

	return inputs
}
