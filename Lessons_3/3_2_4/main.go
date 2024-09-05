package main

import (
    "fmt"
)

func main() {
    var a int
    fmt.Scan(&a)
    b := a /100 // 1ое
    c := a/10%10 // 2ое
    d := a%10 // 3ее
    if b != c && b != d && c != d {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    } 
}