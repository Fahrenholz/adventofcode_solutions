package main

import (
	"github.com/go-test/deep"
	"testing"
)

func TestParseInputs(t *testing.T) {
	tts := []struct {
		name    string
		set     []string
		exp     []vector
		expMaxX int
		expMaxY int
	}{
		{"simple-vertical", []string{"1,1 -> 1,50"}, []vector{{from: coords{x: 1, y: 1}, to: coords{x: 1, y: 50}}}, 1, 50},
		{"simple-horizontal", []string{"1,1 -> 50,1"}, []vector{{from: coords{x: 1, y: 1}, to: coords{x: 50, y: 1}}}, 50, 1},
		{"simple-inverted-horizontal", []string{"50,1 -> 1,1"}, []vector{{from: coords{x: 50, y: 1}, to: coords{x: 1, y: 1}}}, 50, 1},
		{"diagonal-down-up", []string{"1,1 -> 50,50"}, []vector{{from: coords{x: 1, y: 1}, to: coords{x: 50, y: 50}}}, 50, 50},
		{"diagonal-up-down", []string{"1,50 -> 50,1"}, []vector{{from: coords{x: 1, y: 50}, to: coords{x: 50, y: 1}}}, 50, 50},
		{"diagonal-complicated", []string{"770,861 -> 915,716"}, []vector{{from: coords{x: 770, y: 861}, to: coords{x: 915, y: 716}}}, 915, 861},
		{"mixed", []string{"1,1 -> 70,50"}, nil, 0, 0},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			result, maxX, maxY := getVectors(tt.set)
			if diffs := deep.Equal(result, tt.exp); diffs != nil {
				t.Errorf("\tFAIL: should not get any differences: %+v", diffs)
			} else {
				t.Logf("\tPASS: no differences.")
			}

			if maxX != tt.expMaxX {
				t.Errorf("\tFAIL: should not get any different maxX: %d instead of %d", maxX, tt.expMaxX)
			} else {
				t.Logf("\tPASS: no differences.")
			}

			if maxY != tt.expMaxY {
				t.Errorf("\tFAIL: should not get any different maxY: %d instead of %d", maxY, tt.expMaxY)
			} else {
				t.Logf("\tPASS: no differences.")
			}
		})
	}
}

func TestBuildMatrix(t *testing.T) {
	tts := []struct {
		name         string
		set          []vector
		setMaxX      int
		setMaxY      int
		expCoordsTwo coords
		expCoordsOne coords
	}{
		{"small-two-vectors", []vector{{from: coords{x: 1, y: 1}, to: coords{x: 1, y: 70}}, {from: coords{1, 1}, to: coords{50, 1}}}, 50, 70, coords{x: 1, y: 1}, coords{x: 50, y: 1}},
		{"three-vectors-with-diagonal", []vector{{from: coords{x: 1, y: 1}, to: coords{x: 1, y: 70}}, {from: coords{1, 1}, to: coords{50, 1}}, {from: coords{2, 2}, to: coords{49, 49}}}, 50, 70, coords{x: 1, y: 1}, coords{x: 11, y: 11}},
		{"three-vectors-with-diagonal", []vector{{from: coords{x: 1, y: 1}, to: coords{x: 1, y: 70}}, {from: coords{1, 1}, to: coords{50, 1}}, {from: coords{49, 2}, to: coords{2, 49}}}, 50, 70, coords{x: 1, y: 1}, coords{x: 3, y: 48}},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			matrix := buildMatrix(tt.set, tt.setMaxX, tt.setMaxY)

			if len(matrix) != tt.setMaxY+1 {
				t.Errorf("\tFAIL: wrong ordinate: %d instead of %d", len(matrix), tt.setMaxY)
			} else {
				t.Logf("\tPASS: correct ordinate.")
			}

			if len(matrix[0]) != tt.setMaxX+1 {
				t.Errorf("\tFAIL: wrong abscissa: %d instead of %d", len(matrix[0]), tt.setMaxX)
			} else {
				t.Logf("\tPASS: correct abscissa.")
			}

			if matrix[tt.expCoordsTwo.y][tt.expCoordsTwo.x] != 2 {
				t.Errorf("\tFAIL: expects tested coordinate to have 2, got %d", matrix[tt.expCoordsTwo.y][tt.expCoordsTwo.x])
			} else {
				t.Logf("\tPASS: tested coordinate has 2")
			}

			if matrix[tt.expCoordsOne.y][tt.expCoordsOne.x] != 1 {
				t.Errorf("\tFAIL: expects tested coordinate to have 1, got %d", matrix[tt.expCoordsOne.y][tt.expCoordsOne.x])
			} else {
				t.Logf("\tPASS: tested coordinate has 1")
			}
		})
	}
}
