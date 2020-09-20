// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Add takes two or subsequent numbers and returns the result of adding them together.
func Add(a, b float64, nums ...float64) float64 {
	result := a + b
	for _, n := range nums {
		result += n
	}
	return result
}

// Subtract takes two or subsequent numbers and returns the result of subtracting the subsequent
// number from previous.
func Subtract(a, b float64, nums ...float64) float64 {
	result := a - b
	for _, n := range nums {
		result -= n
	}
	return result
}

// Multiply takes two or subsequent numbers and returns the result of multiplication of the
// given numbers
func Multiply(a, b float64, nums ...float64) float64 {
	result := a * b
	for _, n := range nums {
		result *= n
	}
	return result
}

// Divide takes two or subsequent numbers and returns the result of dividing the given numbers.
// It returns error if the division operation is not valid
func Divide(a, b float64, nums ...float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divide by Zero")
	}
	result := a / b

	for _, n := range nums {
		if n == 0 {
			return 0, errors.New("Divide by Zero")
		}
		result /= n
	}
	return result, nil
}

// SquareRoot returns the square root of the provided number, or an error if the input
// number is invalid
func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input %f; square root is undefined for negative values", a)
	}
	return math.Sqrt(a), nil
}

// Evaluate processes the input string as two operands separated by the operator
// and returns the result of that operation, or an error if the string is not in
// the acceptable form "operand1 operator operand2". Leading, embedded and trailing
// spaces are allowed between operands and operator.
// Supported operators are : + , -, * and /
func Evaluate(expr string) (float64, error) {
	opIndex := strings.IndexAny(expr, "+-*/")
	if opIndex == -1 {
		return 0, fmt.Errorf("Supported operator (+,-,*,/) not found")
	}

	operand1 := strings.TrimSpace(expr[0:opIndex])
	operand2 := strings.TrimSpace(expr[opIndex+1:])

	num1, err := strconv.ParseFloat(operand1, 64)
	if err != nil {
		return 0, errors.New("error parsing the first number")
	}

	num2, err := strconv.ParseFloat(operand2, 64)
	if err != nil {
		return 0, errors.New("error parsing the second number")
	}

	switch expr[opIndex] {
	case '+':
		return num1 + num2, nil
	case '-':
		return num1 - num2, nil
	case '*':
		return num1 * num2, nil
	case '/':
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("unsupported operator %v in expression %s ", expr[opIndex], expr)
	}
}
