package main

import (
	"fmt"
    "math"

)

func main() {
	var m  float64

	fmt.Scan(&m)
    if m == 0 {
        fmt.Println("Обратного числа не существует")
    }    else {
        fmt.Println(math.Pow(m, -1))
    }
    
}
