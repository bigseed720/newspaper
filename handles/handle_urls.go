package handles

import (
	"fmt"
	"net/http"

)




func index(web http.ResponseWriter , req *http.Request){
    fmt.Fprintf(web , ReadPage("assets/index.html"))
}

func file(web http.ResponseWriter , req *http.Request){
	fmt.Fprintf(web , ReadPage("assets/index.html"))
	web.Header().Set("Content-Type","application/x-www-form-urlencoded")
	req.ParseForm()
  	fmt.Println(req.Form["search"][0])
  	fmt.Println(req.Form["search"][0])
	
}

func Handle() {
	http.HandleFunc("/" , index)
	http.HandleFunc("/search" , file)
}
