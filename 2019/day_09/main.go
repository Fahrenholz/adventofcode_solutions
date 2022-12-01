package main

import "fmt"

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	chIn := make(chan int, 0)
	chOut := make(chan int, 100)

	runProgram(2, chIn, chOut, nil, false)

	close(chOut)
	close(chIn)

	for output := range chOut {
		fmt.Println("Output PART TWO: ", output)
	}
}

func partOne() {
	chIn := make(chan int, 0)
	chOut := make(chan int, 100)

	runProgram(1, chIn, chOut, nil, false)

	close(chOut)
	close(chIn)

	for output := range chOut {
		fmt.Println("Output PART ONE: ", output)
	}
}
