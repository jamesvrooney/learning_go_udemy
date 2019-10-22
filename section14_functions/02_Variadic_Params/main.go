package main

import "fmt"

func main() {
	foo(1, 2, 3, 4, 5)
}

func foo(x ...int) {
	fmt.Printf("%T\n", x)

	for i, val := range x {
		fmt.Println("Value at index:", i, "is", val)
	}
}
