package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
	"os"
	"sort"
)

func main() {
	var monkeys, monkeysp2 []Monkey
	var modifier int
	var debug bool

	switch {
	case len(os.Args) == 1:
		monkeys = getMonkeys()
		monkeysp2 = getMonkeys()
		modifier = 11 * 19 * 5 * 3 * 13 * 17 * 7 * 2
		debug = false
	case os.Args[1] == "test":
		monkeys = getTestMonkeys()
		monkeysp2 = getTestMonkeys()
		modifier = 23 * 19 * 13 * 17
		debug = true
	}

	for round := 1; round <= 20; round++ {

		for i, _ := range monkeys {
			for _, v := range monkeys[i].Items {
				if debug {
					fmt.Println("Monkey ", i, " inspects an item with a worry level of ", v)
				}
				monkeys[i].InspectedItems++
				v = monkeys[i].Op(v)
				if debug {
					fmt.Println("Worry level gets changed to ", v)
				}
				v = v / 3
				if debug {
					fmt.Println("Worry level gets changed to ", v, " as monkey gets bored")
				}

				v = v % modifier

				if monkeys[i].Test(v) {
					if debug {
						fmt.Println("test successful. Monkey throws it to ", monkeys[i].True)
					}
					monkeys[monkeys[i].True].Items = append(monkeys[monkeys[i].True].Items, v)
				} else {
					if debug {
						fmt.Println("test failed. Monkey throws it to ", monkeys[i].False)
					}
					monkeys[monkeys[i].False].Items = append(monkeys[monkeys[i].False].Items, v)
				}
			}
			monkeys[i].Items = []int{}
			if debug {
				fmt.Println()
			}
		}
		if debug {
			fmt.Println("Monkey 0: ", monkeys[0].Items)
			fmt.Println("Monkey 1: ", monkeys[1].Items)
			fmt.Println("Monkey 2: ", monkeys[2].Items)
			fmt.Println("Monkey 3: ", monkeys[3].Items)
			fmt.Println()
		}
	}

	var inspected []int
	for i, _ := range monkeys {
		inspected = append(inspected, monkeys[i].InspectedItems)
	}
	sort.Ints(inspected)
	inspected = aoch.Reverse(inspected)

	fmt.Println("Part one: ", inspected[0]*inspected[1])

	for round := 1; round <= 10000; round++ {
		for i, _ := range monkeysp2 {
			for _, v := range monkeysp2[i].Items {
				monkeysp2[i].InspectedItems++
				v = monkeysp2[i].Op(v)
				v = v % modifier
				if monkeysp2[i].Test(v) {
					monkeysp2[monkeysp2[i].True].Items = append(monkeysp2[monkeysp2[i].True].Items, v)
				} else {
					monkeysp2[monkeys[i].False].Items = append(monkeysp2[monkeysp2[i].False].Items, v)
				}
			}
			monkeysp2[i].Items = []int{}
		}
		if debug {
			fmt.Println("Monkey 0: ", monkeysp2[0].Items)
			fmt.Println("Monkey 1: ", monkeysp2[1].Items)
			fmt.Println("Monkey 2: ", monkeysp2[2].Items)
			fmt.Println("Monkey 3: ", monkeysp2[3].Items)
			fmt.Println()
		}
	}

	inspected = []int{}
	for i, _ := range monkeysp2 {
		inspected = append(inspected, monkeysp2[i].InspectedItems)
	}
	sort.Ints(inspected)
	inspected = aoch.Reverse(inspected)

	fmt.Println("Part two: ", inspected[0]*inspected[1])
}
