package learning_github_actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubtraction(t *testing.T) {
	type Test struct {
		i      int
		j      int
		result int
	}
	var array []Test
	array = append(array, Test{i: 12, j: 7, result: 5})
	array = append(array, Test{i: 42, j: 5, result: 37})
	for _, test := range array {
		result := Subtraction(test.i, test.j)
		assert.Equal(t, test.result, result)
	}
}
