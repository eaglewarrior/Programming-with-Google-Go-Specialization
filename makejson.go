package main

import (
    "encoding/json"
    "fmt"
)

func main() {
var name string
var addr string
fmt.Println("Please enter your name")
fmt.Scan(&name)
fmt.Println("Please enter your address")
fmt.Scan(&addr)
var m map[string]string
m = make(map[string]string)
m["name"]=name
m["address"]=addr
j, _:= json.Marshal(m)
fmt.Println(string(j))

}