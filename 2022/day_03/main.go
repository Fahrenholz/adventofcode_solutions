package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
)

var priorities = map[string]int{}

func init() {
	for i := 'a'; i <= 'z'; i++ {
		priorities[string(i)] = int(i) - int('a') + 1
	}
	for i := 'A'; i <= 'Z'; i++ {
		priorities[string(i)] = int(i) - int('A') + 27
	}
}

func main() {
	fmt.Println(aoch.Sum(aoch.Map(aoch.GetInputsAsLinesOfStringSlices(""), func(elem []string) int {
		tmp := []map[string]bool{aoch.SliceToMap(elem[:len(elem)/2]), aoch.SliceToMap(elem[len(elem)/2:])}
		return aoch.SumScore(aoch.MapKeysToSlice(aoch.FilterMap(tmp[0], func(key string) bool {
			return tmp[1][key]
		})), func(elem string) int {
			return priorities[elem]
		})
	})))

	fmt.Println(aoch.Sum(aoch.Map(aoch.GroupNItems(aoch.GetInputsAsLinesOfStringSlices(""), 3), func(elem [][]string) int {
		tmp := []map[string]bool{aoch.SliceToMap(elem[0]), aoch.SliceToMap(elem[1]), aoch.SliceToMap(elem[2])}
		return aoch.Reduce(aoch.MapKeysToSlice(aoch.FilterMap(aoch.FilterMap(tmp[0], f(tmp[1])), f(tmp[2]))), 0, func(sum int, elem string) int {
			return sum + priorities[elem]
		})
	})))
}

func f(contains map[string]bool) func(elem string) bool {
	return func(elem string) bool {
		return contains[elem]
	}
}
