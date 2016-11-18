package main

import (
	"fmt"
	"time"
	)

func say(s string){
	fmt.Println(s)
	time.Sleep(100)
}


func main(){
	go say("hello")
	say("world")
	for {
	
	}
}
