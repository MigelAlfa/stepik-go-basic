
package main

import "fmt"

func main() {
	var a, sum, count int
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		fmt.Scan(&count)
		if count == 0 {
			sum ++
		}
	}
    if sum == 0 { 
    fmt.Println("NO")
    } else {
    fmt.Println("YES")
    }
}