package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type point struct {
	height  int
	basinNb int
}

func main() {
	inputs := parseInputs(getInputsByLine())

	lowerPointSum := sumLowerPoints(inputs)

	fmt.Printf("Solution 1: %d\n", lowerPointSum)

	determineBasins(inputs)
	sizes := getBasinSizesSorted(inputs)
	fmt.Printf("Solution 2: %d\n", sizes[0]*sizes[1]*sizes[2])
}

func getBasinSizesSorted(inputs [][]point) []int {
	sizeMap := make(map[int]int)
	for _, vec := range inputs {
		for _, pt := range vec {
			if pt.basinNb == 0 {
				continue
			}

			if _, ok := sizeMap[pt.basinNb]; !ok {
				sizeMap[pt.basinNb] = 0
			}

			sizeMap[pt.basinNb]++
		}
	}

	var sizes []int

	for _, v := range sizeMap {
		sizes = append(sizes, v)
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return sizes
}

func parseInputs(inp []string) [][]point {
	res := make([][]point, len(inp))
	for i, v := range inp {
		res[i] = make([]point, len(v))
		for pos, r := range v {
			res[i][pos] = point{height: int(r - '0')}
		}
	}

	return res
}

func determineBasins(inputs [][]point) {
	currentBasinNb := 1
	for ord, _ := range inputs {
		for abs, _ := range inputs[ord] {
			if inputs[ord][abs].height != 9 && inputs[ord][abs].basinNb == 0 {
				markBasin(inputs, ord, abs, currentBasinNb)
				currentBasinNb++
			}
		}
	}
}

func markBasin(inputs [][]point, ord int, abs int, currentBasinNb int) {
	if inputs[ord][abs].height == 9 || inputs[ord][abs].basinNb != 0 {
		return
	}

	inputs[ord][abs].basinNb = currentBasinNb
	if abs != len(inputs[ord])-1 {
		markBasin(inputs, ord, abs+1, currentBasinNb)
	}
	if abs != 0 {
		markBasin(inputs, ord, abs-1, currentBasinNb)
	}
	if ord != len(inputs)-1 {
		markBasin(inputs, ord+1, abs, currentBasinNb)
	}
	if ord != 0 {
		markBasin(inputs, ord-1, abs, currentBasinNb)
	}

	return
}

func sumLowerPoints(inputs [][]point) int {
	sum := 0
	for ord, vec := range inputs {
		for abs, pt := range vec {
			isLower := true

			if abs != len(vec)-1 {
				isLower = isLower && pt.height < vec[abs+1].height
			}
			if abs != 0 {
				isLower = isLower && pt.height < vec[abs-1].height
			}

			if ord != len(inputs)-1 {
				isLower = isLower && pt.height < inputs[ord+1][abs].height
			}
			if ord != 0 {
				isLower = isLower && pt.height < inputs[ord-1][abs].height
			}

			if isLower {
				sum += 1 + pt.height
			}
		}
	}

	return sum
}

func getInputsByLine() []string {
	inputFile, err := os.Open("./inputs.txt")
	if err != nil {
		fmt.Println("could not find file")
		os.Exit(1)
	}

	defer inputFile.Close()

	var inputs []string

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return inputs
}
