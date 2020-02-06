package main

import "fmt"

func main() {
    var i float64
    fmt.Println("Please enter a float number")
    fmt.Scan(&i)
    fmt.Println("Truncated number", int64(i), "from stdin")
}