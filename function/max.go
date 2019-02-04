package main

import "fmt"

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	}else{
		result = num2
	}

	return result
}

func printVar(parameter1 string, parameter2 int) {
	// fmt.Printf(parameter1, parameter2)
	fmt.Println(parameter1, parameter2)
}

func main(){
	var max int = max(3,9)

	fmt.Printf("max is %d\n", max)

	printVar("hello", 3)

}
