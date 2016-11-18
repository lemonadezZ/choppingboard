package main

import "fmt"


type I interface {
	getname() string
	getage() int
}

type T struct {
	name string
	age int
}


func (p T)getage()int{
	fmt.Println(p.age)
	return p.age
}


func (p T)getname()string{
	fmt.Println(p.age)
	return p.name
}

func main(){
  	var p I=T{"jingdor",18}
	p.getage()
	fmt.Println(p)
}
