package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")
	if input == "" {
		return "", fmt.Errorf("Bad input, empty string %w", errorEmptyInput)
	}

	re := regexp.MustCompile(`[+-]*[^+^-]+[+-]{1}`)
	indexes := re.FindAllIndex([]byte(input), 2)

	if len(indexes) != 1 {
		return "", fmt.Errorf("Bad input, wrong operands count: %w", errorNotTwoOperands)
	}
	index := indexes[0]

	number1 := input[index[0] : index[1]-1]
	operation := string(input[index[1]-1])
	number2 := input[index[1]:]

	number1Int, err := strconv.Atoi(number1)
	if err != nil {
		return "", fmt.Errorf("first number convert to int error: %w", err)
	}

	number2Int, err := strconv.Atoi(number2)
	if err != nil {
		return "", fmt.Errorf("second number convert to int error: %w", err)
	}

	var answer int
	if operation == "+" {
		answer = number1Int + number2Int
	} else {
		answer = number1Int - number2Int
	}

	return strconv.Itoa(answer), nil
}
