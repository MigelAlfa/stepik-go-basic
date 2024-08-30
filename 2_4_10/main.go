package main

import (
	"fmt"
)

func main() {
	var a,b,c,sum int
	fmt.Scan(&a,&b,&c)
	a = a*100
	sum = ((a+b)*c)

	fmt.Println(sum/100, sum%100)

}
