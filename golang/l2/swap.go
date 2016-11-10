package main

import "fmt"

func swap(x ,y int) (int,int) {
  return y,x
}


func swapstring(x ,y string) (string,string) {
  return y,x
}
func main(){
  a,b:=swap(1,2)
  c,d:=swapstring("1","2")
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
