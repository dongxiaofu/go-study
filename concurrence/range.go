package main

import (
  "fmt"
)

func fobonacci(n int, ch chan int){
  x,y := 0,1

  for i := 0; i < n; i++ {
    // invalid operation: <-x (receive from non-chan type int)
    // ch := <- x
    ch <- x
    x,y = y,x + y
  }

  close(ch)
}

func main(){
  ch := make(chan int, 19)

  var num int = 8;
  go fobonacci(num, ch)

  // too many variables in range
  // for _, v := range ch {
  for v := range ch {
    fmt.Println(v)
  }
}
