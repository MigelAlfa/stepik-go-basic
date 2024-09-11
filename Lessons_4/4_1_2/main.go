
package main

import "fmt"

func main() {
    var a,b int
    fmt.Scan(&a,&b)
    if a%2 ==0 {
    for; a <= b; a=a+2 { 
        fmt.Println(a)
    }
    } else {
     for a := a+1; a <= b; a=a+2 {
        fmt.Println(a)
     }
    }
}
  