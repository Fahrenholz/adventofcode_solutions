package aochelper

import "strconv"

func ForceInt(s string) int {
	res, _ := strconv.Atoi(s)

	return res
}
