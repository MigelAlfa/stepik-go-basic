
package main

import "fmt"

func main() {
    var a int
    fmt.Scan(&a)
    count := 0
    for i :=1 ; i <= a; i++ { 
        if a % i == 0{
            count ++
        }
}
fmt.Println(count)
}
  