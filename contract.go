package calculate

type NumberInterface interface {
	SetNum1(num1 float64)
	GetNum1() float64
	SetNum2(num2 float64)
	GetNum2() float64
	SetInput(input string)
	GetCalc(input string) string
}
