package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const PROMPT = "> "
const ERROR = "ERROR"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(PROMPT)

		scanner.Scan()

		raw_text := scanner.Text()

		text := removeDuplicateSpaces(raw_text)

		tokens := strings.Split(text, " ")

		result, err := process(tokens)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(text, "=", result)
		}
	}
}

func process(tokens []string) (result string, err error) {
	tokens, err = calculate(tokens, "*", "/")
	if err != nil {
		return
	}

	tokens, err = calculate(tokens, "+", "-")
	if err != nil {
		return
	}

	if len(tokens) == 1 {
		result = tokens[0]
	} else {
		err = errors.New(ERROR)
	}

	return
}

func calculate(tmp_tokens []string, operator_1, operator_2 string) (tokens []string, err error) {
	tokens = tmp_tokens

	for {
		index := -1
		index, err = find_valid_operator(tokens, operator_1, operator_2)
		if err != nil {
			return
		}

		if index == -1 {
			break
		}

		var result float64
		result, err = evaluate(tokens[index-1 : index+2])
		if err != nil {
			return
		}

		tokens[index+1] = fmt.Sprintf("%g", result)
		tokens = append(tokens[:index-1], tokens[index+1:]...)
	}
	return
}

func evaluate(formula []string) (result float64, err error) {

	left_operand, err_1 := strconv.ParseFloat(formula[0], 10)
	right_operand, err_2 := strconv.ParseFloat(formula[2], 10)

	if err_1 != nil || err_2 != nil {
		err = errors.New(ERROR)
		return
	}

	operator := formula[1]

	switch operator {
	case "+":
		result, _ = add(left_operand, right_operand)
		break
	case "-":
		result, _ = subtract(left_operand, right_operand)
		break
	case "*":
		result, _ = multiply(left_operand, right_operand)
		break
	case "/":
		result, err = divide(left_operand, right_operand)
		break
	default:
		err = errors.New(ERROR)
		return
	}

	return
}

func add(left_operand, right_operand float64) (result float64, err error) {
	result = left_operand + right_operand
	return
}

func subtract(left_operand, right_operand float64) (result float64, err error) {
	result = left_operand - right_operand
	return
}

func multiply(left_operand, right_operand float64) (result float64, err error) {
	result = left_operand * right_operand
	return
}

func divide(left_operand, right_operand float64) (result float64, err error) {

	if right_operand == 0 {
		err = errors.New(ERROR)
	} else {
		result = left_operand / right_operand
	}

	return
}

func find_valid_operator(tokens []string, operand_1, operand_2 string) (index int, err error) {
	for i, token := range tokens {
		if token == operand_1 || token == operand_2 {

			// operator must be at odd position and not the last one in array
			if i%2 == 0 || i == len(tokens)-1 {
				err = errors.New(ERROR)
				return
			}

			index = i
			return
		}
	}

	index = -1
	return
}

func removeDuplicateSpaces(raw_text string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(raw_text, " ")
}
