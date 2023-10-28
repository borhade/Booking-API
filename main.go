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

type Student struct {
	Id    int    `json:"Id"`
	Name  string `json:"Name"`
	Posts []StudentPost
}

type StudentPost struct {
	PostId   int    `json:"PostId"`
	PostName string `json:"PostName"`
	Comments string `json:"Comments"`
}

func main() {

	data := make([]map[int]string, 0)
	for i := 0; i < 5; i++ {
		test := map[int]string{
			i: fmt.Sprintf("Alice is:%d", i),
		}
		data = append(data, test)
	}

	fmt.Println(data)
	//b := 300
	//fmt.Println("address of a ", &a)
	//fmt.Println("address of a ", &b)
	//SumOfDigit(&a)
	//fmt.Println("after function address of a ", a)
	//fmt.Println("address of a ", b)
}

func SumOfDigit(a *int) {
	*a *= *a
	//fmt.Println("address of a...", *a)
}

func HandlePanic() {
	r := recover()
	if r != nil {
		fmt.Println("Handle Panic", r)
	}
}
