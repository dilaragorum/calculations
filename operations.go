package calculations

import (
	"errors"
	"strings"
)

func Sum(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divider cannot be zero")
	}
	return a / b, nil
}

func IsHttp(url string) bool {
	return strings.HasPrefix(url, "http://")
}

// TODO: IsHttp -> http:// --> true ya da false
// TODO: IsHttps -> https://
