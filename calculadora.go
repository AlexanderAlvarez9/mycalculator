package mycalculator

import "fmt"

type calc struct {
}

func (calc) operate(inputsValue [2]int, operator string) {
	firstNumber := inputsValue[0]
	secondNumber := inputsValue[1]
	switch operator {
	case "+":
		fmt.Printf("Result of %d %s %d = %d\n", firstNumber, operator, secondNumber, (firstNumber + secondNumber))

	case "-":
		fmt.Printf("Result of %d %s %d = %d\n", firstNumber, operator, secondNumber, (firstNumber - secondNumber))

	case "/":
		fmt.Printf("Result of %d %s %d = %d\n", firstNumber, operator, secondNumber, (firstNumber / secondNumber))

	case "*":
		fmt.Printf("Result of %d %s %d = %d\n", firstNumber, operator, secondNumber, (firstNumber * secondNumber))

	default:
		fmt.Println("Really nigga?", firstNumber, operator, secondNumber)
	}
}

// GetData get data from user
func getData() (int, int, string) {
	var firstNumber, secondNumber int
	var operator string

	fmt.Printf("Operation: + - / *\n")
	fmt.Scan(&operator)

	fmt.Printf("Primer valor\n")
	fmt.Scan(&firstNumber)

	fmt.Printf("Segundo valor\n")
	fmt.Scan(&secondNumber)
	return firstNumber, secondNumber, operator
}

func calculator() {
	print("Welcome to goCalc\n")
	v1, v2, op := getData()
	c := calc{}
	values := [2]int{v1, v2}
	c.operate(values, op)
}

func main() {
	calculator()
}
