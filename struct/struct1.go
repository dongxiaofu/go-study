package main

import (
  "fmt"
)

type Books struct{
  title string
  author string
  subject string
  id int
}

func main(){
  fmt.Println(Books{"go","cg","program",1})
  fmt.Println(Books{title : "go", author : "cg", subject : "program", id : 1})
  // ./struct1.go:17:26: too few values in Books literal
  // fmt.Println(Books{"go","cg"})
  // 有field时可缺省
  fmt.Println(Books{title : "go", id : 1})
}
