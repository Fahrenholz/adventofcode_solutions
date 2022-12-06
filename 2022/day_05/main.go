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
	stacksP1 := make([][]string, nbStacks)
	stacksP2 := make([][]string, nbStacks)

	// Setup initial state
	aoch.Walk(aoch.Filter(aoch.GetInputAsLines(), func(el string) bool { return strings.Contains(el, "[") }), func(el string) {
		for currSt := 0; currSt < nbStacks; currSt++ {
			if currSt*4 > len(el) {
				break
			}

			cr := el[currSt*4 : (currSt*4)+3]
			if cr[0:1] == "[" {
				stacksP1[currSt] = append(stacksP1[currSt], cr[1:2])
				stacksP2[currSt] = append(stacksP2[currSt], cr[1:2])
			}
		}
	})
	stacksP1 = aoch.Map(stacksP1, func(el []string) []string { return aoch.Reverse(el) })
	stacksP2 = aoch.Map(stacksP2, func(el []string) []string { return aoch.Reverse(el) })

	// Now applying instructions - Twice
	aoch.Walk(aoch.Map(aoch.Filter(aoch.GetInputAsLines(), func(el string) bool { return strings.Contains(el, "move") }), func(el string) [3]int {
		els := strings.Split(el, " ")
		return [3]int{aoch.ForceInt(els[1]), aoch.ForceInt(els[3]) - 1, aoch.ForceInt(els[5]) - 1}
	}), func(el [3]int) {
		stacksP1[el[2]] = append(stacksP1[el[2]], aoch.Reverse(stacksP1[el[1]][len(stacksP1[el[1]])-el[0]:])...)
		stacksP1[el[1]] = stacksP1[el[1]][0 : len(stacksP1[el[1]])-el[0]]
		stacksP2[el[2]] = append(stacksP2[el[2]], stacksP2[el[1]][len(stacksP2[el[1]])-el[0]:]...)
		stacksP2[el[1]] = stacksP2[el[1]][0 : len(stacksP2[el[1]])-el[0]]
	})

	// Display
	fmt.Println("Part One: ", aoch.Reduce(stacksP1, "", func(curr string, el []string) string {
		return fmt.Sprintf("%s%s", curr, el[len(el)-1])
	}))
	fmt.Println("Part Two: ", aoch.Reduce(stacksP2, "", func(curr string, el []string) string {
		return fmt.Sprintf("%s%s", curr, el[len(el)-1])
	}))
}
