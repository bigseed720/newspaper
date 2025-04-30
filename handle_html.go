package main

import (
	"os"

)


func readpage(name string) string {
    data , err := os.ReadFile(name)
    if err != nil{
    panic(err)
    }
    return string(data)
}
