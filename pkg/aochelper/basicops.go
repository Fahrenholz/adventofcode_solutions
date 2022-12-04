package aochelper

func SumScore[T any](vals []T, score func(elem T) int) int {
	var sum int

	for i, _ := range vals {
		sum += score(vals[i])
	}

	return sum
}

func Filter[T any](sl []T, cond func(i int) bool) []T {
	var res []T

	for i, _ := range sl {
		if cond(i) {
			res = append(res, sl[i])
		}
	}

	return res
}

func FilterMap[K comparable, T any](vals map[K]T, cond func(key K) bool) map[K]T {
	res := make(map[K]T)

	for k, _ := range vals {
		if cond(k) {
			res[k] = vals[k]
		}
	}

	return res
}

func SliceToMap[T comparable](vals []T) map[T]bool {
	res := make(map[T]bool)

	for _, v := range vals {
		res[v] = true
	}

	return res
}

func MapToSlice[K comparable, T any](vals map[K]T) []T {
	var res []T

	for _, v := range vals {
		res = append(res, v)
	}

	return res
}

func MaxScore[T any](vals []T, score func(elem T) int) T {
	var max int
	var maxElem T

	for i, _ := range vals {
		curr := score(vals[i])

		if curr > max {
			max = curr
			maxElem = vals[i]
		}
	}

	return maxElem
}

func MinScore[T any](vals []T, score func(elem T) int) T {
	var min int
	var minElem T

	for i, _ := range vals {
		curr := score(vals[i])

		if curr < min {
			min = curr
			minElem = vals[i]
		}
	}

	return minElem
}
