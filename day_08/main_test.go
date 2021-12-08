package main

import "testing"

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

			if res != tt.exp {
				t.Errorf("\tFAIL: expected %d got %d", tt.exp, res)
			} else {
				t.Logf("\tPASS: expected value returned")
			}
		})
	}
}

func TestContains(t *testing.T) {
	res := contains("cdbaf", "ab")

	if !res {
		t.Errorf("\tFAIL: should be true")
	} else {
		t.Logf("\tPASS: expected value returned")
	}

	res2 := contains("cdfgeb", "ab")
	if res2 {
		t.Errorf("\tFAIL: should be false")
	} else {
		t.Logf("\tPASS: expected value returned")
	}
}
