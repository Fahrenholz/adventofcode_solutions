package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type scanner struct {
	number  int
	reports [][]coordinates
}

type normalizedScanner struct {
	number           int
	adjustmentToOrig coordinates
	reports          []coordinates
}

func normalizeByOverlaps(sc scanner, compareTo *normalizedScanner) *normalizedScanner {
	for _, axis := range sc.reports {
		for _, coordOnPlane := range compareTo.reports {
			for _, candidate := range axis {
				possibleAdjustment := coordinates{x: coordOnPlane.x - candidate.x, y: coordOnPlane.y - candidate.y, z: coordOnPlane.z - candidate.z}
				matches := 0
				for _, validationItem := range axis {
					possibleMatch := coordinates{x: validationItem.x + possibleAdjustment.x, y: validationItem.y + possibleAdjustment.y, z: validationItem.z + possibleAdjustment.z}
					if contains(compareTo.reports, possibleMatch) {
						matches++
					}
				}

				if matches >= 12 {
					return &normalizedScanner{
						number:           sc.number,
						adjustmentToOrig: coordinates{x: compareTo.adjustmentToOrig.x + possibleAdjustment.x, y: compareTo.adjustmentToOrig.y + possibleAdjustment.y, z: compareTo.adjustmentToOrig.z + possibleAdjustment.z},
						reports:          axis,
					}
				}
			}
		}
	}

	return nil
}

type coordinates struct {
	x int
	y int
	z int
}

func main() {
	inputs := parseInputs(getInputsByLine())

	scanners := map[int]*normalizedScanner{0: {
		number:           inputs[0].number,
		adjustmentToOrig: coordinates{x: 0, y: 0, z: 0},
		reports:          inputs[0].reports[0],
	}}

	for len(scanners) < len(inputs) {
		for _, ns := range scanners {
			for _, sc := range inputs {
				if _, ok := scanners[sc.number]; ok {
					continue
				}

				nss := normalizeByOverlaps(sc, ns)
				if nss != nil {
					scanners[nss.number] = nss
				}
			}
		}
	}

	beacons := make(map[coordinates]bool)
	maxManhattan := 0

	for i := range scanners {
		v := scanners[i]
		for _, rp := range v.reports {
			coordsFromOrig := coordinates{x: rp.x + v.adjustmentToOrig.x, y: rp.y + v.adjustmentToOrig.y, z: rp.z + v.adjustmentToOrig.z}
			beacons[coordsFromOrig] = true
		}

		for osci := range scanners {
			if osci == i {
				continue
			}

			x := abs(scanners[osci].adjustmentToOrig.x - v.adjustmentToOrig.x)
			y := abs(scanners[osci].adjustmentToOrig.y - v.adjustmentToOrig.y)
			z := abs(scanners[osci].adjustmentToOrig.z - v.adjustmentToOrig.z)

			max := x + y + z
			if maxManhattan < max {
				maxManhattan = max
			}
		}
	}

	fmt.Printf("Solution 1: %d\n", len(beacons))
	fmt.Printf("Solution 2: %d\n", maxManhattan)
}

func contains(reports []coordinates, coords coordinates) bool {
	for _, v := range reports {
		if v.x == coords.x && v.y == coords.y && v.z == coords.z {
			return true
		}
	}

	return false
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

func parseInputs(inp []string) []scanner {
	var scs []scanner
	nLi := regexp.MustCompile(`--- scanner (\d+) ---`)
	var cur scanner
	cur.reports = make([][]coordinates, 1)
	for _, li := range inp {
		if li == "" {
			cur.reports = permuteAndNegate(cur.reports[0])
			scs = append(scs, cur)
			cur = scanner{}
			cur.reports = make([][]coordinates, 1)
			continue
		}
		if matches := nLi.FindStringSubmatch(li); len(matches) == 2 {
			cur.number, _ = strconv.Atoi(matches[1])
			continue
		}

		xyz := strings.Split(li, ",")
		rep := coordinates{}
		rep.x, _ = strconv.Atoi(xyz[0])
		rep.y, _ = strconv.Atoi(xyz[1])
		rep.z, _ = strconv.Atoi(xyz[2])
		cur.reports[0] = append(cur.reports[0], rep)
	}

	cur.reports = permuteAndNegate(cur.reports[0])
	scs = append(scs, cur)
	return scs
}

func permuteAndNegate(report []coordinates) [][]coordinates {
	res := make([][]coordinates, 6*8)
	possibleNegVals := []coordinates{{x: 1, y: 1, z: 1}, {x: 1, y: 1, z: -1}, {x: 1, y: -1, z: 1}, {x: -1, y: 1, z: 1}, {x: -1, y: -1, z: 1}, {x: -1, y: -1, z: -1}, {x: 1, y: -1, z: -1}, {x: -1, y: 1, z: -1}}

	for _, v := range report {
		for i := 0; i < len(possibleNegVals); i++ {
			res[(i*6)+0] = append(res[(i*6)+0], coordinates{x: v.x * possibleNegVals[i].x, y: v.y * possibleNegVals[i].y, z: v.z * possibleNegVals[i].z})
			res[(i*6)+1] = append(res[(i*6)+1], coordinates{x: v.x * possibleNegVals[i].x, y: v.z * possibleNegVals[i].y, z: v.y * possibleNegVals[i].z})
			res[(i*6)+2] = append(res[(i*6)+2], coordinates{x: v.y * possibleNegVals[i].x, y: v.x * possibleNegVals[i].y, z: v.z * possibleNegVals[i].z})
			res[(i*6)+3] = append(res[(i*6)+3], coordinates{x: v.y * possibleNegVals[i].x, y: v.z * possibleNegVals[i].y, z: v.x * possibleNegVals[i].z})
			res[(i*6)+4] = append(res[(i*6)+4], coordinates{x: v.z * possibleNegVals[i].x, y: v.x * possibleNegVals[i].y, z: v.y * possibleNegVals[i].z})
			res[(i*6)+5] = append(res[(i*6)+5], coordinates{x: v.z * possibleNegVals[i].x, y: v.y * possibleNegVals[i].y, z: v.x * possibleNegVals[i].z})
		}
	}

	return res
}

func getInputsByLine() []string {
	inputFile, err := os.Open(os.Args[1])
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
