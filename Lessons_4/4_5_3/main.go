package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)

    for n > 0 { // пока цифры числа не закончатся
        digit := n % 10   // последняя цифра числа

        n = n / 10 // избавляемся от последней цифры.
        fmt.Print(digit)

    }

}

