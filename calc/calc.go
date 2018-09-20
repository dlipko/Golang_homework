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
	if !(len(os.Args) == 2) {
		panic("usage go run main.go 25+=")
	}

	expression := strings.Replace(strings.Join(os.Args[1:], ""), " ", "", -1)
	fmt.Println(expression)
	err := calc(out, expression)
	if err != nil {
		fmt.Fprintf(out, "invalid syntax")
	}
}

func pop(stack *[]int) (int, []int) {
	return (*stack)[len(*stack)-1], (*stack)[:len(*stack)-1]
}

func calc(out io.Writer, expression string) error {
	var stack []int
	for _, char := range expression {
		var x, y int
		switch string(char) {
		case "+":
			{
				x, stack = pop(&stack)
				y, stack = pop(&stack)
				stack = append(stack, y+x)

			}
		case "-":
			{
				x, stack = pop(&stack)
				y, stack = pop(&stack)
				stack = append(stack, y-x)

			}
		case "*":
			{
				x, stack = pop(&stack)
				y, stack = pop(&stack)
				stack = append(stack, y*x)
			}
		case "/":
			{
				x, stack = pop(&stack)
				y, stack = pop(&stack)
				stack = append(stack, y/x)
			}

		case "=":
			{
				fmt.Fprintf(out, strconv.Itoa((stack[len(stack)-1])))
			}

		default:
			{
				number, error := strconv.Atoi(string(char))
				if error != nil {
					return error
				}
				stack = append(stack, number)
			}
		}

	}
	return nil
}
