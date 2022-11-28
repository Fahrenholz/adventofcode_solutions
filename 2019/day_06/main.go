package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type OrbitMass struct {
	label    string
	parent   *OrbitMass
	children []*OrbitMass
}

func main() {
	inputs := getInputsByLine()
	comMass := &OrbitMass{label: "COM", parent: nil}

	startingMass := buildOrbitMassTree(comMass, inputs)
	nbSteps := step(startingMass, 0)
	orbitalTransfers := nbSteps
	if nbSteps > 1 {
		orbitalTransfers = nbSteps - 2
	}

	fmt.Println("Part 1: ", getChecksum(comMass, 0))
	fmt.Println("Part 2: ", orbitalTransfers)
}

func buildOrbitMassTree(mass *OrbitMass, list map[string][]string) *OrbitMass {
	var youMass *OrbitMass

	for _, v := range list[mass.label] {
		tmpMass := &OrbitMass{label: v, parent: mass}
		foundMass := buildOrbitMassTree(tmpMass, list)
		if foundMass != nil {
			youMass = foundMass
		}
		mass.children = append(mass.children, tmpMass)
		if v == "YOU" {
			youMass = tmpMass
		}
	}

	return youMass
}

func step(mass *OrbitMass, it int) int {
	if mass.label == "SAN" {
		return it
	}
	//fmt.Println(mass.label)

	for _, m := range mass.children {
		ch := strings.Join(getLabels(m), ",")
		//fmt.Println(ch)
		if strings.Contains(ch, "SAN") {
			//fmt.Println("IT ", it, ": go to ", m.label)
			return step(m, it+1)
		}
	}

	//fmt.Println("IT ", it, ": return to ", mass.parent.label)
	return step(mass.parent, it+1)
}

func getLabels(mass *OrbitMass) []string {
	var labels []string

	for _, m := range mass.children {
		currentLabels := getLabels(m)
		labels = append(labels, currentLabels...)
	}
	labels = append(labels, mass.label)

	return labels
}

func getChecksum(mass *OrbitMass, indirectOrbits int) int {
	sum := indirectOrbits
	for _, v := range mass.children {
		sum += getChecksum(v, indirectOrbits+1)
	}

	return sum
}

func getInputsByLine() map[string][]string {
	inputFile, err := os.Open(fmt.Sprintf("./%s.txt", os.Args[1]))
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

	list := make(map[string][]string)

	for _, v := range inputs {
		tmp := strings.Split(v, ")")
		list[tmp[0]] = append(list[tmp[0]], tmp[1])
	}

	return list
}
