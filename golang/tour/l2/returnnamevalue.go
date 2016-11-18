package main

import (
  "fmt"
)
func swap(x,y int)(z,w int){
  z=y
  w=x
  return
}

func main(){
  a,b:=swap(1,2)
  fmt.Println(a)
  fmt.Println(b)
}
