package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInputs(t *testing.T) {
	grid, instructions := getInputsByLine("./testinputs.txt")

	assert.Equal(t, 18, len(grid.values))
	assert.Equal(t, 10, grid.maxX)
	assert.Equal(t, 14, grid.maxY)

	assert.Equal(t, 2, len(instructions))
}
