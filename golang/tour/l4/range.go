package main

import (
		"fmt"
       )


func main(){
	a:=make([]int,5)
	a=append(a,2)
	a=append(a,2,3,4,5,6,7)
	for i,v:=range a {
		
		fmt.Printf("index: %d:",i)
		fmt.Printf("val: %d \n",v)
	}
	for _,v:=range a {
		
		fmt.Printf("val: %d \n",v)
	}
	fmt.Println(a)
}
