package main

import (
	"fmt"

)

func main() {
	var m  int

	fmt.Scan(&m)
    if m%2 != 0 {
        fmt.Println("YES")
    }    else if m%2 == 0 && m >=2 && m<=5 {
        fmt.Println("NO")
    } else if m%2 == 0 && m >=6 && m<=20 {
        fmt.Println("YES")
    } else if m%2 == 0 && m >20 {
        fmt.Println("NO")
    }
    
}
