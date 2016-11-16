package main

import (
		"fmt"
       )


func main(){
	var a  [1]string
		a[0]="jingdor" 
		a[0]="alice" 
		fmt.Println(a[0])
		b:=[4]string{"jingdor","alice","viking","cgoer"}
		fmt.Println(b)
		fmt.Println(b[:2])
		fmt.Println(b[2:3])
		fmt.Println(b[2:])
		fmt.Println(b[0:0])
		s:=[]struct{
			name string
			age int
		}{
			{"jingdor",18},
			{"alice",20},
			{"cgoer",28},
		}
	
		fmt.Println(s)
}
