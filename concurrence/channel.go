package main

import "fmt"

func sum(s []int, ch chan int){
	sum := 0
	for _, v := range s {
		sum += v
	}

	ch <- sum
}

func main(){
	nums := []int{1,2,3,4,5,6}

	ch := make(chan int)
	go sum(nums, ch)
	x := <- ch
	fmt.Println(x)
	fmt.Printf("nums : %v\n", nums)
	// fmt.Println(string(x))
	// fmt.Println("x = " + string(x))
	go sum(nums[:len(nums)/2], ch)
	y := <- ch
	fmt.Println(y)
	fmt.Printf("nums[:len[num/2]] : %v\n", nums[:len(nums)/2])

	go sum(nums[len(nums)/2:], ch)
	z := <- ch
	fmt.Println(z)
	fmt.Printf("nums[len(nums/2):] : %v\n", nums[len(nums)/2:])

	// 切片，endIndex缺省时，从startIndex到len(nums)
	// 无缺省时，包括startIndex，不包括endIndex
	fmt.Printf("nums[2,3] : %v\n", nums[2:3])
	fmt.Printf("nums[2:] : %v\n", nums[2:])
	fmt.Printf("nums[:2] : %v\n", nums[:2])
}
