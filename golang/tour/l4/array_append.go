package main

import (
		"fmt"
       )


func main(){
	a:=make([]int,5)
	a=append(a,2)
	a=append(a,2,3,4,5,6,7)
	fmt.Println(a)
}
