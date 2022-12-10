package main

import (
	"fmt"
	aoch "github.com/fahrenholz/adventOfCode/pkg/aochelper"
	"reflect"
	"strings"
)

const threshold = 100000
const secondThreshold = 30000000
const fsSize = 70000000

type sizeable interface {
	getSize() int
}

type Dir struct {
	Name     string
	children map[string]sizeable
	parent   *Dir
}

func (d *Dir) getSize() int {
	sum := 0
	for _, child := range d.children {
		sum += child.getSize()
	}
	return sum
}

type File struct {
	name string
	size int
}

func (f *File) getSize() int {
	return f.size
}

func main() {

	lines := aoch.GetInputAsLines()

	var dirStack *Dir
	var curr *Dir

	for _, line := range lines {
		switch {
		case strings.Contains(line, "$ cd .."):
			curr = curr.parent
		case strings.Contains(line, "$ cd /") && dirStack == nil:
			dirStack = &Dir{
				Name:     strings.ReplaceAll(line, "$ cd", ""),
				children: make(map[string]sizeable),
				parent:   nil,
			}
			curr = dirStack
		case strings.Contains(line, "dir "):
			if _, ok := curr.children[strings.ReplaceAll(line, "dir ", "")]; !ok {
				newDir := &Dir{
					Name:     strings.ReplaceAll(line, "dir ", ""),
					children: make(map[string]sizeable),
					parent:   curr,
				}
				curr.children[strings.ReplaceAll(line, "dir ", "")] = newDir
			}
		case strings.Contains(line, "$ cd"):
			if _, ok := curr.children[strings.ReplaceAll(line, "$ cd ", "")]; ok {
				curr = curr.children[strings.ReplaceAll(line, "$ cd ", "")].(*Dir)
			}
		case strings.Contains(line, "$ ls"):
		default:
			file := strings.Split(line, " ")
			curr.children[file[1]] = &File{name: file[1], size: aoch.ForceInt(file[0])}
		}
	}

	fmt.Println("Part One: ", aoch.Sum(getDirSizesBelowThreshold(dirStack)))
	fmt.Println("Part Two: ", aoch.Min(getDirSizesAboveSecondThreshold(dirStack, fsSize-dirStack.getSize())))
}

func getDirSizesBelowThreshold(curr *Dir) []int {
	var sizes []int
	for _, v := range curr.children {
		if getType(v) == "*Dir" {
			sizes = append(sizes, getDirSizesBelowThreshold(v.(*Dir))...)
		}
	}

	currSize := curr.getSize()
	if currSize <= threshold {
		sizes = append(sizes, currSize)
	}

	return sizes
}

func getDirSizesAboveSecondThreshold(curr *Dir, currFreeSize int) []int {
	var sizes []int
	for _, v := range curr.children {
		if getType(v) == "*Dir" {
			sizes = append(sizes, getDirSizesAboveSecondThreshold(v.(*Dir), currFreeSize)...)
		}
	}

	currSize := curr.getSize()
	if currFreeSize+currSize >= secondThreshold {
		sizes = append(sizes, currSize)
	}

	return sizes
}

func getType(myvar interface{}) (res string) {
	t := reflect.TypeOf(myvar)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
		res += "*"
	}
	return res + t.Name()
}
