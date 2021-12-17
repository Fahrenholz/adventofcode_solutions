package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var targetArea = map[string]int{
	"minX": 0,
	"maxX": 0,
	"minY": 0,
	"maxY": 0,
}

type coordinates struct {
	x int
	y int
}

func main() {
	inputs := getInputsByLine()
	re := regexp.MustCompile(`target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)`)
	sub := re.FindStringSubmatch(inputs[0])
	targetArea["minX"], _ = strconv.Atoi(sub[1])
	targetArea["maxX"], _ = strconv.Atoi(sub[2])
	targetArea["minY"], _ = strconv.Atoi(sub[3])
	targetArea["maxY"], _ = strconv.Atoi(sub[4])

	highestY := 0
	velocities := 0

	for y := targetArea["minY"]; y < targetArea["minY"]*-1; y++ {
		for x := 0; x <= targetArea["maxX"]; x++ {
			pos := coordinates{x: 0, y: 0}
			velo := coordinates{x: x, y: y}
			highestLocalY := 0
			for pos.y > targetArea["minY"] {
				pos, velo = step(pos, velo)
				if highestLocalY < pos.y {
					highestLocalY = pos.y
				}
				if pos.x <= targetArea["maxX"] && pos.x >= targetArea["minX"] && pos.y >= targetArea["minY"] && pos.y <= targetArea["maxY"] {
					velocities++
					if highestY < highestLocalY {
						highestY = highestLocalY
					}
					break
				}
			}

		}
	}

	fmt.Printf("Solution 1: %d\n", highestY)
	fmt.Printf("Solution 2: %d\n", velocities)
}

func step(pos coordinates, velocity coordinates) (newPos coordinates, newVelocity coordinates) {
	newPos.x = pos.x + velocity.x
	newPos.y = pos.y + velocity.y
	newVelocity.y = velocity.y - 1
	newVelocity.x = velocity.x

	if velocity.x > 0 {
		newVelocity.x -= 1
	}

	if velocity.x < 0 {
		newVelocity.x += 1
	}

	return
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
