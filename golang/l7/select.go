package main

import (
	"fmt"
	)

func fib(c chan int,quit chan int){
	x,y:=0,1
	for {
		select {
			//从 c channel 读取数据
			case c <- x:
				x,y =y,x+y
			//从 quit channel 获取到了数据就退出
			case <-quit:
				fmt.Println("quit")
				return
		}
	}
}

func main(){
	c :=make(chan int);
	quit :=make(chan int);
	go func(){
		for i:=0;i<10;i++ {
			//计算阶乘
			fmt.Println(<-c)
		}
		//发送退出chanel
		quit <-0
	}()
	fib(c,quit)
	//fmt.Println(s,x,y)
}
