package main

import "fmt"

func main(){
  var a int = 4
  var b int32
  var c float32
  var ptr *int

  ptr = &a

  fmt.Printf("a 的数据类型是：%T\n", a)
  fmt.Printf("b 的数据类型是：%T\n", b)
  fmt.Printf("c 的数据类型是： %T\n", c)
  fmt.Printf("ptr = %d\n", *ptr)
}
