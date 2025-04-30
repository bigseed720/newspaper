package main

import (
	"fmt"
	"net/http"

)



func index(web http.ResponseWriter , req *http.Request){
    fmt.Fprintf(web , readpage("assets/index.html"))
}


func handle() {
	http.HandleFunc("/" , index)
}
