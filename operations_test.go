package calculations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(2, 2)
	assert.Equal(t, 4, result)
}

func TestMinus(t *testing.T) {
	result := Minus(2, 1)
	assert.Equal(t, 1, result)
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 2)
	assert.Equal(t, 4, result)
}

func TestIsHttp(t *testing.T) {
	result := IsHttp("http")
	assert.Equal(t, false, result)
}

func TestIsHttps(t *testing.T) {
	result := IsHttps("https://")
	assert.Equal(t, true, result)
}

func TestDivide(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result, err := Divide(5, 1)
		assert.Equal(t, 5, result)
		assert.Nil(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		_, err := Divide(5, 0)
		assert.Errorf(t, err, "divider cannot be zero")
	})
}
