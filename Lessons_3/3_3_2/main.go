package main

import "fmt"

func main() {
    var a int
    fmt.Scan(&a)
    if a <60   {
        fmt.Print("Легкий вес")
    } else {
        if a<64 {
            fmt.Print("Первый полусредний вес")
            } else {
                fmt.Print("Полусредний вес")

        }
    }
}