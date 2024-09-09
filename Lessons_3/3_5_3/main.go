package main

import (
	"fmt"

)

func main() {
	var m  int

	fmt.Scan(&m)
    if m > -30 && m <= -2 || m > 7 && m <= 25 {
        fmt.Println("Принадлежит")
    }    else {
        fmt.Println("Не принадлежит")
    }
    
}
