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
	s:=[]int{1,2,3,4}
	c:=make(chan int,3)
	go loop(s[:len(s)/2],c)
	go sum(s[:len(s)/2],c)
	//go sum(s[len(s)/2:],c)
	x, y := <-c ,<-c
	fmt.Println(s,x,y)
}
