package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInputsForTwo(t *testing.T) {
	inp := []string{
		"12",
		"34",
	}
	expMaxX := 9
	expMaxY := 9

	expRes := map[point]int{
		point{x: 0, y: 0}: 1,
		point{x: 1, y: 0}: 2,
		point{x: 2, y: 0}: 2,
		point{x: 3, y: 0}: 3,
		point{x: 4, y: 0}: 3,
		point{x: 5, y: 0}: 4,
		point{x: 6, y: 0}: 4,
		point{x: 7, y: 0}: 5,
		point{x: 8, y: 0}: 5,
		point{x: 9, y: 0}: 6,
		point{x: 0, y: 1}: 3,
		point{x: 1, y: 1}: 4,
		point{x: 2, y: 1}: 4,
		point{x: 3, y: 1}: 5,
		point{x: 4, y: 1}: 5,
		point{x: 5, y: 1}: 6,
		point{x: 6, y: 1}: 6,
		point{x: 7, y: 1}: 7,
		point{x: 8, y: 1}: 7,
		point{x: 9, y: 1}: 8,
		point{x: 0, y: 2}: 2,
		point{x: 1, y: 2}: 3,
		point{x: 2, y: 2}: 3,
		point{x: 3, y: 2}: 4,
		point{x: 4, y: 2}: 4,
		point{x: 5, y: 2}: 5,
		point{x: 6, y: 2}: 5,
		point{x: 7, y: 2}: 6,
		point{x: 8, y: 2}: 6,
		point{x: 9, y: 2}: 7,
	}

	res, start, goal := parseInputsForTwo(inp)

	for i := range expRes {
		assert.Equal(t, expRes[i], res[i], "expects equality for point (x: %d, y:%d)", i.x, i.y)
	}

	assert.Equal(t, 0, start.x)
	assert.Equal(t, 0, start.y)
	assert.Equal(t, expMaxX, goal.x)
	assert.Equal(t, expMaxY, goal.y)
}

func TestParseInputsForTwoSecond(t *testing.T) {
	inp := []string{"8"}

	exp := map[point]int{
		point{x: 0, y: 0}: 8,
		point{x: 1, y: 0}: 9,
		point{x: 2, y: 0}: 1,
		point{x: 3, y: 0}: 2,
		point{x: 4, y: 0}: 3,
		point{x: 0, y: 1}: 9,
		point{x: 1, y: 1}: 1,
		point{x: 2, y: 1}: 2,
		point{x: 3, y: 1}: 3,
		point{x: 4, y: 1}: 4,
		point{x: 0, y: 2}: 1,
		point{x: 1, y: 2}: 2,
		point{x: 2, y: 2}: 3,
		point{x: 3, y: 2}: 4,
		point{x: 4, y: 2}: 5,
	}

	res, _, _ := parseInputsForTwo(inp)

	fmt.Println(res)

	for i := range exp {
		assert.Equal(t, exp[i], res[i], "expects equality for point (x: %d, y:%d)", i.x, i.y)
	}
}

func TestSolve(t *testing.T) {
	inp := []string{
		"12",
		"34",
	}

	r1, s1, g1 := parseInputs(inp)
	slv1 := solve(r1, s1, g1)
	assert.Equal(t, 6, slv1)
}
