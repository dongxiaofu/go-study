package main

import (
  "fmt"
)

func main(){
  var countryCapital map[string]string
  countryCapital = make(map[string]string)

  countryCapital["China"] = "北京"
  countryCapital["Japan"] = "东京"
  countryCapital["America"] = "华盛顿"
  countryCapital["France"] = "巴黎"

  for country := range countryCapital {
    fmt.Println(country + "首都是" + countryCapital[country])
  }

  fmt.Println()

  for country,capital := range countryCapital {
    fmt.Println(country + "首都是" + capital)
  }

  capital, ok := countryCapital["India"]
  if ok {
    fmt.Printf("India 的首都是 %s\n", capital)
  }else{
    fmt.Println("没有India")
  }

  var capital2 string = countryCapital["China"]
  fmt.Printf("capital = %s\n", capital2)

  capital3 := countryCapital["China"]
  fmt.Printf("capital3 = %s\n", capital3)
  fmt.Println()
  // missing second (key) argument to delete
  // delete(countryCapital["China"])
  delete(countryCapital, "China")

  for country := range countryCapital {
    fmt.Println(country, "的首都是" ,countryCapital[country])
  }
}
