package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func runProgram(phase int, chIn, chOut, chExit chan int, debug bool) {
	// instantiate memory
	memory := getProgrammCode(os.Args[1])
	phaseDone := false

	instructionPointer := 0
	relativeBase := 0

	//exec
programmExec:
	for instructionPointer < len(memory)-1 {
		opcode, modeC, modeB, modeA := getParameterModes(memory[instructionPointer])

		switch opcode {
		case 1: //addition
			termA := memory[instructionPointer+1]
			if modeC == 0 {
				termA = memory[memory[instructionPointer+1]]
			}
			if modeC == 2 {
				termA = memory[memory[instructionPointer+1]+relativeBase]
			}

			termB := memory[instructionPointer+2]
			if modeB == 0 {
				termB = memory[memory[instructionPointer+2]]
			}
			if modeB == 2 {
				termB = memory[memory[instructionPointer+2]+relativeBase]
			}

			resultMemorySlot := memory[instructionPointer+3]
			if modeA == 2 {
				resultMemorySlot += relativeBase
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], ", ", memory[instructionPointer+2], ", ", memory[instructionPointer+3], " => Address ", resultMemorySlot, " = ", termA, " + ", termB, " = ", termA+termB)
			}

			memory[resultMemorySlot] = termA + termB
			instructionPointer += 4
		case 2: //multiplication
			termA := memory[instructionPointer+1]
			if modeC == 0 {
				termA = memory[memory[instructionPointer+1]]
			}
			if modeC == 2 {
				termA = memory[memory[instructionPointer+1]+relativeBase]
			}

			termB := memory[instructionPointer+2]
			if modeB == 0 {
				termB = memory[memory[instructionPointer+2]]
			}
			if modeB == 2 {
				termB = memory[memory[instructionPointer+2]+relativeBase]
			}

			resultMemorySlot := memory[instructionPointer+3]
			if modeA == 2 {
				resultMemorySlot += relativeBase
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], ", ", memory[instructionPointer+2], ", ", memory[instructionPointer+3], " => Address ", resultMemorySlot, " = ", termA, " * ", termB, " = ", termA*termB)
			}

			memory[resultMemorySlot] = termA * termB
			instructionPointer += 4
		case 3: //input

			var valueToInput int

			if !phaseDone {
				valueToInput = phase
				phaseDone = true
			} else {
				valueToInput = <-chIn
			}

			resultMemorySlot := memory[instructionPointer+1]
			if modeC == 2 {
				resultMemorySlot += relativeBase
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], " => Input ", valueToInput, " to ", resultMemorySlot)
			}

			memory[resultMemorySlot] = valueToInput
			instructionPointer += 2

		case 4: //output
			valueToDisplay := memory[instructionPointer+1]
			if modeC == 0 {
				valueToDisplay = memory[memory[instructionPointer+1]]
			}
			if modeC == 2 {
				valueToDisplay = memory[memory[instructionPointer+1]+relativeBase]
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], " => Display ", valueToDisplay)
			}

			chOut <- valueToDisplay
			if chExit != nil {
				chExit <- valueToDisplay
			}
			instructionPointer += 2

		case 5: // jump-if-true
			testedVal := memory[instructionPointer+1]
			if modeC == 0 {
				testedVal = memory[memory[instructionPointer+1]]
			}
			if modeC == 2 {
				testedVal = memory[memory[instructionPointer+1]+relativeBase]
			}

			jumpTarget := memory[instructionPointer+2]
			if modeB == 0 {
				jumpTarget = memory[memory[instructionPointer+2]]
			}
			if modeB == 2 {
				jumpTarget = memory[memory[instructionPointer+2]+relativeBase]
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], ", ", memory[instructionPointer+2], " => Jump to ", jumpTarget, " if ", testedVal, " != 0")
			}

			if testedVal != 0 {
				instructionPointer = jumpTarget
			} else {
				instructionPointer += 3
			}
		case 6: // jump-if-false
			testedVal := memory[instructionPointer+1]
			if modeC == 0 {
				testedVal = memory[memory[instructionPointer+1]]
			}
			if modeC == 2 {
				testedVal = memory[memory[instructionPointer+1]+relativeBase]
			}

			jumpTarget := memory[instructionPointer+2]
			if modeB == 0 {
				jumpTarget = memory[memory[instructionPointer+2]]
			}
			if modeB == 2 {
				jumpTarget = memory[memory[instructionPointer+2]+relativeBase]
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], ", ", memory[instructionPointer+2], " => Jump to ", jumpTarget, " if ", testedVal, " == 0")
			}

			if testedVal == 0 {
				instructionPointer = jumpTarget
			} else {
				instructionPointer += 3
			}
		case 7: // less-than
			termA := memory[instructionPointer+1]
			if modeC == 0 {
				termA = memory[memory[instructionPointer+1]]
			}
			if modeC == 2 {
				termA = memory[memory[instructionPointer+1]+relativeBase]
			}

			termB := memory[instructionPointer+2]
			if modeB == 0 {
				termB = memory[memory[instructionPointer+2]]
			}
			if modeB == 2 {
				termB = memory[memory[instructionPointer+2]+relativeBase]
			}

			resultMemorySlot := memory[instructionPointer+3]
			if modeA == 2 {
				resultMemorySlot += relativeBase
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], ", ", memory[instructionPointer+2], ", ", memory[instructionPointer+3], " => ", resultMemorySlot, " = 1", " if ", termA, " < ", termB, ", otherwise 0")
			}

			if termA < termB {
				memory[resultMemorySlot] = 1
			} else {
				memory[resultMemorySlot] = 0
			}

			instructionPointer += 4
		case 8: // equal
			termA := memory[instructionPointer+1]
			if modeC == 0 {
				termA = memory[memory[instructionPointer+1]]
			}
			if modeC == 2 {
				termA = memory[memory[instructionPointer+1]+relativeBase]
			}

			termB := memory[instructionPointer+2]
			if modeB == 0 {
				termB = memory[memory[instructionPointer+2]]
			}
			if modeB == 2 {
				termB = memory[memory[instructionPointer+2]+relativeBase]
			}

			resultMemorySlot := memory[instructionPointer+3]
			if modeA == 2 {
				resultMemorySlot += relativeBase
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], ", ", memory[instructionPointer+2], ", ", memory[instructionPointer+3], " => ", resultMemorySlot, " = 1", " if ", termA, " == ", termB, ", otherwise 0")
			}

			if termA == termB {
				memory[resultMemorySlot] = 1
			} else {
				memory[resultMemorySlot] = 0
			}

			instructionPointer += 4
		case 9:
			adjustment := memory[instructionPointer+1]
			if modeC == 0 {
				adjustment = memory[memory[instructionPointer+1]]
			}
			if modeC == 2 {
				adjustment = memory[memory[instructionPointer+1]+relativeBase]
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], " => Adjust relative base by ", adjustment, " from ", relativeBase, " to ", relativeBase+adjustment)
			}

			relativeBase += adjustment
			instructionPointer += 2
		case 99: //halt
			instructionPointer += 1
			break programmExec
		default:
			panic("unknown opcode")
		}
	}
}

func getParameterModes(modeDescriptor int) (int, int, int, int) {
	modeA := modeDescriptor / 10000
	modeB := (modeDescriptor - modeA*10000) / 1000
	modeC := (modeDescriptor - modeA*10000 - modeB*1000) / 100
	opCode := modeDescriptor - modeA*10000 - modeB*1000 - modeC*100

	return opCode, modeC, modeB, modeA
}

func getProgrammCode(filename string) []int {
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

	inputsInt := make([]int, len(inputsExploded)*1000)

	for i, v := range inputsExploded {
		inputsInt[i], _ = strconv.Atoi(v)
	}

	return inputsInt
}
