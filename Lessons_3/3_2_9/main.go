package main

import (
    "fmt"
)

func main() {
    var a,b,c,d int
    fmt.Scan(&a,&b,&c,&d)

    if a == c || b == d {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    } 
}