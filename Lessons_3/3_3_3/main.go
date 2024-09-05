package main

import "fmt"

func main() {
    var a,b float32
    fmt.Scan(&a,&b)
    if a >0 && b > 0   {
        fmt.Print(1)
    } else {
        if a <0 && b>0 {
            fmt.Print(2)
            } else {
                if a <0 && b<0 {
                    fmt.Print(3)
                }else {
                    fmt.Print(4)
                }                
        }
    }
}