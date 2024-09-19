package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    plus := 0
    minus := 0


    for n != 0 {
        if n >= 0 {
        plus ++
        fmt.Scan(&n)

       } else if n <= 0 {
        minus ++
        fmt.Scan(&n)
       
    }
}
    fmt.Println(plus - minus)

}