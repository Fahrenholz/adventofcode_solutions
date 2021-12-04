package main

import (
	"regexp"
	"strconv"
)

type bingoBoard struct {
	rows []row
}

type row struct {
	cells []cell
}

type cell struct {
	value  int
	marked bool
}

func NewBoard(buf []string) *bingoBoard {
	board := bingoBoard{rows: make([]row, 5)}
	reg := regexp.MustCompile(`([0-9]{1,2})`)
	for i, v := range buf {
		matches := reg.FindAllString(v, 5)
		row := row{cells: make([]cell, 5)}
		for j, val := range matches {
			valNb, _ := strconv.Atoi(val)
			row.cells[j] = cell{value: valNb}
		}
		board.rows[i] = row
	}

	return &board
}

func (b *bingoBoard) playRound(draftNb int) bool {
	for i, _ := range b.rows {
		for j, _ := range b.rows[i].cells {
			if b.rows[i].cells[j].value == draftNb {
				b.rows[i].cells[j].marked = true

				return true
			}
		}
	}

	return false
}

func (b *bingoBoard) hasBingo() (bool, int) {
	// rows
	var bingo bool
	var unmarkedSum int

	for _, v := range b.rows {
		cellsMarked := true
		for _, c := range v.cells {
			if !c.marked {
				cellsMarked = false
				unmarkedSum += c.value
			}
		}

		if cellsMarked {
			bingo = true
		}
	}

	if !bingo {
		// columns
		for i, _ := range b.rows[0].cells {
			cellsMarked := true
			for _, v := range b.rows {
				if !v.cells[i].marked {
					cellsMarked = false
				}
			}

			if cellsMarked {
				bingo = true
				break
			}
		}
	}

	return bingo, unmarkedSum
}

func (b *bingoBoard) playDraft(draft []int) (bool, int, int) {
	for round, d := range draft {
		b.playRound(d)
		bingo, s := b.hasBingo()

		if bingo {
			return true, round + 1, s * d
		}
	}

	return false, 0, 0
}
