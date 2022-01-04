package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var lastTerm *term
var nextTermAdds *term

type term struct {
	literal *int
	left    *term
	right   *term
}

func (t *term) getMagnitude() int {
	if t.literal != nil {
		return *t.literal
	}

	return (3 * t.left.getMagnitude()) + (2 * t.right.getMagnitude())
}

func (t *term) String() string {
	if t.literal != nil {
		return strconv.Itoa(*t.literal)
	}

	return "[" + t.left.String() + "," + t.right.String() + "]"
}

func main() {
	inputs := getInputsByLine(os.Args[1])
	var sum *term

	var parsedInputs []*term

	for _, inp := range inputs {
		parsed, _ := parseString(inp, 0)
		parsedInputs = append(parsedInputs, parsed)

		if sum == nil {
			sum = parsed
			continue
		}

		sum = &term{
			left:  sum,
			right: parsed,
		}

		sum = sum.reduceTerm()

	}

	fmt.Printf("Solution 1: %d\n", sum.getMagnitude())

	maxMagnitude := 0

	for i, t1S := range inputs {
		for j, t2S := range inputs {
			if i == j {
				continue
			}

			t1, _ := parseString(t1S, 0)
			t2, _ := parseString(t2S, 0)

			secondSum := &term{
				left:  t1,
				right: t2,
			}

			secondSum = secondSum.reduceTerm()
			magn := secondSum.getMagnitude()

			if maxMagnitude < magn {
				maxMagnitude = magn
			}
		}
	}

	fmt.Printf("Solution 2: %d\n", maxMagnitude)

}

func (t *term) reduceTerm() *term {
	res := t
	executedAction := true
	for executedAction {
		executedAction = res.explodeTerm(1, false)
		if !executedAction {
			executedAction = res.split()
		}
	}
	return res
}

func (t *term) split() bool {
	var executedAction bool

	if t.literal == nil {
		executedAction = t.left.split()
		if !executedAction {
			executedAction = t.right.split()
		}

		return executedAction
	}

	if *t.literal > 9 {
		v1 := *t.literal / 2
		v2 := *t.literal - v1
		t.literal = nil
		t.left = &term{literal: &v1}
		t.right = &term{literal: &v2}
		executedAction = true
	}

	return executedAction
}

func (t *term) explodeTerm(depth int, executedAction bool) bool {
	if depth == 1 {
		nextTermAdds = nil
		lastTerm = nil
	}

	if depth >= 5 && !executedAction && t.literal == nil && t.left.literal != nil && t.right.literal != nil {
		nextTermAdds = t.right
		n := 0

		if lastTerm != nil {
			*lastTerm.literal += *t.left.literal
		}

		t.left = nil
		t.right = nil
		t.literal = &n

		return true
	}

	var exAct bool

	if t.literal == nil {
		exAct = t.left.explodeTerm(depth+1, executedAction)
		exAct = t.right.explodeTerm(depth+1, exAct)

		return exAct
	}

	if nextTermAdds != nil {
		*t.literal += *nextTermAdds.literal
		nextTermAdds = nil
	}

	lastTerm = t

	return executedAction
}

func parseString(str string, nextChar int) (*term, int) {
	nc := nextChar
	var res term

	switch str[nextChar] {
	case '[':
		res.left, nc = parseString(str, nc+1)
		res.right, nc = parseString(str, nc+1)
	default:
		v, _ := strconv.Atoi(string(str[nc]))
		res.literal = &v
	}

	return &res, nc + 1
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
