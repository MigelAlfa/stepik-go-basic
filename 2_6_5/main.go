package main

import "fmt"

func main() {
	var a float64
	var b int
	fmt.Scan(&a)
	b = int(a)

	fmt.Println(a-(float64(b)))


}