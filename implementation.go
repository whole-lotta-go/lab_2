package lab2

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// EvalPostfix evaluates a postfix (Reverse Polish Notation) expression and returns the result.
// The input is a string consisting of space-separated tokens, where tokens can be integers or operators.
// Supported operators are: + (addition), - (subtraction), * (multiplication), / (division), and ^ (exponentiation).
//
// Parameters:
//   - input: A string representing the postfix expression (e.g., "3 4 + 2 *").
//
// Returns:
//   - int: The result of evaluating the postfix expression.
//   - error: An error if the expression is invalid.
func EvalPostfix(input string) (int, error) {
	stack := []int{}
	tokens := strings.Fields(input)

	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, errors.New("invalid postfix expression: not enough operands")
			}

			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result int
			switch token {
			case "+":
				result = operand1 + operand2
			case "-":
				result = operand1 - operand2
			case "*":
				result = operand1 * operand2
			case "/":
				if operand2 == 0 {
					return 0, errors.New("division by zero")
				}
				result = operand1 / operand2
			case "^":
				result = int(math.Pow(float64(operand1), float64(operand2)))
			default:
				return 0, fmt.Errorf("invalid operator: %s", token)
			}

			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid postfix expression: too many operands")
	}

	return stack[0], nil
}
