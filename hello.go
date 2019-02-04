package main

import "fmt"

var (
	a int
	b bool
)

func main() {
	var str string = "hello, world + str"
	fmt.Println(str)

	printVar1()

	printVar2()

	printVar3()

	printVar4()

	a = 1
	b = true

	println(a, b)
}

func printVar1() {
	var a string = "study go"
	var b = "study c"
	var c bool

	println(a, b, c)
}

func printVar2() {
	a := "hello"
	var b = "hello"
	// var c string = "hello"

	println(a)
	println(b)
}

func printVar3() {
	var a, b, c = "how", "are", "you"
	println(a, b, c)
	d, e, f := "Fine,", "thank", "you"
	println(d, e, f)
}

func printVar4() {
	var (
		a int
		b bool
	)

	a = 3
	b = false

	println(a, b)
}
