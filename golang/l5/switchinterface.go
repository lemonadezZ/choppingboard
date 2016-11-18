package main

import "fmt"

func do(i interface{}){
	switch v := i.(type) {
		case int:
			fmt.Printf("type %T \n",v,v)
		case string:
			fmt.Printf("type %T \n",v,v)
		case bool:
			fmt.Printf("type %T \n",v,v)
		default:
			fmt.Printf("type %T \n",v,v)
			
	}
}

func main(){
	do(11)
	do("11")
	do(true)
	do(11.1)
}
