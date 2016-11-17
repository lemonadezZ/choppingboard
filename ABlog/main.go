package main 

import (
	"fmt"
	"time"
//	"strings"
	"net/http"
)

//发布

func log(response http.ResponseWriter,req *http.Request){
	t:=time.Now().Format("2006-01-02 15:04:05");
	id:="info";
	fmt.Printf("["+t+"] ["+id+"] ["+req.URL.Path+"] 处理log信息到日志\r\n")
	switch req.Method {
		case "GET":			
			response.Write([]byte(req.Method));	
		break;
		case "POST":
			response.Write([]byte(req.Method));	
		break;
		default:
			
		break;	
	}
}

func main() {
	http.HandleFunc("/",log)
	http.ListenAndServe(":8001",nil)	
}
