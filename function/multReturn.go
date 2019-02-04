package main

import "fmt"

func swap(var1 , var2 string) (string, string) {
	return var2, var1
}

func main(){
	a, b := swap("a", "b")
	fmt.Println(a, b)
}
