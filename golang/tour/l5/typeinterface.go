package main

import "fmt"


func main(){

	var i interface{}="hello world"

	fmt.Println(i.(string));
	fmt.Println(i.(type));
}
