package main
import (
    "fmt"
    "strings"
)

func main() {
    var i string

    fmt.Println("Please enter a string")
    fmt.Scan(&i)

    if strings.HasPrefix(strings.ToLower(i), "i") && strings.Contains(strings.ToLower(i), "a") && strings.Contains(strings.ToLower(i), "a"){
fmt.Println("Found")
     }else
     {
fmt.Println("Not Found")
     }
    
}