package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs := getInputsByLine()

	draft, boards := parseInputs(inputs)

	var winningScore int
	lowerRound := len(draft) + 1
	var upperRound int
	var loosingBoardScore int

	for _, b := range boards {
		bingo, round, score := b.playDraft(draft)
		if (bingo && lowerRound > round) || (bingo && lowerRound == round && winningScore < score) {
			lowerRound = round
			winningScore = score
		}

		if (bingo && upperRound < round) || (bingo && upperRound == round && loosingBoardScore > score) {
			loosingBoardScore = score
			upperRound = round
		}
	}

	fmt.Printf("Winning score: %d\n", winningScore)
	fmt.Printf("Loosing score: %d\n", loosingBoardScore)
}

func parseInputs(inputs []string) ([]int, []*bingoBoard) {
	var draft []int
	var buf []string
	var boards []*bingoBoard

	for idx, v := range inputs {
		if idx == 0 {
			//the first line is our number draft
			draft = parseNumberDraft(v)
			continue
		}

		if strings.TrimSpace(v) == "" && len(buf) != 0 {
			//empty line is new board, we create the previous one from the buffer and reset the buffer
			boards = append(boards, NewBoard(buf))
			buf = []string{}
			continue
		}

		if strings.TrimSpace(v) == "" {
			continue
		}

		buf = append(buf, v)
	}
	return draft, boards
}

func parseNumberDraft(str string) []int {
	strJSON := fmt.Sprintf("[%s]", str)
	var res []int

	_ = json.Unmarshal([]byte(strJSON), &res)

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
