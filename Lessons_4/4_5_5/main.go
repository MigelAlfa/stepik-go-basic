package main

import (
    "fmt"
)

func main() {
    var n, sum int
    fmt.Scan(&n)

    for n > 0 {
        sum += n%2
        n = n / 2
    }
    fmt.Println(sum)
}

