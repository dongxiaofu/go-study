package main

import (
  "fmt"
)

type error interface {
  Error() string
}

type DivideError struct {
  dividee int
  divider int
}

// func (de *DivideError) Error() string {
//   strFormat := `
//   Cannot proceed, the divider is zero.
//   dividee : %d
//   divider : 0
// `
//
//   return fmt.Sprintf(strFormat, de.dividee)
// }


func (de *DivideError) Error() string {
    strFormat := "Cannot proceed, the divider is zero."

    return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
  if varDividee == 0 {
    data := DivideError {
      dividee : varDividee,
      // divider : varDivider
      divider : varDivider,
    }
    errorMsg = data.Error()
    return
  }else{
    return varDividee / varDivider, ""
  }
}

func main() {
  // if result,errorMsg := Divide(100, 10); errorMsg == "" {
  //   fmt.Println("100/10 = ", result)
  // }

  if _, errorMsg := Divide(100, 0); errorMsg != "" {
    fmt.Println("errorMsg is : ", errorMsg)
  }
}
