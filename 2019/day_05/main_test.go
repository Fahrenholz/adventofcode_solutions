package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetParameterModes(t *testing.T) {
	tts := []struct {
		name          string
		parameterMode int
		expOpcode     int
		expModeC      int
		expModeB      int
		expModeA      int
	}{
		{"simple-opcode", 3, 3, 0, 0, 0},
		{"two-digit-opcode", 99, 99, 0, 0, 0},
		{"modeA-immediate", 10003, 3, 0, 0, 1},
		{"modeB-immediate", 1003, 3, 0, 1, 0},
		{"modeC-immediate", 103, 3, 1, 0, 0},
		{"modeAB-immediate", 11003, 3, 0, 1, 1},
		{"modeABC-immediate", 11103, 3, 1, 1, 1},
		{"twodigit-opcode-modeAB-immediate", 11099, 99, 0, 1, 1},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			opCode, modeC, modeB, modeA := getParameterModes(tt.parameterMode)

			assert.Equal(t, tt.expOpcode, opCode, "opcode should be equal to expected")
			assert.Equal(t, tt.expModeC, modeC, "C should be equal to expected")
			assert.Equal(t, tt.expModeB, modeB, "B should be equal to expected")
			assert.Equal(t, tt.expModeA, modeA, "A should be equal to expected")
		})
	}
}
