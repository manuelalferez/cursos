package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calculator struct{}

func (calculator) calculate(operators string, operation string) int {
	numbers := strings.Split(operators, operation)
	firstOperator := parseInt(numbers[0])
	secondOperator := parseInt(numbers[1])

	switch operation {
	case "+":
		return firstOperator + secondOperator
	case "-":
		return firstOperator - secondOperator
	case "*":
		return firstOperator * secondOperator
	case "/":
		return firstOperator / secondOperator
	default:
		return 0
	}
}

func read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func parseInt(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}

func main() {
	fmt.Println("Introduce operation (+,-,* ó /): ")
	operation := read()
	fmt.Println("Introduce operators (ex: 2+2): ")
	operators := read()
	c := calculator{}
	fmt.Println("Result: ")
	fmt.Println(c.calculate(operators, operation))
}
