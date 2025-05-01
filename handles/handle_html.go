package handles

import (
	"os"

)


func ReadPage(name string) string {
    data , err := os.ReadFile(name)
    if err != nil{
    panic(err)
    }
    return string(data)
}
