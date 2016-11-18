package main

import (
	"fmt"
	)

func sum(s []int ,c chan int){
	sum:=0
	for _ ,v := range s{
	 	sum+=v
	}
//	c<-sum
}

func loop(s []int ,c chan int){
	for{
	}
}

func main(){
	c:=make(chan int,2)
	c <- 1
	z:=<-c
	c <- 2
	c <- 3
	x, y := <-c ,<-c
	fmt.Println(z,x,y)
}
