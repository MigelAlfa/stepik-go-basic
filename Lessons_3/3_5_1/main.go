package main

import (
	"fmt"

)

func main() {
	var m  int
    var m1 string
	fmt.Scan(&m, &m1)
    if m >=12 && m <=18 && m1 == "m"{
        fmt.Println("YES")
    }    else {
        fmt.Println("NO")
    }
    
}
