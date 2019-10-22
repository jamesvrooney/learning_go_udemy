package main

import "fmt"

func main() {
	a := "James"
	num := 45

	foo()
	bar(a)
	add(&num)

	fmt.Println("My name is still", a)
	fmt.Println("Main: The num:", num)
}

func foo() {
	fmt.Println("My name is James")
}

func bar(s string) {
	s = "John"

	fmt.Println("My name is", s)
}

func add(x *int) {
	fmt.Println("In add func - num:", *x)

	*x++

	fmt.Println(*x)
}
