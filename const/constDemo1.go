package main

import "fmt"

func main(){
  const WIDTH int = 5
  const HEIGHT int = 4
  var area int
  var a, b, c = 3, "hi", true

  area = WIDTH * HEIGHT
  fmt.Printf("area is : %d", area)

  println()

  println(a, b, c)
}
