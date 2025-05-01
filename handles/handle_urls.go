package handles

import (
	"fmt"
	"net/http"

)




func index(web http.ResponseWriter , req *http.Request){
    fmt.Fprintf(web , ReadPage("assets/index.html"))
}


func Handle() {
	http.HandleFunc("/" , index)
}
