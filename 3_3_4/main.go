package main

import (
    "fmt"
    "strconv"
)

func main() {
    var a,b int
    var c string
    fmt.Scan(&a,&b,&c)
    _, err := strconv.Atoi(c)
    if err == nil {
		fmt.Print("Неверная операция")
    } else {
        if c == "/" && b == 0   {
            fmt.Print("На ноль делить нельзя!")
            } else {
                if c == "+"{
                fmt.Print(a+b)
                } else {
                    if c == "-"{
                        fmt.Print(a-b)
                    } else {
                        if c == "*"{
                            fmt.Print(a*b)
                        } else {
                            if c == "/"{
                                fmt.Print(float64(a)/float64(b))
                            }
                        }
                    }
            }               
        
        } 
    
    }
}