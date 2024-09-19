package main

import (
    "fmt"
)

func main() {
    var n, sum, count float64
    fmt.Scan(&n)

    for n != 0 {
        sum = sum +n
        fmt.Scan(&n)

    }
    fmt.Println(sum/count)

}