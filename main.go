package main

import (
	"fmt"
)

/*type Rectangle struct {
	height int
		width  int
}

type MathOperation interface {
	Add(a, b int) int
	Area(a, b int) int
}

func (r Rectangle) Add(a, b int) int {
	res := r.height + r.width + a + b
	return res
}
func (r Rectangle) Area(a, b int) int {
	res := r.height * r.width * a * b
	return res
}*/

type CustomError struct {
	errorMessage string
}

func (err *CustomError) Error() string {
	return err.errorMessage
}

func main() {
	a := 2
	b := 3
	res, err := SumOfDigit(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("my name is vishal", res)
}

func SumOfDigit(a, b int) (int, error) {

	defer HandlePanic()
	if b > a {
		panic("Number is small ")
	}
	sum := a + b
	return sum, nil

	//panic("divison by zero")
}

func HandlePanic() {
	r := recover()
	if r != nil {
		fmt.Println("Handle Panic", r)
	}
}
