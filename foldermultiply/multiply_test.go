package foldermultiply

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiply(t *testing.T) {
	type Test struct {
		i      int
		j      int
		result int
	}
	var array []Test
	array = append(array, Test{i: 3, j: 5, result: 15})
	array = append(array, Test{i: 6, j: 8, result: 48})
	for _, test := range array {
		result := Multiply(test.i, test.j)
		assert.Equal(t, test.result, result)
	}
}
