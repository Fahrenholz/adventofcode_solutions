package main

import (
	"encoding/json"
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
	"reflect"
	"sort"
	"strings"
)

func main() {
	inputs := aoch.GetInputAsLines()

	var pairs [][2][]any
	var currentPair [2][]any

	for i, _ := range inputs {
		if (i+1)%3 == 0 {
			pairs = append(pairs, currentPair)
			currentPair = [2][]any{}
			continue
		}
		err := json.Unmarshal([]byte(inputs[i]), &currentPair[i%3])
		if err != nil {
			panic(err)
		}
	}

	var rightIndices []int
	var pkgs []any

	for i, _ := range pairs {
		pkgs = append(pkgs, pairs[i][0], pairs[i][1])
		if cmp(pairs[i][0], pairs[i][1]) == -1 {
			rightIndices = append(rightIndices, i+1)
		}
	}

	fmt.Println("Part one: ", aoch.Sum(rightIndices))
	var d1 []any
	var d2 []any

	_ = json.Unmarshal([]byte("[[2]]"), &d1)
	_ = json.Unmarshal([]byte("[[6]]"), &d2)
	pkgs = append(pkgs, d1, d2)

	sort.Slice(pkgs, func(i, j int) bool {
		return cmp(pkgs[i], pkgs[j]) == -1
	})

	var indices []int
	for i, v := range pkgs {
		if cmp(v, d1) == 0 || cmp(v, d2) == 0 {
			indices = append(indices, i+1)
		}
	}

	fmt.Println("Part two: ", aoch.Reduce(indices, 1, func(mul int, el int) int { return mul * el }))
}

func cmp(left any, right any) int {
	isListLeft := isList(reflect.TypeOf(left))
	isListRight := isList(reflect.TypeOf(right))
	switch {
	case !isListLeft && !isListRight:
		leftInt := decodeNumber(left)
		rightInt := decodeNumber(right)
		switch {
		case leftInt < rightInt:
			return -1
		case leftInt > rightInt:
			return 1
		default:
			return 0
		}
	case isListLeft && isListRight:
		leftList := left.([]any)
		rightList := right.([]any)
		for i := 0; i < len(leftList) && i < len(rightList); i++ {
			res := cmp(leftList[i], rightList[i])
			if res != 0 {
				return res
			}
		}
		return cmp(len(leftList), len(rightList))
	case !isListLeft:
		leftInt := decodeNumber(left)
		return cmp([]any{leftInt}, right)
	default:
		rightInt := decodeNumber(right)
		return cmp(left, []any{rightInt})
	}
}

func isList(t reflect.Type) bool {
	return strings.HasPrefix(t.String(), "[]")
}

func decodeNumber(v any) int {
	if reflect.TypeOf(v).Name() == "int" {
		return v.(int)
	} else {
		return int(v.(float64))
	}
}
