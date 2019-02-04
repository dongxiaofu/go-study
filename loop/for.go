package main

import "fmt"

func main(){
  var a int = 20
  var b int

  numbers := [5]int{1,2,3}

  for b := 0; b < 10; b++ {
    fmt.Printf("b = %d\n", b)
  }

  for(b < a){
    b++
    fmt.Printf("b = %d\n", b)
  }

  for k,v := range numbers {
    fmt.Printf("k = %d, v = %d\n", k, v)
  }
}
