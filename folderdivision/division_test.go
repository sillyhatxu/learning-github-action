package folderdivision

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivision(t *testing.T) {
	type Test struct {
		i      int
		j      int
		result int
	}
	var array []Test
	array = append(array, Test{i: 48, j: 3, result: 16})
	array = append(array, Test{i: 14, j: 4, result: 3})
	for _, test := range array {
		result := Division(test.i, test.j)
		assert.Equal(t, test.result, result)
	}
}
