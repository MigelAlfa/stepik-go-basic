package main

import (
    "fmt"
)

func main() {
    var a int
    fmt.Scan(&a)

    if a <= 13  {
        fmt.Println("детство")

    } else if a >= 14 && a <= 24 {
        fmt.Println("молодость")

    } else if a >= 25 && a <= 59 {
        fmt.Println("зрелость")
    } else {
        fmt.Println("старость")

    }

}