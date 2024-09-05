package main

import "fmt"

func main(){
    var a,b int
    fmt.Scan(&a,&b)
    if b !=0 {
        if (a%b) == 0  {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
        
    }

}
