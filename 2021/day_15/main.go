package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

var directions = []point{
	{x: -1, y: 0},
	{x: 1, y: 0},
	{x: 0, y: -1},
	{x: 0, y: 1},
}

func main() {
	riskGrid, start, goal := parseInputs(getInputsByLine())
	fmt.Printf("Solution 1: %d\n", solve(riskGrid, start, goal))
	riskGrid, start, goal = parseInputsForTwo(getInputsByLine())
	fmt.Printf("Solution 2: %d\n", solve(riskGrid, start, goal))
}

func solve(riskGrid map[point]int, start, goal point) int {
	queue := make(PriorityQueue, 0)
	heap.Init(&queue)
	heap.Push(&queue, queueItem{pos: start})

	pointsCost := make(map[point]int)

	for queue.Len() > 0 {
		cur := heap.Pop(&queue).(queueItem)
		if cur.pos == goal {
			break
		}

		for i := 0; i < 4; i++ {
			np := point{x: cur.pos.x + directions[i].x, y: cur.pos.y + directions[i].y}
			if np.x > goal.x || np.x < 0 || np.y > goal.y || np.y < 0 {
				continue
			}
			cost := cur.riskLevel + riskGrid[np]

			if v, ok := pointsCost[np]; !ok || v > cost {
				pointsCost[np] = cost
				heap.Push(&queue, queueItem{pos: np, riskLevel: cost})
			}
		}
	}

	return pointsCost[goal]
}

func parseInputs(inp []string) (map[point]int, point, point) {
	res := make(map[point]int)
	maxord := len(inp) - 1
	var maxAbs int
	for ord, v := range inp {
		vv := strings.Split(v, "")
		maxAbs = len(vv) - 1
		for abs, rlStr := range vv {
			rl, _ := strconv.Atoi(rlStr)
			res[point{x: abs, y: ord}] = rl
		}
	}

	return res, point{x: 0, y: 0}, point{x: maxAbs, y: maxord}
}

func parseInputsForTwo(inp []string) (map[point]int, point, point) {
	res := make(map[point]int)
	maxord := (len(inp) * 5) - 1
	var maxAbs int
	for _, tilesY := range []int{0, 1, 2, 3, 4} {
		for ord, v := range inp {
			vv := strings.Split(v, "")
			maxAbs = (len(vv) * 5) - 1
			for _, tilesX := range []int{0, 1, 2, 3, 4} {
				for abs, rlStr := range vv {
					rl, _ := strconv.Atoi(rlStr)
					crl := rl + tilesX + tilesY
					if crl-9 > 0 {
						crl = crl % 9
					}
					res[point{x: abs + (len(vv) * tilesX), y: ord + (len(v) * tilesY)}] = crl
				}
			}
		}
	}

	return res, point{x: 0, y: 0}, point{x: maxAbs, y: maxord}
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
