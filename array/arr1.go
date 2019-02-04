package main

import (
	"fmt"
)

func main(){
	// var arr [5]int = {1,2,3,4,5}
	var arr  = [5]int{1,2,3,4,5}
	fmt.Printf("arr[3] = %d\n", arr[3])

	// var arr2 int[10]
	var arr2 [10]int

	for i := 0; i < 10; i++ {
		arr2[i] = i
	}

	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d] = %d\n", arr2[i], arr2[i])
	}
}
