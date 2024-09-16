package main

import "fmt"

func main() {
    var a ,b  int
	fact := 1
    fmt.Scan(&a, &b)
	if a <= b && a < 100 && b <100 {
    for  ; a <= b; a++ {
        fact = fact * a 

    }
}
fmt.Println(fact)
}
