package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testPatterns = parseInputs([]string{
	"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf",
	"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
	"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
})

func TestDecodeEntry(t *testing.T) {
	tts := []struct {
		name string
		set  entry
		exp  int
	}{
		{"pattern-1", testPatterns[0], 5353},
		{"pattern-2", testPatterns[1], 8394},
		{"pattern-3", testPatterns[2], 9781},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			res := decodeEntry(tt.set)
			assert.Equal(t, tt.exp, res, "should find equal four-digit-value")
		})
	}
}

func TestContains(t *testing.T) {
	res := contains("cdbaf", "ab")
	assert.True(t, res, "should find stringparts")

	res2 := contains("cdfgeb", "ab")
	assert.False(t, res2, "should not find stringparts")
}
