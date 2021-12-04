package main

import (
	"github.com/go-test/deep"
	"testing"
)

func TestParseNumberDraft(t *testing.T) {
	set := "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15"
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	res := parseNumberDraft(set)

	if diffs := deep.Equal(res, exp); diffs != nil {
		t.Errorf("\t FAIL not as expected. Diffs : %+v", diffs)
	} else {
		t.Logf("\t PASS: as expected")
	}
}
