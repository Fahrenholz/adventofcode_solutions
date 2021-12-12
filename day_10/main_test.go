package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorruptedStrings(t *testing.T) {
	tts := []struct {
		name string
		set  string
		exp  int
	}{
		{"ex-1", "{([(<{}[<>[]}>{[]{[(<()>", 1197},
		{"ex-2", "[[<[([]))<([[{}[[()]]]", 3},
		{"ex-3", "[{[{({}]{}}([{[{{{}}([]", 57},
		{"ex-4", "[<(<(<(<{}))><([]([]()", 3},
		{"ex-5", "<{([([[(<>()){}]>(<<{{", 25137},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			res, compl := evalLines(tt.set)

			assert.Equal(t, "", compl, "compl should be empty")
			assert.Equal(t, tt.exp, res, "score should be as expected")
		})
	}
}

func TestIncompleteStrings(t *testing.T) {
	tts := []struct {
		name string
		set  string
		exp  string
	}{
		{"ex-1", "[({(<(())[]>[[{[]{<()<>>", "}}]])})]"},
		{"ex-2", "[(()[<>])]({[<{<<[]>>(", ")}>]})"},
		{"ex-3", "(((({<>}<{<{<>}{[]{[]{}", "}}>}>))))"},
		{"ex-4", "{<[[]]>}<{[{[{[]{()[[[]", "]]}}]}]}>"},
		{"ex-5", "<{([{{}}[<[[[<>{}]]]>[]]", "])}>"},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			sc, res := evalLines(tt.set)

			assert.Equal(t, 0, sc, "should have score to 0")
			assert.Equal(t, tt.exp, res, "should have correct result")
		})
	}
}

func TestCalcCompletionScore(t *testing.T) {
	tts := []struct {
		name string
		set  string
		exp  int
	}{
		{"ex-1", "}}]])})]", 288957},
		{"ex-2", ")}>]})", 5566},
		{"ex-3", "}}>}>))))", 1480781},
		{"ex-4", "]]}}]}]}>", 995444},
		{"ex-5", "])}>", 294},
		{"b-1", ")", 1},
		{"b-1", "]", 2},
		{"b-1", "}", 3},
		{"b-1", ">", 4},
		{"b-1", "))", 6},
		{"b-1", "]]", 12},
		{"b-1", "}}", 18},
		{"b-1", ">>", 24},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			res := calcCompletionScore(tt.set)

			assert.Equal(t, tt.exp, res, "should find the right score")
		})
	}
}

func TestFindMiddleScore(t *testing.T) {
	scores := []int{288957, 5566, 1480781, 995444, 294}
	res := findMiddleScore(scores)

	assert.Equal(t, 288957, res, "should find correct middle score")
}
