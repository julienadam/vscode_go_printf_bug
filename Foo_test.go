package Foo

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_can_load_real_data(t *testing.T) {

	result := DoStuffAndPrint()
	assert.Equal(t, false, result)
}
