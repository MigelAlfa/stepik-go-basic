package main

import (
"fmt"
       "math"
    )

func main() {
    var a ,b  int
	fact := 1
    fmt.Scan(&a, &b)
	if a <= b && a < 150 && b <150 {
    for  ; a <= b; a++ {
        if math.Abs(float64(a%10)) == 7 {
        fact = fact * a 
        }
    }
}
fmt.Println(fact)
}
