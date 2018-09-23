package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	out := os.Stdout

	var expression string
	var isNumbers bool

	if os.Args[1] == "-numbers" {
		expression = strings.Join(os.Args[2:], " ")
		isNumbers = true
	} else {
		expression = strings.Join(os.Args[1:], "")
		isNumbers = false
	}

	fmt.Println(isNumbers, expression)

	calc(out, expression, isNumbers)
}

func pop(stack *[]int) (value int) {
	value = (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return value
}

func numberCheck(stack *[]int, number *int) (int, int, *int) {
	var y int
	x := pop(stack)
	if number != nil {
		y = *number
		number = nil
	} else {
		y = pop(stack)
	}
	return x, y, number
}

func catchWrongExpr() {
	if err := recover(); err != nil { //catch
		fmt.Println("Wrong expression")
	}
}

func calc(out io.Writer, expression string, isNumbers bool) {
	defer catchWrongExpr()

	var stack []int
	var number *int = nil
	for _, char := range expression {
		var x, y int
		switch string(char) {
		case "+":
			x, y, number = numberCheck(&stack, number)
			fmt.Println(x, " + ", y, " = ", x+y)
			stack = append(stack, y+x)

		case "-":
			x, y, number = numberCheck(&stack, number)
			fmt.Println(x, " - ", y, " = ", x-y)
			stack = append(stack, y-x)

		case "*":
			x, y, number = numberCheck(&stack, number)
			stack = append(stack, y*x)

		case "/":
			x, y, number = numberCheck(&stack, number)
			stack = append(stack, y/x)

		case "=":
			fmt.Fprintf(out, strconv.Itoa((stack[len(stack)-1])))

		default:
			numeral, error := strconv.Atoi(string(char))

			if isNumbers {

				if error != nil {
					if number != nil {
						stack = append(stack, *number)
					}
					number = nil
				} else {
					if number == nil {
						number = new(int)
					}
					(*number) *= 10
					(*number) += numeral
				}

			} else {

				if error != nil {
					fmt.Println("find syntax error")
				} else {
					stack = append(stack, numeral)
				}

			}

		}

	}
}
