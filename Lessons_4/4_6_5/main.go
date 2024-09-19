package main

import "fmt"

func main() {
    var n, m, count int
    fmt.Scan(&n)
    for n != 0 {
    m = n
    fmt.Scan(&n)
        if m < 0 && n > 0 || m > 0 && n < 0 {
          count++
      }
    }
    fmt.Println(count)
}