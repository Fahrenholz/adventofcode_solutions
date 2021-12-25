package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type instruction struct {
	putOn  bool
	xRange [2]int
	yRange [2]int
	zRange [2]int
}

type coordinates struct {
	x, y, z int
}

type crange struct {
	xRange [2]int
	yRange [2]int
	zRange [2]int
}

func main() {
	instructions := parseInputs(getInputsByLine(os.Args[1]))
	partOne(instructions)
	partTwo(instructions)
}

func partOne(instructions []instruction) {
	reactor := make(map[coordinates]bool)

	for _, v := range instructions {
		xR, xInVA := restrictVA(v.xRange)
		yR, yInVA := restrictVA(v.yRange)
		zR, zInVa := restrictVA(v.zRange)

		if !xInVA || !yInVA || !zInVa {
			continue
		}

		for x := xR[0]; x <= xR[1]; x++ {
			for y := yR[0]; y <= yR[1]; y++ {
				for z := zR[0]; z <= zR[1]; z++ {
					reactor[coordinates{x: x, y: y, z: z}] = v.putOn
				}
			}
		}
	}

	cubesOn := 0
	for v := range reactor {
		if reactor[v] {
			cubesOn++
		}
	}

	fmt.Printf("Solution 1: %d cubes put on.\n", cubesOn)
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

func partTwo(instructions []instruction) {
	cubesOn := uint64(0)
	var previous []crange
	for _, v := range instructions {
		rg := crange{xRange: v.xRange, yRange: v.yRange, zRange: v.zRange}

		var np []crange
		for _, pr := range previous {
			rgs := subRange(pr, rg)
			np = append(np, rgs...)
		}
		if v.putOn {
			np = append(np, rg)
		}

		previous = np
	}

	for _, rg := range previous {
		cubesOn += uint64(abs(rg.xRange[1]-rg.xRange[0]+1) * abs(rg.yRange[1]-rg.yRange[0]+1) * abs(rg.zRange[1]-rg.zRange[0]+1))
	}

	fmt.Printf("Solution 2: %d cubes put on.\n", cubesOn)
}

func subRange(a, b crange) []crange {
	ho, ol := determineOverlap(a, b)
	if !ho {
		return []crange{a}
	}

	var res []crange
	var xrs [][2]int
	if a.xRange[0] < ol.xRange[0] {
		xrs = append(xrs, [2]int{a.xRange[0], ol.xRange[0] - 1})
	}
	if a.xRange[1] > ol.xRange[1] {
		xrs = append(xrs, [2]int{ol.xRange[1] + 1, a.xRange[1]})
	}
	xrs = append(xrs, ol.xRange)

	var yrs [][2]int
	if a.yRange[0] < ol.yRange[0] {
		yrs = append(yrs, [2]int{a.yRange[0], ol.yRange[0] - 1})
	}
	if a.yRange[1] > ol.yRange[1] {
		yrs = append(yrs, [2]int{ol.yRange[1] + 1, a.yRange[1]})
	}
	yrs = append(yrs, ol.yRange)

	var zrs [][2]int
	if a.zRange[0] < ol.zRange[0] {
		zrs = append(zrs, [2]int{a.zRange[0], ol.zRange[0] - 1})
	}
	if a.zRange[1] > ol.zRange[1] {
		zrs = append(zrs, [2]int{ol.zRange[1] + 1, a.zRange[1]})
	}
	zrs = append(zrs, ol.zRange)

	for xri := range xrs {
		for yri := range yrs {
			for zri := range zrs {
				if xri == len(xrs)-1 && yri == len(yrs)-1 && zri == len(zrs)-1 {
					continue
				}
				res = append(res, crange{xRange: xrs[xri], yRange: yrs[yri], zRange: zrs[zri]})
			}
		}
	}

	return res
}

func determineOverlap(a, b crange) (bool, crange) {
	if a.xRange[1] < b.xRange[0] || b.xRange[1] < a.xRange[0] {
		return false, crange{}
	}
	if a.yRange[1] < b.yRange[0] || b.yRange[1] < a.yRange[0] {
		return false, crange{}
	}
	if a.zRange[1] < b.zRange[0] || b.zRange[1] < a.zRange[0] {
		return false, crange{}
	}

	//we have an overlap
	nxr := [2]int{int(math.Max(float64(a.xRange[0]), float64(b.xRange[0]))), int(math.Min(float64(a.xRange[1]), float64(b.xRange[1])))}
	nyr := [2]int{int(math.Max(float64(a.yRange[0]), float64(b.yRange[0]))), int(math.Min(float64(a.yRange[1]), float64(b.yRange[1])))}
	nzr := [2]int{int(math.Max(float64(a.zRange[0]), float64(b.zRange[0]))), int(math.Min(float64(a.zRange[1]), float64(b.zRange[1])))}

	return true, crange{xRange: nxr, yRange: nyr, zRange: nzr}
}

func parseInputs(inp []string) []instruction {
	r := regexp.MustCompile(`(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)
	res := make([]instruction, len(inp))

	for i, v := range inp {
		matches := r.FindStringSubmatch(v)
		res[i].putOn = matches[1] == "on"
		res[i].xRange[0], _ = strconv.Atoi(matches[2])
		res[i].xRange[1], _ = strconv.Atoi(matches[3])
		res[i].yRange[0], _ = strconv.Atoi(matches[4])
		res[i].yRange[1], _ = strconv.Atoi(matches[5])
		res[i].zRange[0], _ = strconv.Atoi(matches[6])
		res[i].zRange[1], _ = strconv.Atoi(matches[7])
	}

	return res
}

func restrictVA(o [2]int) ([2]int, bool) {
	if o[1] < -50 || o[0] > 50 {
		return [2]int{0, 0}, false
	}

	res := [2]int{o[0], o[1]}

	if o[0] < -50 {
		res[0] = -50
	}

	if o[1] > 50 {
		res[1] = 50
	}

	return res, true
}

func getInputsByLine(file string) []string {
	inputFile, err := os.Open(file)
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

func min(ins []instruction) coordinates {
	minX := math.MaxInt
	minY := math.MaxInt
	minZ := math.MaxInt

	for _, v := range ins {

		if v.xRange[0] < minX {
			minX = v.xRange[0]
		}

		if v.yRange[0] < minY {
			minY = v.yRange[0]
		}

		if v.zRange[0] < minZ {
			minZ = v.zRange[0]
		}
	}

	return coordinates{x: minX, y: minY, z: minZ}
}

func max(ins []instruction) coordinates {
	maxX := 0
	maxY := 0
	maxZ := 0

	for _, v := range ins {
		if v.xRange[1] > maxX {
			maxX = v.xRange[1]
		}

		if v.yRange[1] > maxY {
			maxY = v.yRange[1]
		}

		if v.zRange[1] > maxZ {
			maxZ = v.zRange[1]
		}
	}

	return coordinates{x: maxX, y: maxY, z: maxZ}
}
