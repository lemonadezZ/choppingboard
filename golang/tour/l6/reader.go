package main

import (
	"fmt"
	"io"
	"strings"

)

func main(){
  a:= strings.NewReader("hello world");
	
  b:=make([]byte,8)
	for{
		n,err:=a.Read(b);
		fmt.Printf("n:%v  error:%v  b :%v \n",n,err,b)
		fmt.Printf("%q\n" , b[:n])
		if err == io.EOF{
			break
		}
	}
}
