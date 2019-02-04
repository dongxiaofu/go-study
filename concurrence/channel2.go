package main

import (
  "fmt"
)

func sum(s []int, ch chan int){
  var sum int = 0
  for _, v := range s {
    sum += v
  }

  ch <- sum
}

func main(){
  nums := []int{7, 2, 8, -9, 4, 0}

  ch := make(chan int,2)
  go sum(nums[len(nums)/2:], ch)
  x := <- ch
  go sum(nums[:len(nums)/2], ch)


  y := <- ch

  fmt.Printf("x = %d, y = %d\n", x, y)

}
