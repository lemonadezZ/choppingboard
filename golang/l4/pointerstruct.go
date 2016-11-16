package main

import (
  "fmt"
)

type person struct {
  name string
  age int
}

func main(){
  
  var person1=person{"jingdor",18}
  p	:=&person1
  p.age=20
  fmt.Println(person1.age)



}
