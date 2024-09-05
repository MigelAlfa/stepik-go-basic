package main

import (
	"fmt"
	"math"
)

func main() {
    var a,b,c float64
    fmt.Scan(&a,&b,&c)
    d := ((b*b) - (4*a*c))
        if d < 0   {
            } else {
                if d == 0 {
                fmt.Print((-b)/(2*a))
                } else {
                    if d > 0{
                        x1 := (-float64(b)+math.Sqrt(float64(d))) / (2*float64(a))
                        x2 := (-float64(b)-math.Sqrt(float64(d))) / (2*float64(a))
                        if x1 < x2 {
                            fmt.Println(x1)
                            fmt.Println(x2)

                        }else {
                            fmt.Println(x2)
                            fmt.Println(x1)
                        }
                        }
                    }
            }               
    
        } 

    
    