package calculate

import "fmt"

type Number struct {
	num1  float64
	num2  float64
	input string
}

func NewNumber() *Number {
	return &Number{}
}

func (n *Number) SetNum1(num1 float64) {
	n.num1 = num1
}
func (n *Number) GetNum1() float64 {
	return n.num1
}
func (n *Number) SetNum2(num2 float64) {
	n.num2 = num2
}

func (n *Number) GetNum2() float64 {
	return n.num2
}

func (n *Number) SetInput(input string) {
	n.input = input
}

func (n *Number) GetCalc(input string) string {
	var result float64

	switch input {
	case "1": // Add
		result = n.GetNum1() + n.GetNum2()
	case "2": // Subtract
		result = n.GetNum1() - n.GetNum2()
	case "3": // Multiply
		result = n.GetNum1() * n.GetNum2()
	case "4": // Divide
		if n.num2 != 0 {
			result = n.GetNum1() / n.GetNum2()
		} else {
			return "Cannot divide by zero"
		}
	case "5":
		result = n.GetNum1() / 100
	case "6":
		return "Exit"
	default:
		return "Invalid operation"
	}

	return fmt.Sprintf("%f", result)
}
