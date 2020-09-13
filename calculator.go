// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// AddN takes a slice of numbers and returns the addition of the given numbers
func AddN(numbers []float64) float64 {
	result := 0.0
	for _, n := range numbers {
		result += n
	}
	return result
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// SubtractN takes a slice of numbers and returns the subtraction of the given numbers
// num1 - num2 - num3 etc..
func SubtractN(numbers []float64) float64 {
	length := len(numbers)

	if length == 0 {
		return 0
	}
	if length == 1 {
		return numbers[0]
	}

	result := numbers[0]
	for i := 1; i < length; i++ {
		result -= numbers[i]
	}
	return result
}

// Multiply takes two numbers and returns the result of multiplication of both
// the numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// MultiplyN takes a slice of numbers and returns result of multiplication of
// all the numbers in the slice
func MultiplyN(numbers []float64) float64 {
	length := len(numbers)

	if length == 0 {
		return 0
	}
	if length == 1 {
		return numbers[0]
	}

	result := numbers[0]
	for i := 1; i < length; i++ {
		result *= numbers[i]
	}
	return result
}

// Divide takes two numbers and returns the result of dividing the first number
// by second number. It returns error if the division operation is not valid
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divide by Zero")
	}
	return a / b, nil
}

// DivideN takes a slice of numbers and returns the division of the given numbers
// ((num1 / num2) / num3) / num4 etc..
func DivideN(numbers []float64) (float64, error) {
	length := len(numbers)

	if length == 0 {
		return 0, nil
	}
	if length == 1 {
		return numbers[0], nil
	}

	result := numbers[0]
	for i := 1; i < length; i++ {
		if numbers[i] == 0 {
			return 0, errors.New("Divide by 0")
		}
		result /= numbers[i]
	}
	return result, nil
}

// SquareRoot returns the square root of the provided number , error if the input
// number is invalid
func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("negative numbers don't have square root")
	}
	return math.Sqrt(a), nil
}

// ComputeString processes the input string as two operands separated by the operator
// and returns the result of that operation. Returns error if the string is not in
// the acceptable form.
func ComputeString(str string) (float64, error) {
	for _, op := range "+-*/" {
		if !strings.ContainsAny(str, string(op)) {
			continue
		}

		elements := strings.Split(str, string(op))

		if len(elements) != 2 {
			return 0, errors.New("did not get two operands")
		}

		num1, err := strconv.ParseFloat(strings.TrimSpace(elements[0]), 64)
		if err != nil {
			return 0, errors.New("error parsing the first number")
		}

		num2, err := strconv.ParseFloat(strings.TrimSpace(elements[1]), 64)
		if err != nil {
			return 0, errors.New("error parsing the second number")
		}

		switch op {
		case '+':
			return num1 + num2, nil
		case '-':
			return num1 - num2, nil
		case '*':
			return num1 * num2, nil
		case '/':
			return num1 / num2, nil
		}
	}
	return 0, errors.New("invalid input string")
}
