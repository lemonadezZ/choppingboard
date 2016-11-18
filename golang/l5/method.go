package main

import "fmt"


type person struct {
	name string
	age int
}


func (p person)getage()int{
	fmt.Println(p.age)
	return p.age
}


func main(){
  	p:=person{"jingdor",18}
	p.getage()
	fmt.Println(p)
}
