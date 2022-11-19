package eval

import (
	"math"
)

// mathFunc defines a mathematical func
type mathFunc func(x, y int) int

// funcMap holds all allowed mathFunc
type funcMap map[string]mathFunc

// doFunc runs a mathFunc
func DoFunc(op string, x int, y int) int {
	f := getFunc(op)
	return f(x, y)
}

// getFunc returns a mathFunc based on input operator
func getFunc(op string) mathFunc {
	var functions = funcMap{
		"+": Sum,
		"-": Sub,
		"*": Mul,
		"/": Div,
		"^": Pow,
	}
	return functions[op]
}

/* Allowed math functions */

func Sum(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return y - x
}

func Mul(x, y int) int {
	return x * y
}

func Div(x, y int) int {
	return int(y / x)
}

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
