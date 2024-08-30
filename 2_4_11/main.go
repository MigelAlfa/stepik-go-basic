package main

import (
	"fmt"
)

func main() {
	var a,sum,h int
	fmt.Scan(&a)
	h = a/60

	sum = a%60

	fmt.Println(a, "мин - это", h, "час", sum, "минут.")

}
