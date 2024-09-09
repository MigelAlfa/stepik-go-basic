package main

import (
	"fmt"
    "math"
)

func main() {
	var c1, c2, c3 int
	fmt.Scan(&c1, &c2, &c3)
    res1 := math.Ceil(float64(c1)/2)
    res2 := math.Ceil(float64(c2)/2)
    res3 := math.Ceil(float64(c3)/2)
    fmt.Println(res1+res2+res3)
     

    
}
