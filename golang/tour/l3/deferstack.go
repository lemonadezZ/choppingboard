package main

import "fmt"

func _exit(a int){  
  fmt.Println(a)
}

func main(){
  defer _exit(1)
  defer _exit(2)
  var a string="hello";
  fmt.Println(a)
}
