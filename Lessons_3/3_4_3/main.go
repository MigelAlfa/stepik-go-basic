package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	b := a / 1000
	c := a /100%10
	d := a /10%10
	e := a%10
	if a>1000 && a <=9999  {
		if b == e && c == d {
			fmt.Println("YES")
        } else {
            fmt.Println("NO")
        } 

	}

}
