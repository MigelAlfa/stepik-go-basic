package main

import (
"fmt"

    )

func main() {
    var a   int
	fact := 1
    fmt.Scan(&a)
	if a >= 2 && a <=15 {
    for i :=1 ; i <= a; i++ {
        if i%2 == 0 {
        fact = fact *i
        }
    }
}
fmt.Println(fact)
}
