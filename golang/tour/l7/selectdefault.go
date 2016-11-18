package main

import (
	"time"
	"fmt"
	)

func fib(c chan int,quit chan int,o chan int){
	x,y:=0,1
	for {
		select {
			//从 c channel 读取数据
			case c <- x:
				x,y =y,x+y
			//从 quit channel 获取到了数据就退出
			case <- quit:
				fmt.Println("quit")
				return
			default:
				fmt.Println("    .")
				time.Sleep(50 * time.Millisecond)		
		}
	}
}

func main(){
	c :=make(chan int)
	quit :=make(chan int)
	o :=make(chan int)
	go func(){
		for i:=0;i<10;i++ {
			//计算阶乘
			fmt.Println(<-c)
		}
		//发送退出chanel
		o <-0
		quit <-0
	}()
	
	fmt.Println("计算阶乘")
	fib(c,quit,o)
	//fmt.Println(s,x,y)
}
