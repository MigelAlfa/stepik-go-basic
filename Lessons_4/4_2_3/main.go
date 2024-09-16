
package main

import "fmt"

func main() {
    var a,sum,count int
    fmt.Scan(&a)
    for i :=1 ; i <= a; i++ { 
            fmt.Scan(&count)
            if count%2 == 0 && count%3 != 0{
            sum = sum + count
            }
        }
fmt.Println(sum)
}
  