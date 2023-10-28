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

func main() {
	fmt.Println("hi my name is vishal")
	SlicePanic()
	fmt.Println("hi my name is borhade")
}

func SlicePanic() {
	defer HandlePanic()
	res := []int{1, 2, 3}
	fmt.Println("len", res[3])
	//defer HandlePanic()
	//panic("divison by zero")
}

func HandlePanic() {
	r := recover()
	if r != nil {
		fmt.Println("Handle Panic", r)
	}
}
