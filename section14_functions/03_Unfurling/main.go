package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}

	foo(xi...)
}

func foo(x ...int) {
	for _, v := range x {
		fmt.Println("Value:", v)
	}
}
