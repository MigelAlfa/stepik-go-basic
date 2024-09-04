package main

import "fmt"

func main() {
    var a int
    fmt.Scan(&a)
    if a >= -3 && a<=1  {
        fmt.Print("YES")
    } else {
        if a>=5 && a<=9 {
            fmt.Print("YES")
            } else {
                fmt.Print("NO")

        }
    }
}