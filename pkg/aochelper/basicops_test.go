package aochelper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	strsl := []string{"a", "b", "c", "d"}
	filtered := Filter(strsl, func(i int) bool {
		return strsl[i] == "c"
	})

	assert.Equal(t, filtered, []string{"c"}, "should have filtered string slice")

	intsl := []int{1, 2, 3, 4}
	filteredtwo := Filter(intsl, func(i int) bool {
		return intsl[i] == 2
	})

	assert.Equal(t, []int{2}, filteredtwo, "should have filtered int slice")
}
