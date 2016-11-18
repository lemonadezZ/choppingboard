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

func getage1(p person)int{
	fmt.Println(p.age)
	return p.age
}

func main(){
  	p:=person{"jingdor",18}
	p.getage()
	getage1(p)
	fmt.Println(p)
}
