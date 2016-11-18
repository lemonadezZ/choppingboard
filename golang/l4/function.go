package main

import (
		"fmt"
       )


func main(){
	f:=func(a int)(int){
		return a
	}
	fmt.Println(f(11))
}
