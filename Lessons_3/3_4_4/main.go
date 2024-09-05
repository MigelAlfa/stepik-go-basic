package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	c := a /100%10
	d := a /10%10
	e := a%10
	if a>100 && a <=999  {
		if c < d && d < e {
			fmt.Println("YES")
        } else {
            fmt.Println("NO")
        } 

	}

}
