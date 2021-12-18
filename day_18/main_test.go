package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExplode(t *testing.T) {
	tts := []struct {
		name string
		set  string
		exp  string
	}{
		{"ex-1", "[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"ex-2", "[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"ex-3", "[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"ex-4", "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"ex-5", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			term, _ := parseString(tt.set, 0)
			_ = term.explodeTerm(1, false)
			assert.Equal(t, tt.exp, term.String())
		})
	}
}

func TestSplit(t *testing.T) {
	v := 13
	v1 := 6
	v2 := 7

	tts := []struct {
		name string
		set  *term
		exp  string
	}{
		{"ex-1", &term{literal: &v}, "[6,7]"},
		{"ex-2", &term{left: &term{literal: &v}, right: &term{literal: &v}}, "[[6,7],13]"},
		{"ex-3", &term{left: &term{left: &term{literal: &v1}, right: &term{literal: &v2}}, right: &term{literal: &v}}, "[[6,7],[6,7]]"},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.split()
			assert.Equal(t, tt.exp, tt.set.String())
		})
	}
}

func TestReduce(t *testing.T) {
	tts := []struct {
		name string
		set1 string
		set2 string
		exp  string
	}{
		{"ex-1", "[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
		{"ex-2", "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]", "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]", "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			t1, _ := parseString(tt.set1, 0)
			t2, _ := parseString(tt.set2, 0)
			term := &term{left: t1, right: t2}

			res := term.reduceTerm()

			assert.Equal(t, tt.exp, res.String())
		})
	}
}
