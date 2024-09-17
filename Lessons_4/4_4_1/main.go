package main

import (
	"fmt"
)

func main() {
	var n, count, l int
	fmt.Scan(&n)
	i := 0
	l = n
	for i <= n {
		if l%3 == 0 && l != 0 {
			count ++
			l = l / 3
		}
        i++

	}
    fmt.Println(count)

}
