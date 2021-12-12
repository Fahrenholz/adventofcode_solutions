package main

import (
	"github.com/idealo/abs-billing/pkg/assert"
	"testing"
)

var testInputs = [][]string{{"start", "A"}, {"start", "b"}, {"A", "c"}, {"A", "b"}, {"b", "d"}, {"A", "end"}, {"b", "end"}}

func TestRecursion(t *testing.T) {
	res := recursePathFinding(testInputs, []string{"start"}, false)
	res2 := recursePathFinding(testInputs, []string{"start"}, true)

	assert.Eq(t, 10, len(res), "should find 10 paths")
	assert.Eq(t, 36, len(res2), "should find 10 paths")
}
