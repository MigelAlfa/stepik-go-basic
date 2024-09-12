


package main

import "fmt"

func main() {
    var a,sum,count int
    fmt.Scan(&a)
    for i :=1 ; i <= a; i++ { 
            fmt.Scan(&count)
            sum = sum + count
        }
fmt.Println(sum)
}
  