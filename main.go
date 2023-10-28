package main

import (
	"fmt"
	"time"
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
	ch1 := make(chan int)
	ch2 := make(chan string)
	go second(ch2)
	go first(ch1)

	fmt.Println("hi this gorutine in goalng")
	select {
	case val1 := <-ch1:
		fmt.Println("Yes I am execute", val1)
	case val2 := <-ch2:
		fmt.Println("I am second", val2)
		//default:
		//fmt.Println("Yes I am default case")
	}

	//b := 300
	//fmt.Println("address of a ", &a)
	//fmt.Println("address of a ", &b)
	//SumOfDigit(&a)
	//fmt.Println("after function address of a ", a)
	//fmt.Println("address of a ", b)
}

func first(ch1 chan int) {
	time.Sleep(4 * time.Second)
	ch1 <- 23
	//fmt.Println("My name is vishal")\
	//fmt.Println("address of a...", *a)
}

func second(ch2 chan string) {
	time.Sleep(2 * time.Second)
	ch2 <- "Hello, Go!"
}

func HandlePanic() {
	r := recover()
	if r != nil {
		fmt.Println("Handle Panic", r)
	}
}
