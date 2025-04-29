package main

import (
	"fmt"
	"net/http"
	"os"
)


func readpage(name string) string {
	data , err := os.ReadFile(name)
	if err != nil{
	panic(err)
	}
	return string(data)
}




func index(web http.ResponseWriter , req *http.Request){
	fmt.Fprintf(web , readpage("assets/index.html"))
}




func main(){
	http.HandleFunc("/" , index)
	http.ListenAndServe("0.0.0.0:8000" , nil)
}
