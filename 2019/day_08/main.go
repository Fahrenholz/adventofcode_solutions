package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := getInputsByLine()
	image := toImage(inputs[0])

	min := 6*25 + 1
	minLayerCS := -1

	for _, currLayer := range image {
		cnt := make(map[int]int)
		for _, row := range currLayer {
			for _, px := range row {
				cnt[px]++
			}
		}

		if cnt[0] < min {
			min = cnt[0]
			minLayerCS = cnt[1] * cnt[2]
		}
	}

	fmt.Println("PART ONE | ", minLayerCS)

	var finalImage [6][25]int

	for layerIdx := len(image) - 1; layerIdx >= 0; layerIdx-- {
		for rowIdx, _ := range image[layerIdx] {
			for cellIdx, _ := range image[layerIdx][rowIdx] {
				if image[layerIdx][rowIdx][cellIdx] != 2 {
					finalImage[rowIdx][cellIdx] = image[layerIdx][rowIdx][cellIdx]
				}
			}
		}
	}

	ascii := '█'

	fmt.Println("PART TWO |")
	for _, row := range finalImage {
		for _, cell := range row {
			if cell == 1 {
				fmt.Printf("%c", ascii)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func toImage(pixels string) [][6][25]int {
	pixelSl := strings.Split(pixels, "")
	nbLayers := len(pixelSl) / (6 * 25)

	img := make([][6][25]int, nbLayers)

	for layer := 0; layer < nbLayers; layer++ {
		l := [6][25]int{}
		for row := 0; row < 6; row++ {
			for col := 0; col < 25; col++ {
				l[row][col], _ = strconv.Atoi(pixelSl[(layer*6*25)+(row*25)+col])
			}
		}
		img[layer] = l
	}

	return img
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
