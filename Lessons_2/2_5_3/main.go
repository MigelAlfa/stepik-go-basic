package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	if a <999 || a >= 100{
    b:= a/100 // 1ое
	
	c:= a/10%10 // 2ое

	d:= a%10 // 3ee

	fmt.Println(b+c+d)
}


}