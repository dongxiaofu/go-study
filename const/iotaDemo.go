package main

const (
  a = iota
  b = iota
  c = iota
  d
  e = "a"
  f
  g = iota
  h
)

const (
  i = 1 << iota
  j = 3 << iota
  k
  l
)

func main(){
  println(a, b, c, d, e, f, g, h)
  println()
  println(i, j, k, l)
}
