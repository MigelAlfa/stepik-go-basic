package main

import (
    "fmt"
)

func main() {
    n := 6808
    cnt := 1
    for n > 0 {
        cnt = cnt + 1
        n = n / 10

    }
	fmt.Println(cnt)
}