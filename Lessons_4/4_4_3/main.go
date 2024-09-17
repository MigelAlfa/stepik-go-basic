package main

import (
	"fmt"
)

func main() {
	var n, mnoz int
	fmt.Scan(&n)
	mnoz = 1

	i := 0

	for mnoz < n {
		if n == 1 {
			break
		} 
        mnoz = mnoz * 2
		i++
		
	}
	fmt.Println(i)
}
