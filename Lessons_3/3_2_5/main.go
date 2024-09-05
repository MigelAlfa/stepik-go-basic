package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    a := n/1000

    b := a /100 // 1ое
    c := a/10%10 // 2ое
    d := a%10 // 3ее

    x := n%1000
    
    k := x /100 // 1ое
    l := x/10%10 // 2ое
    m := x%10 // 3ее

    if (b+c+d) == (k+l+m) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    } 
}