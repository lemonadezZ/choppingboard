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
  fmt.Println(person1.name)
  fmt.Println(person1.age)



}
