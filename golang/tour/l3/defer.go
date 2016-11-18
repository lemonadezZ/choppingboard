package main

import "fmt"

func _exit(){  
  fmt.Println("world")
}

func main(){
  defer _exit()
  var a string="hello";
  fmt.Println(a)
}
