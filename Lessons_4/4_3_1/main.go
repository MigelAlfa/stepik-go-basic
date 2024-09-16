
package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    fact := 1 
    if n <= 12 {
    for i := 1; i <= n; i++ {
        fact = fact * i 
    }
}
    fmt.Println(fact)
}