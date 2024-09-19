package main

import (
	"fmt"
)

func main() {
	var n, sum, count int
	fmt.Scan(&n)

	for n != 0 {
		sum = n
		fmt.Scan(&n)
		if n > sum && sum!=0 {
            count++
		}
		
	}
	fmt.Println(count)

}
