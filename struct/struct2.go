package main

import (
  "fmt"
)

type Books struct{
  title string
  author string
  id int
}

func printBooks(books Books){
  fmt.Printf("id = %d\n", books.id)
  fmt.Printf("title = %s\n", books.title)
  fmt.Printf("author = %s\n", books.author)
}

func main(){
  var books1 Books
  var books2 Books

  books1.title = "A"
  books1.author = "a"
  books1.id = 1

  books2.title = "B"
  books2.author = "b"
  books2.id = 2

  printBooks(books1)

  fmt.Println()

  printBooks(books2)


}
