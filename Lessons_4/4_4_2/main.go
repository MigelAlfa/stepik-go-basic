package main

import (

    "fmt"

)

func main() {
    var n, step int
    fmt.Scan(&n)
    step = 1

    i := 0
    for step <= n {
        if step <= n {
       fmt.Println(step)
		}	
		step = step * 2
        i++
    }

}