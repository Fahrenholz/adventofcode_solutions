package main

import (
	"github.com/go-test/deep"
	"testing"
)

var draft = []int{14, 30, 18, 8, 3, 10, 77, 4, 48, 67, 28, 38, 63, 43, 62, 12, 68, 88, 54, 32, 17, 21, 83, 64, 97, 53, 24, 2, 60, 96, 86, 23, 20, 93, 65, 34, 45, 46, 42, 49, 71, 9, 61, 16, 31, 1, 29, 40, 59, 87, 95, 41, 39, 27, 6, 25, 19, 58, 80, 81, 50, 79, 73, 15, 70, 37, 92, 94, 7, 55, 85, 98, 5, 84, 99, 26, 66, 57, 82, 75, 22, 89, 74, 36, 11, 76, 56, 33, 13, 72, 35, 78, 47, 91, 51, 44, 69, 0, 90, 52}

func TestNewBoard(t *testing.T) {
	set := []string{
		"62 42 24  0 53",
		"41 94 70 88 33",
		"32 19 43 21 23",
		"84 98 60 39 36",
		" 5  4 49 76 82",
	}

	exp := &bingoBoard{rows: []row{
		{cells: []cell{{value: 62}, {value: 42}, {value: 24}, {value: 0}, {value: 53}}},
		{cells: []cell{{value: 41}, {value: 94}, {value: 70}, {value: 88}, {value: 33}}},
		{cells: []cell{{value: 32}, {value: 19}, {value: 43}, {value: 21}, {value: 23}}},
		{cells: []cell{{value: 84}, {value: 98}, {value: 60}, {value: 39}, {value: 36}}},
		{cells: []cell{{value: 5}, {value: 4}, {value: 49}, {value: 76}, {value: 82}}},
	}}

	res := NewBoard(set)

	if diffs := deep.Equal(exp, res); diffs != nil {
		t.Errorf("\t FAILED: diffs found: %+v", diffs)
	} else {
		t.Logf("\t PASS: no diffs found")
	}
}

func TestPlayDraft(t *testing.T) {

	tts := []struct {
		name string
		set  []int
		exp  int
	}{
		{"draft-bingo-row", []int{62, 42, 24, 0, 53}, 997 * 53},
		{"draft-bingo-column", []int{62, 41, 32, 84, 5}, 954 * 5},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			board := &bingoBoard{rows: []row{
				{cells: []cell{{value: 62}, {value: 42}, {value: 24}, {value: 0}, {value: 53}}},
				{cells: []cell{{value: 41}, {value: 94}, {value: 70}, {value: 88}, {value: 33}}},
				{cells: []cell{{value: 32}, {value: 19}, {value: 43}, {value: 21}, {value: 23}}},
				{cells: []cell{{value: 84}, {value: 98}, {value: 60}, {value: 39}, {value: 36}}},
				{cells: []cell{{value: 5}, {value: 4}, {value: 49}, {value: 76}, {value: 82}}},
			}}

			bingo, round, score := board.playDraft(tt.set)

			if !bingo {
				t.Errorf("\t FAIL: should have bingo")
			} else {
				t.Logf("\t PASS: has bingo")
			}

			if round != 5 {
				t.Errorf("\t FAIL: should return in round 5")
			} else {
				t.Logf("\t PASS: returns in round 5")
			}

			if score != tt.exp {
				t.Errorf("\t FAIL: should have score equal to %d, %d found", tt.exp, score)
			} else {
				t.Logf("\t PASS: has score equal to %d", tt.exp)
			}
		})
	}
}
