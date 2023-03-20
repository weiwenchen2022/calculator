// Package calculator provides a library for
// simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the
// result of adding them together.
func Add(a, b float64, c ...float64) float64 {
	a += b
	for _, v := range c {
		a += v
	}

	return a
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64, c ...float64) float64 {
	a -= b
	for _, v := range c {
		a -= v
	}

	return a
}

func Multiply(a, b float64, c ...float64) float64 {
	a *= b
	for _, v := range c {
		a *= v
	}

	return a
}

func Divide(a, b float64, c ...float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divided by zero")
	}

	a /= b
	for _, v := range c {
		if v == 0 {
			return 0, errors.New("divided by zero")
		}

		a /= v
	}

	return a, nil
}

func Sqrt(a float64) (float64, error) {
	if a <= 0 {
		return 0, errors.New("a <=0")
	}

	return math.Sqrt(a), nil
}
