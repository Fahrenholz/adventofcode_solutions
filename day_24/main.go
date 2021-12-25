package main

import "fmt"

// These are input-specific
var zDivisors = [14]int{1, 1, 1, 1, 1, 26, 1, 26, 26, 1, 26, 26, 26, 26}
var checks = [14]int{10, 12, 10, 12, 11, -16, 10, -11, -13, 13, -8, -1, -4, -14}
var offsets = [14]int{12, 7, 8, 8, 15, 12, 8, 13, 3, 13, 3, 9, 4, 13}

// These are for running the program
var minNumber = [14]int{}
var maxNumber = [14]int{}
var stack [][2]int

func main() {
	for step := 0; step < 14; step++ {
		if zDivisors[step] == 1 { //push on stack
			push([2]int{step, offsets[step]})
		} else {
			stackItem := pop()
			min := -1
			max := 0
			for i := 1; i <= 9; i++ {
				possibleJ := (i + stackItem[1]) + checks[step]
				if 0 < possibleJ && possibleJ < 10 {
					if min == -1 {
						min = i
					}
					max = i
				}
			}
			minNumber[stackItem[0]] = min
			minNumber[step] = (min + stackItem[1]) + checks[step]
			maxNumber[stackItem[0]] = max
			maxNumber[step] = (max + stackItem[1]) + checks[step]
		}
	}

	fmt.Printf("Solution 1: %s\n", printNumber(maxNumber))
	fmt.Printf("Solution 2: %s\n", printNumber(minNumber))
}

func printNumber(nb [14]int) string {
	res := ""
	for _, v := range nb {
		res = fmt.Sprintf("%s%d", res, v)
	}

	return res
}

func pop() [2]int {
	nb := stack[0]
	stack = stack[1:]

	return nb
}

func push(nb [2]int) {
	nStack := [][2]int{nb}
	stack = append(nStack, stack...)
}
