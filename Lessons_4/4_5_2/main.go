package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    s := n
    sum := 0    // сумматор
    for n > 0 { // пока цифры числа не закончатся
        digit := n % 10   // последняя цифра числа
        sum = sum + digit // складываем в сумматор
        n = n / 10 // избавляемся от последней цифры.
  
        }
        if s % sum == 0 {
        fmt.Println("YES")

        } else {
        fmt.Println("NO") 

        }
}

