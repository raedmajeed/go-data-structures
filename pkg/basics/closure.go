package basics

import "fmt"

func checkClosure() {
	returnedFunc := closure()
	fmt.Println(returnedFunc(5))
}

func closure() func(int) int {
	sum := 10
	return func(i int) int {
		return sum + i
	}
}
