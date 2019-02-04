package main

import "unsafe"

const (
  a = "abc"
  b = len(a)
  c = unsafe.Sizeof(a)
  d
)

func main(){
  println(a, b, c, d)
}
