package main

import (
		"fmt"
       )


func main(){
	var a  [1]string
		a[0]="jingdor" 
		a[0]="alice" 
		fmt.Println(a[0])
		b:=[10]string{"jingdor","alice","viking","cgoer"}
	
		fmt.Println(len(b))
		fmt.Println(cap(b))
		fmt.Println(b)
		fmt.Println(b[:2])
		fmt.Println(b[2:3])
		fmt.Println(b[2:])
		fmt.Println(b[0:0])
		fmt.Println(b[:])
}
