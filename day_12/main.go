package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs := parseInputs(getInputsByLine())

	paths := recursePathFinding(inputs, []string{"start"}, false)

	fmt.Printf("Solution 1: %d paths visit small caves at most once\n", len(paths))

	morePaths := recursePathFinding(inputs, []string{"start"}, true)

	fmt.Printf("Solution 2: %d paths visit small caves at most twice\n", len(morePaths))
}

func parseInputs(inp []string) [][]string {
	res := make([][]string, len(inp))
	for i, v := range inp {
		res[i] = strings.Split(v, "-")
	}

	return res
}

func recursePathFinding(inputs [][]string, points []string, allowPassTwice bool) [][]string {
	var res [][]string
	for _, v := range inputs {
		isPossible, dst := getDst(points[len(points)-1], v)
		if isPossible {
			apt := allowPassTwice
			//starting point for current segment is end point of last
			if dst == "end" {
				pts := make([]string, len(points)+1)
				copy(pts, points)
				pts[len(pts)-1] = "end"
				res = append(res, pts)

				continue
			}

			if strings.ToLower(dst) == dst {
				//is lowercase, we have to check if already passed before continuing
				if alreadyInPath(points, dst) {
					if !apt {
						continue //ignoring this possibility
					}

					apt = false
				}
			}

			cp := make([]string, len(points)+1)
			copy(cp, points)
			cp[len(cp)-1] = dst
			res = append(res, recursePathFinding(inputs, cp, apt)...)
		}
	}

	return res
}

func getDst(lastPt string, current []string) (bool, string) {
	if current[0] == lastPt {
		return true, current[1]
	}
	if current[1] == lastPt && current[0] != "start" {
		return true, current[0]
	}

	return false, ""
}

func alreadyInPath(pts []string, pt string) bool {
	found := false
	for _, p := range pts {
		if p == pt {
			found = true
		}
	}

	return found
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
