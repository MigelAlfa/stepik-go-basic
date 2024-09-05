package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println("Следующее за числом", a, "число:", a+1)
	fmt.Println("Для числа", a, "предыдущее число:", a-1)


}