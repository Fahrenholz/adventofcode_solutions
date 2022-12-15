package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
	"regexp"
)

func main() {
	beacons := make(map[[2]int]bool)
	safelyNoBeacon := make(map[[2]int]bool)
	re := regexp.MustCompile("^Sensor at x=([-]*\\d+), y=([-]*\\d+): closest beacon is at x=([-]*\\d+), y=([-]*\\d+)$")

	inputs := aoch.Map(aoch.GetInputAsLines(), func(el string) [2][2]int {
		matches := re.FindStringSubmatch(el)
		beacon := [2]int{aoch.ForceInt(matches[3]), aoch.ForceInt(matches[4])}
		beacons[beacon] = true

		return [2][2]int{{aoch.ForceInt(matches[1]), aoch.ForceInt(matches[2])}, beacon}
	})

	aoch.Walk(inputs, func(el [2][2]int) {
		for x := el[0][0] - manhattanDistance(el[0], el[1]); x <= el[0][0]+manhattanDistance(el[0], el[1]); x++ {
			if manhattanDistance(el[0], [2]int{x, 2000000}) <= manhattanDistance(el[0], el[1]) && !beacons[[2]int{x, 2000000}] {
				safelyNoBeacon[[2]int{x, 2000000}] = true
			}
		}
	})

	fmt.Println("Part one: ", len(safelyNoBeacon))

	foundBeaconPos := false
	for y := 0; y <= 4000000 && !foundBeaconPos; y++ {
		for x := 0; x <= 4000000 && !foundBeaconPos; x++ {
			isInRange := false
			for _, el := range inputs {
				if manhattanDistance(el[0], [2]int{x, y}) <= manhattanDistance(el[0], el[1]) {
					isInRange = true
					x += manhattanDistance(el[0], el[1]) - manhattanDistance(el[0], [2]int{x, y})
					break
				}
			}
			if !isInRange {
				//10961118625406
				fmt.Println("Part two: ", x*4000000+y, "(", x, ", ", y, ")")
				foundBeaconPos = true
			}
		}
	}
}

func manhattanDistance(point1, point2 [2]int) int {
	return aoch.Abs(point1[0]-point2[0]) + aoch.Abs(point1[1]-point2[1])
}
