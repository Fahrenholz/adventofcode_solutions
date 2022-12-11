package main

type Monkey struct {
	Items          []int
	Op             func(int) int
	Test           func(int) bool
	True           int
	False          int
	InspectedItems int
}

func getMonkeys() []Monkey {
	return []Monkey{
		{
			Items: []int{73, 77},
			Op:    func(el int) int { return el * 5 },
			Test:  func(el int) bool { return el%11 == 0 },
			True:  6,
			False: 5,
		},
		{
			Items: []int{57, 88, 80},
			Op:    func(el int) int { return el + 5 },
			Test:  func(el int) bool { return el%19 == 0 },
			True:  6,
			False: 0,
		},
		{
			Items: []int{61, 81, 84, 69, 77, 88},
			Op:    func(el int) int { return el * 19 },
			Test:  func(el int) bool { return el%5 == 0 },
			True:  3,
			False: 1,
		},
		{
			Items: []int{78, 89, 71, 60, 81, 84, 87, 75},
			Op:    func(el int) int { return el + 7 },
			Test:  func(el int) bool { return el%3 == 0 },
			True:  1,
			False: 0,
		},
		{
			Items: []int{60, 76, 90, 63, 86, 87, 89},
			Op:    func(el int) int { return el + 2 },
			Test:  func(el int) bool { return el%13 == 0 },
			True:  2,
			False: 7,
		},
		{
			Items: []int{88},
			Op:    func(el int) int { return el + 1 },
			Test:  func(el int) bool { return el%17 == 0 },
			True:  4,
			False: 7,
		},
		{
			Items: []int{84, 98, 78, 85},
			Op:    func(el int) int { return el * el },
			Test:  func(el int) bool { return el%7 == 0 },
			True:  5,
			False: 4,
		},
		{
			Items: []int{98, 89, 78, 73, 71},
			Op:    func(el int) int { return el + 4 },
			Test:  func(el int) bool { return el%2 == 0 },
			True:  3,
			False: 2,
		},
	}
}

func getTestMonkeys() []Monkey {
	return []Monkey{
		{
			Items: []int{79, 98},
			Op:    func(el int) int { return el * 19 },
			Test:  func(el int) bool { return el%23 == 0 },
			True:  2,
			False: 3,
		},
		{
			Items: []int{54, 65, 75, 74},
			Op:    func(el int) int { return el + 6 },
			Test:  func(el int) bool { return el%19 == 0 },
			True:  2,
			False: 0,
		},
		{
			Items: []int{79, 60, 97},
			Op:    func(el int) int { return el * el },
			Test:  func(el int) bool { return el%13 == 0 },
			True:  1,
			False: 3,
		},
		{
			Items: []int{74},
			Op:    func(el int) int { return el + 3 },
			Test:  func(el int) bool { return el%17 == 0 },
			True:  0,
			False: 1,
		},
	}
}
