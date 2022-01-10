package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateValues(t *testing.T) {
	tts := []struct {
		name string
		set  boardingPass
		exp  boardingPass
	}{
		{"ex-1", boardingPass{row: "FBFBBFF", column: "RLR"}, boardingPass{rowInt: 44, columnInt: 5, id: 357}},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.set.calculateAttributes()
			assert.Equal(t, tt.exp.rowInt, p.rowInt)
			assert.Equal(t, tt.exp.columnInt, p.columnInt)
			assert.Equal(t, tt.exp.id, p.id)
		})
	}
}
