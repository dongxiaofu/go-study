package main

import "fmt"

func swap(a *int, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func main(){
	var a, b int = 100, 200
	fmt.Printf("before swap: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("after swap: a = %d, b = %d\n", a, b)
}
