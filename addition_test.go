package learning_github_actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddition(t *testing.T) {
	type Test struct {
		i      int
		j      int
		result int
	}
	var array []Test
	array = append(array, Test{i: 5, j: 8, result: 13})
	array = append(array, Test{i: 1, j: 5, result: 6})
	for _, test := range array {
		result := Addition(test.i, test.j)
		assert.Equal(t, test.result, result)
	}
}
