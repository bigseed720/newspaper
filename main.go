package main

import (
	"net/http"
	"newspaper/handles"
)







func main(){
	handles.Handle()
	http.ListenAndServe("0.0.0.0:8000" , nil)
}
