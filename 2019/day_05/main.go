package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	runProgram(-1, -1, true)
}

func runProgram(noun, verb int, debug bool) int {
	// instantiate memory
	memory := getInputsByLine(os.Args[1])

	// init parameters
	if noun != -1 {
		memory[1] = noun
	}

	if noun != -1 {
		memory[2] = verb
	}

	instructionPointer := 0

	//exec
programmExec:
	for instructionPointer < len(memory)-1 {
		opcode, modeC, modeB, _ := getParameterModes(memory[instructionPointer])

		switch opcode {
		case 1: //addition
			termA := memory[instructionPointer+1]
			if modeC == 0 {
				termA = memory[memory[instructionPointer+1]]
			}

			termB := memory[instructionPointer+2]
			if modeB == 0 {
				termB = memory[memory[instructionPointer+2]]
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], ", ", memory[instructionPointer+2], ", ", memory[instructionPointer+3], " => Address ", memory[instructionPointer+3], " = ", termA, " + ", termB, " = ", termA+termB)
			}

			memory[memory[instructionPointer+3]] = termA + termB
			instructionPointer += 4
		case 2: //multiplication
			termA := memory[instructionPointer+1]
			if modeC == 0 {
				termA = memory[memory[instructionPointer+1]]
			}

			termB := memory[instructionPointer+2]
			if modeB == 0 {
				termB = memory[memory[instructionPointer+2]]
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], ", ", memory[instructionPointer+2], ", ", memory[instructionPointer+3], " => Address ", memory[instructionPointer+3], " = ", termA, " * ", termB, " = ", termA*termB)
			}

			memory[memory[instructionPointer+3]] = termA * termB
			instructionPointer += 4
		case 3: //input
			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], " => Input to ", memory[instructionPointer+1])
			}
			var i int
			fmt.Printf("Input asked: ")
			_, _ = fmt.Scanf("%d", &i)

			memory[memory[instructionPointer+1]] = i
			instructionPointer += 2

		case 4: //input
			termA := memory[instructionPointer+1]
			if modeC == 0 {
				termA = memory[memory[instructionPointer+1]]
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], " => Display ", termA)
			}
			fmt.Println("Output: ", termA)
			//nextOpcode, _, _, _ := getParameterModes(memory[instructionPointer+2])
			//if termA != 0 && nextOpcode != 99 {
			//	panic("Wrong output")
			//}
			instructionPointer += 2

		case 5: // jump-if-true
			termA := memory[instructionPointer+1]
			if modeC == 0 {
				termA = memory[memory[instructionPointer+1]]
			}

			termB := memory[instructionPointer+2]
			if modeB == 0 {
				termB = memory[memory[instructionPointer+2]]
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], ", ", memory[instructionPointer+2], " => Jump to ", termB, " if ", termA, " != 0")
			}

			if termA != 0 {
				instructionPointer = termB
			} else {
				instructionPointer += 3
			}
		case 6: // jump-if-false
			termA := memory[instructionPointer+1]
			if modeC == 0 {
				termA = memory[memory[instructionPointer+1]]
			}

			termB := memory[instructionPointer+2]
			if modeB == 0 {
				termB = memory[memory[instructionPointer+2]]
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], ", ", memory[instructionPointer+2], " => Jump to ", termB, " if ", termA, " == 0")
			}

			if termA == 0 {
				instructionPointer = termB
			} else {
				instructionPointer += 3
			}
		case 7: // less-than
			termA := memory[instructionPointer+1]
			if modeC == 0 {
				termA = memory[memory[instructionPointer+1]]
			}

			termB := memory[instructionPointer+2]
			if modeB == 0 {
				termB = memory[memory[instructionPointer+2]]
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], ", ", memory[instructionPointer+2], ", ", memory[instructionPointer+3], " => ", memory[instructionPointer+3], " = 1", " if ", termA, " < ", termB, ", otherwise 0")
			}

			if termA < termB {
				memory[memory[instructionPointer+3]] = 1
			} else {
				memory[memory[instructionPointer+3]] = 0
			}

			instructionPointer += 4
		case 8: // equal
			termA := memory[instructionPointer+1]
			if modeC == 0 {
				termA = memory[memory[instructionPointer+1]]
			}

			termB := memory[instructionPointer+2]
			if modeB == 0 {
				termB = memory[memory[instructionPointer+2]]
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], ", ", memory[instructionPointer+2], ", ", memory[instructionPointer+3], " => ", memory[instructionPointer+3], " = 1", " if ", termA, " == ", termB, ", otherwise 0")
			}

			if termA == termB {
				memory[memory[instructionPointer+3]] = 1
			} else {
				memory[memory[instructionPointer+3]] = 0
			}

			instructionPointer += 4
		case 99: //halt
			instructionPointer += 1
			break programmExec
		default:
			panic("unknown opcode")
		}
	}

	return memory[0]
}

func getParameterModes(modeDescriptor int) (int, int, int, int) {
	modeA := modeDescriptor / 10000
	modeB := (modeDescriptor - modeA*10000) / 1000
	modeC := (modeDescriptor - modeA*10000 - modeB*1000) / 100
	opCode := modeDescriptor - modeA*10000 - modeB*1000 - modeC*100

	return opCode, modeC, modeB, modeA
}

func getInputsByLine(filename string) []int {
	inputFile, err := os.Open(fmt.Sprintf("./%s.txt", filename))
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
