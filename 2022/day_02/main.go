package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs := getInputsByLine()
	var oppPointsP1, myPointsP1 int
	var oppPointsP2, myPointsP2 int

	for _, v := range inputs {
		topp, tmy := gameStepPone(v)
		oppPointsP1 += topp
		myPointsP1 += tmy
		topp, tmy = gameStepPtwo(v)
		oppPointsP2 += topp
		myPointsP2 += tmy
	}

	fmt.Println("PART ONE ", myPointsP1)
	fmt.Println("PART TWO ", myPointsP2)
}

func gameStepPone(play []string) (int, int) {
	oppPoints := 0
	myPoints := 0

	pointMap := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
	winnerMap := map[string]string{"A": "Z", "B": "X", "C": "Y", "X": "C", "Y": "A", "Z": "B"}

	oppPoints += pointMap[play[0]]
	myPoints += pointMap[play[1]]

	if winnerMap[play[1]] == play[0] {
		myPoints += 6
	} else {
		if winnerMap[play[0]] == play[1] {
			oppPoints += 6
		} else {
			myPoints += 3
			oppPoints += 3
		}
	}

	return oppPoints, myPoints
}

func gameStepPtwo(play []string) (int, int) {
	oppPoints := 0
	myPoints := 0

	pointMap := map[string]int{"A": 1, "B": 2, "C": 3}
	winnerMap := map[string]string{"A": "C", "B": "A", "C": "B"}
	loserMap := map[string]string{"A": "B", "B": "C", "C": "A"}

	var myPlay string

	switch play[1] {
	case "X": //lose
		myPlay = winnerMap[play[0]]
	case "Y": //draw
		myPlay = play[0]
	default:
		myPlay = loserMap[play[0]]
	}

	oppPoints += pointMap[play[0]]
	myPoints += pointMap[myPlay]

	if winnerMap[myPlay] == play[0] {
		myPoints += 6
	} else {
		if winnerMap[play[0]] == myPlay {
			oppPoints += 6
		} else {
			myPoints += 3
			oppPoints += 3
		}
	}

	return oppPoints, myPoints
}

func getInputsByLine() [][]string {
	inputFile, err := os.Open("./inputs.txt")
	if err != nil {
		fmt.Println("could not find file")
		os.Exit(1)
	}

	defer inputFile.Close()

	var inputs [][]string

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {

		inputs = append(inputs, strings.Split(scanner.Text(), " "))
	}

	return inputs
}
