package main

import (
    "fmt"

)

func main() {
    var a int
    fmt.Scan(&a)
    if a >= 100 && a < 1000{
        fmt.Println(a/100)
    } else {
        fmt.Println(a/100%10)
    }    

}