package main

import (
	"fmt"
	"sync"
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

type abc struct {
	x int
	y int
}

var mu = sync.Mutex{}
var money = 100

func debit() {
	for i := 1; i <= 1000; i++ {
		mu.Lock()
		money = money + 10
		mu.Unlock()
	}
	time.Sleep(1 * time.Millisecond)
	fmt.Println("debit done..")
}

func withdraw() {

	for i := 1; i <= 1000; i++ {
		mu.Lock()
		money = money - 10
		mu.Unlock()
	}
	time.Sleep(1 * time.Millisecond)
	fmt.Println("withdraw done..")

}

func main() {
	go debit()
	go withdraw()
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Money...", money)
}

func SwapDigit(i, j *int) {
	*i, *j = *j, *i
}

func HandlePanic() {
	r := recover()
	if r != nil {
		fmt.Println("Handle Panic", r)
	}
}
