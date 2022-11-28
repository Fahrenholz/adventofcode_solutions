package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	max := 0
	var maxPhaseNb [5]int

	var phaseSignals [][5]int

	for phase1 := 5; phase1 < 10; phase1++ {
		for phase2 := 5; phase2 < 10; phase2++ {
			for phase3 := 5; phase3 < 10; phase3++ {
				for phase4 := 5; phase4 < 10; phase4++ {
					for phase5 := 5; phase5 < 10; phase5++ {
						occs := make(map[int]int)

						occs[phase1]++
						occs[phase2]++
						occs[phase3]++
						occs[phase4]++
						occs[phase5]++

						accepted := true

						for ph, _ := range occs {
							if occs[ph] != 1 {
								accepted = false
							}
						}

						if accepted {
							phaseSignals = append(phaseSignals, [5]int{phase1, phase2, phase3, phase4, phase5})
						}
					}
				}
			}
		}
	}

	for _, sig := range phaseSignals {
		chIn1 := make(chan int, 1)
		chIn2 := make(chan int, 1)
		chIn3 := make(chan int, 1)
		chIn4 := make(chan int, 1)
		chIn5 := make(chan int, 1)
		chExit := make(chan int, 100)

		chIn1 <- 0
		var wg sync.WaitGroup

		wg.Add(5)
		go func() { runProgram(sig[0], chIn1, chIn2, nil, false); wg.Done() }()
		go func() { runProgram(sig[1], chIn2, chIn3, nil, false); wg.Done() }()
		go func() { runProgram(sig[2], chIn3, chIn4, nil, false); wg.Done() }()
		go func() { runProgram(sig[3], chIn4, chIn5, nil, false); wg.Done() }()
		go func() { runProgram(sig[4], chIn5, chIn1, chExit, false); wg.Done() }()

		wg.Wait()

		close(chExit)

		var op5 int

		for top5 := range chExit {
			op5 = top5
		}

		if max < op5 {
			max = op5
			maxPhaseNb = sig
		}
	}

	fmt.Println("PART ONE || Max: ", max, ", Phases: ", maxPhaseNb)
}

func partOne() {
	max := 0
	var maxPhaseNb [5]int

	var phaseSignals [][5]int

	for phase1 := 0; phase1 < 5; phase1++ {
		for phase2 := 0; phase2 < 5; phase2++ {
			for phase3 := 0; phase3 < 5; phase3++ {
				for phase4 := 0; phase4 < 5; phase4++ {
					for phase5 := 0; phase5 < 5; phase5++ {
						occs := make(map[int]int)

						occs[phase1]++
						occs[phase2]++
						occs[phase3]++
						occs[phase4]++
						occs[phase5]++

						accepted := true

						for ph, _ := range occs {
							if occs[ph] != 1 {
								accepted = false
							}
						}

						if accepted {
							phaseSignals = append(phaseSignals, [5]int{phase1, phase2, phase3, phase4, phase5})
						}
					}
				}
			}
		}
	}

	for _, sig := range phaseSignals {
		chIn1 := make(chan int, 1)
		chIn2 := make(chan int, 1)
		chIn3 := make(chan int, 1)
		chIn4 := make(chan int, 1)
		chIn5 := make(chan int, 1)

		chIn1 <- 0
		fmt.Println("A1")
		runProgram(sig[0], chIn1, chIn2, nil, true)
		fmt.Println("A2")
		runProgram(sig[1], chIn2, chIn3, nil, true)
		fmt.Println("A3")
		runProgram(sig[2], chIn3, chIn4, nil, true)
		fmt.Println("A4")
		runProgram(sig[3], chIn4, chIn5, nil, true)
		fmt.Println("A5")
		runProgram(sig[4], chIn5, chIn1, nil, true)

		var op5 int

		op5 = <-chIn1

		if max < op5 {
			max = op5
			maxPhaseNb = sig
		}
	}

	fmt.Println("PART ONE || Max: ", max, ", Phases: ", maxPhaseNb)
}

func runProgram(phase int, chIn, chOut, chExit chan int, debug bool) {
	// instantiate memory
	memory := getInputsByLine(os.Args[1])
	phaseDone := false

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

			var i int

			if !phaseDone {
				i = phase
				phaseDone = true
			} else {
				i = <-chIn
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], " => Input ", i, " to ", memory[instructionPointer+1])
			}

			memory[memory[instructionPointer+1]] = i
			instructionPointer += 2

		case 4: //output
			termA := memory[instructionPointer+1]
			if modeC == 0 {
				termA = memory[memory[instructionPointer+1]]
			}

			if debug {
				fmt.Println("Pos ", instructionPointer, ": ", memory[instructionPointer], ", ", memory[instructionPointer+1], " => Display ", termA)
			}

			chOut <- termA
			if chExit != nil {
				chExit <- termA
			}
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
