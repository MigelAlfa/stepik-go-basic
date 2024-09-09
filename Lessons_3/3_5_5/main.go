package main

import (
	"fmt"
    
)

func main() {
	var m  float64

	fmt.Scan(&m)
    if m <=3  {
        fmt.Println("начинающий")
    }    else if m > 3 && m < 8 {
        fmt.Println("младший разработчик")
    }    else if m > 7 && m < 16 {
        fmt.Println("средний разработчик")
    } else if m > 15  {
        fmt.Println("старший разработчик")
    }   
}
