package main

import (
	"fmt"

)

func main() {
	var a1  int
	fmt.Scan(&a1)
    if a1 >= 1 && a1 <=100{
        if a1 %2 == 0 && a1 !=2{
        fmt.Println("YES")

    }    else {
        fmt.Println("NO")
    }
    }
     

}
