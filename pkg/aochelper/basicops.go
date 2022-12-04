package aochelper

func Sum(sl []int) int {
	var sum int

	for i, _ := range sl {
		sum += sl[i]
	}

	return sum
}

func Max(sl []int) int {
	var max int

	for i, _ := range sl {
		if sl[i] > max {
			max = sl[i]
		}
	}

	return max
}

func Min(sl []int) int {
	min := sl[0]

	for i := 1; i < len(sl); i++ {
		if sl[i] < min {
			min = sl[i]
		}
	}

	return min
}
