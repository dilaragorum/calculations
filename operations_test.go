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
