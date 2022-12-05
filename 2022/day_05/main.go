package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
	"strings"
)

func main() {
	// Creating the stacks
	nbStacks := aoch.Map(aoch.Filter(aoch.GetInputAsLines(), func(el string) bool { return strings.Contains(el, "1   2") }), func(el string) int {
		el = strings.ReplaceAll(el, " ", "")
		return aoch.ForceInt(el[len(el)-1:])
	})[0]
	stackP1 := make([][]string, nbStacks)
	stackP2 := make([][]string, nbStacks)

	// Setup initial state
	aoch.Walk(aoch.Filter(aoch.GetInputAsLines(), func(el string) bool { return strings.Contains(el, "[") }), func(el string) {
		for currSt := 0; currSt < nbStacks; currSt++ {
			if currSt*4 > len(el) {
				break
			}

			cr := el[currSt*4 : (currSt*4)+3]
			if cr[0:1] == "[" {
				stackP1[currSt] = append(stackP1[currSt], cr[1:2])
				stackP2[currSt] = append(stackP2[currSt], cr[1:2])
			}
		}
	})
	stackP1 = aoch.Map(stackP1, func(el []string) []string { return aoch.Reverse(el) })
	stackP2 = aoch.Map(stackP2, func(el []string) []string { return aoch.Reverse(el) })

	// Now applying instructions - Twice
	aoch.Walk(aoch.Map(aoch.Filter(aoch.GetInputAsLines(), func(el string) bool { return strings.Contains(el, "move") }), func(el string) [3]int {
		els := strings.Split(el, " ")
		return [3]int{aoch.ForceInt(els[1]), aoch.ForceInt(els[3]) - 1, aoch.ForceInt(els[5]) - 1}
	}), func(el [3]int) {
		stackP1[el[2]] = append(stackP1[el[2]], aoch.Reverse(stackP1[el[1]][len(stackP1[el[1]])-el[0]:])...)
		stackP1[el[1]] = stackP1[el[1]][0 : len(stackP1[el[1]])-el[0]]
		stackP2[el[2]] = append(stackP2[el[2]], stackP2[el[1]][len(stackP2[el[1]])-el[0]:]...)
		stackP2[el[1]] = stackP2[el[1]][0 : len(stackP2[el[1]])-el[0]]
	})

	// Display
	fmt.Println("Part One: ", aoch.Reduce(stackP1, "", func(curr string, el []string) string {
		return fmt.Sprintf("%s%s", curr, el[len(el)-1])
	}))
	fmt.Println("Part Two: ", aoch.Reduce(stackP2, "", func(curr string, el []string) string {
		return fmt.Sprintf("%s%s", curr, el[len(el)-1])
	}))
}
