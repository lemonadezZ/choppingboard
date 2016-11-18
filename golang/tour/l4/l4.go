package main

import (
  "fmt"
)

func main(){
  var a *int
  var b int = 10
  a=&b
  fmt.Println(*a)



}
