package main

import (
	"bufio"
	"fmt"
	"os"
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

		text := scanner.Text()

		texts := strings.Split(text, " ")

		left_operand, err_1 := strconv.ParseFloat(texts[0], 10)
		right_operand, err_2 := strconv.ParseFloat(texts[2], 10)

		if err_1 != nil || err_2 != nil {
			fmt.Println(ERROR)
			continue
		}

		operator := texts[1]

		var result float64

		switch operator {
		case "+":
			result, _ = addition(left_operand, right_operand)
			break
		case "-":
			result, _ = subtraction(left_operand, right_operand)
			break
		case "*":
			result, _ = multiplication(left_operand, right_operand)
			break
		case "/":
			result, _ = division(left_operand, right_operand)
			break
		default:
			fmt.Println(ERROR)
			continue
		}

		fmt.Println(left_operand, operator, right_operand, "=", result)
	}
}

func addition(left_operand, right_operand float64) (result float64, err error) {
	result = left_operand + right_operand
	return
}

func subtraction(left_operand, right_operand float64) (result float64, err error) {
	result = left_operand - right_operand
	return
}

func multiplication(left_operand, right_operand float64) (result float64, err error) {
	result = left_operand * right_operand
	return
}

func division(left_operand, right_operand float64) (result float64, err error) {
	result = left_operand / right_operand
	return
}
