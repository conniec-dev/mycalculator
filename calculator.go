package mycalculator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Calc struct{}

func (Calc) Operate(input string, operation string) (int, error) {
	cleanInput := strings.Split(input, operation)
	operator1 := ParseString(cleanInput[0])
	operator2 := ParseString(cleanInput[1])
	switch operation {
	case "+":
		fmt.Println(operator1 + operator2)
		return operator1 + operator2, nil
	case "-":
		fmt.Println(operator1 - operator2)
		return operator1 - operator2, nil
	case "*":
		fmt.Println(operator1 * operator2)
		return operator1 * operator2, nil
	case "/":
		fmt.Println(operator1 / operator2)
		return operator1 / operator2, nil
	default:
		log.Println(operation, "operation is not supported!")
		return 0, nil

	}
}

func ParseString(input string) int {
	operator, _ := strconv.Atoi(input)
	return operator
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
