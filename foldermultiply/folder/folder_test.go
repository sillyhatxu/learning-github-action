package folder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNeedToTest(t *testing.T) {
	err := NeedToTest()
	assert.Nil(t, err)
}
