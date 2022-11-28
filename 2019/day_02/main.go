package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partTwo() {

	var noun, verb int

search:
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			if runProgram(noun, verb) == 19690720 {
				break search
			}
		}
	}

	fmt.Printf("Part 2: Noun = %d, Verb = %d, result = %d\n", noun, verb, 100*noun+verb)
}

func runProgram(noun, verb int) int {
	// instantiate memory
	memory := getInputsByLine()

	// init parameters
	memory[1] = noun
	memory[2] = verb

	instructionPointer := 0

	//exec
programmExec:
	for instructionPointer < len(memory)-1 {
		switch memory[instructionPointer] {
		case 1: //addition
			memory[memory[instructionPointer+3]] = memory[memory[instructionPointer+1]] + memory[memory[instructionPointer+2]]
			instructionPointer += 4
		case 2: //multiplication
			memory[memory[instructionPointer+3]] = memory[memory[instructionPointer+1]] * memory[memory[instructionPointer+2]]
			instructionPointer += 4
		case 99: //halt
			instructionPointer += -1
			break programmExec
		default:
			panic("unknown opcode")
		}
	}

	return memory[0]
}

func partOne() {
	inputs := getInputsByLine()
	i := 0

	inputs[1] = 12
	inputs[2] = 2

programmExec:
	for i < len(inputs)-1 {
		switch inputs[i] {
		case 1:
			inputs[inputs[i+3]] = inputs[inputs[i+1]] + inputs[inputs[i+2]]
			i += 4
		case 2:
			inputs[inputs[i+3]] = inputs[inputs[i+1]] * inputs[inputs[i+2]]
			i += 4
		case 99:
			break programmExec
		default:
			panic("unknown opcode")
		}
	}

	fmt.Printf("Part 1: %d\n", inputs[0])
}

func getInputsByLine() []int {
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

	inputsExploded := strings.Split(inputs[0], ",")

	inputsInt := make([]int, len(inputsExploded))

	for i, v := range inputsExploded {
		inputsInt[i], _ = strconv.Atoi(v)
	}

	return inputsInt
}
