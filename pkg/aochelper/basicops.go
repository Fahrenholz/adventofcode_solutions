package aochelper

func SumScore[T any](vals []T, score func(elem T) int) int {
	var sum int

	for i, _ := range vals {
		sum += score(vals[i])
	}

	return sum
}

func Sum(vals []int) int {
	return SumScore(vals, func(elem int) int { return elem })
}

func Filter[T any](vars []T, f func(el T) bool) []T {
	var res []T

	for _, v := range vars {
		if f(v) {
			res = append(res, v)
		}
	}

	return res
}

func FilterWithIndex[T any](sl []T, cond func(i int) bool) []T {
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

func Reverse[T any](vals []T) []T {
	res := make([]T, len(vals))

	for i := 0; i < len(vals); i++ {
		res[i] = vals[len(vals)-1-i]
	}

	return res
}

func Map[T, U any](vals []T, f func(T) U) []U {
	var res []U

	for _, v := range vals {
		res = append(res, f(v))
	}

	return res
}

func MapToMap[K comparable, T, U any](vals map[K]T, f func(T) U) map[K]U {
	res := make(map[K]U)

	for i, v := range vals {
		res[i] = f(v)
	}

	return res
}

func Walk[T any](vals []T, f func(T)) {
	for _, v := range vals {
		f(v)
	}
}

func Reduce[T, U any](vals []T, init U, f func(U, T) U) U {
	res := init
	for _, v := range vals {
		res = f(res, v)
	}

	return res
}

func ReduceMap[K comparable, T, U any](vals map[K]T, init U, f func(U, T) U) U {
	res := init
	for _, v := range vals {
		res = f(res, v)
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

func MapKeysToSlice[K comparable, T any](vals map[K]T) []K {
	var res []K

	for v, _ := range vals {
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

func GroupNItems[T any](vals []T, n int) [][]T {
	var res [][]T

	tmp := []T{}

	for i, v := range vals {
		if i != 0 && i%n == 0 {
			res = append(res, tmp)
			tmp = []T{}
		}

		tmp = append(tmp, v)
	}

	res = append(res, tmp)

	return res
}

func Max(vals []int) int {
	return MaxScore(vals, func(elem int) int { return elem })
}

func MinScore[T any](vals []T, score func(elem T) int) T {
	min := score(MaxScore(vals, score))
	var minElem T

	for i, _ := range vals {
		curr := score(vals[i])

		if curr <= min {
			min = curr
			minElem = vals[i]
		}
	}

	return minElem
}

func Min(vals []int) int {
	return MinScore(vals, func(elem int) int { return elem })
}

func MaxIdx(vals []int, lastIdx bool) int {
	var max, maxIdx int

	for i, _ := range vals {
		if vals[i] > max {
			max = vals[i]
			maxIdx = i
		}
		if vals[i] == max && lastIdx {
			max = vals[i]
			maxIdx = i
		}
	}

	return maxIdx
}

func Flip[T any](vals [][]T) [][]T {
	res := make([][]T, len(vals[0]))
	for j, _ := range vals[0] {
		for i, _ := range vals {
			res[j] = append(res[j], vals[i][j])
		}
	}

	return res
}
