package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coordinates struct {
	x int
	y int
}

func main() {
	algorithm, litPixels := parse(getInputsByLine(os.Args[1]))
	//process(algorithm, litPixels, 1)  // 24
	process(algorithm, litPixels, 2)  // 35
	process(algorithm, litPixels, 50) // 3351
}

func process(algorithm []string, litPixels map[coordinates]bool, iterations int) {
	for it := 0; it < iterations; it++ {
		litPixels = enhance(algorithm, litPixels, it)
	}

	fmt.Printf("After %d iterations: %d lit pixels\n", iterations, len(litPixels))
}

func enhance(algorithm []string, litPixels map[coordinates]bool, step int) map[coordinates]bool {
	nLitPixels := make(map[coordinates]bool)
	xMin, xMax, yMin, yMax := minMax(litPixels)
	determinedChar := "."
	if step%2 == 1 {
		determinedChar = algorithm[0]
	}

	for y := yMin - 1; y <= yMax+1; y++ {
		for x := xMin - 1; x <= xMax+1; x++ {
			binNb := ""
			for _, ym := range []int{y - 1, y, y + 1} {
				for _, xm := range []int{x - 1, x, x + 1} {
					if yMin <= ym && ym <= yMax && xMin <= xm && xm <= xMax {
						v, fnd := litPixels[coordinates{x: xm, y: ym}]
						if fnd && v {
							binNb = fmt.Sprintf("%s%d", binNb, 1)
							continue
						}
						binNb = fmt.Sprintf("%s%d", binNb, 0)
						continue
					}

					if determinedChar == "." {
						binNb = fmt.Sprintf("%s%d", binNb, 0)
						continue
					}

					binNb = fmt.Sprintf("%s%d", binNb, 1)
				}
			}

			nb, _ := strconv.ParseInt(binNb, 2, 32)

			if algorithm[int(nb)] == "#" {
				nLitPixels[coordinates{x: x + 1, y: y + 1}] = true
			}
		}
	}

	return nLitPixels
}

func minMax(litPixels map[coordinates]bool) (int, int, int, int) {
	xMin := math.MaxInt
	yMin := math.MaxInt
	xMax := 0
	yMax := 0

	for i := range litPixels {
		if i.x < xMin {
			xMin = i.x
		}
		if i.x > xMax {
			xMax = i.x
		}
		if i.y < yMin {
			yMin = i.y
		}
		if i.y > yMax {
			yMax = i.y
		}
	}

	return xMin, xMax, yMin, yMax
}

func parse(inp []string) ([]string, map[coordinates]bool) {
	algorithm := strings.Split(inp[0], "")
	litPixels := make(map[coordinates]bool)

	gridLines := inp[2:]
	for y, v := range gridLines {
		symbols := strings.Split(v, "")
		for x, s := range symbols {
			if s == "#" {
				litPixels[coordinates{x: x, y: y}] = true
			}
		}
	}

	return algorithm, litPixels
}

func getInputsByLine(fName string) []string {
	inputFile, err := os.Open(fName)
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
