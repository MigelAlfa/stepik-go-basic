package main
import (
    "fmt"
)

func main() {
    var d1, d2, d3 int
    _, _ = fmt.Scan(&d1, &d2, &d3)
    if d3 >= d1+d2 {
        fmt.Print(2 * (d1 + d2))
    } else if d2 >= d1+d3 {
        fmt.Print(2 * (d1 + d3))
    } else if d1 >= d2+d3 {
        fmt.Print(2 * (d2 + d3))
    } else {
        fmt.Print(d1 + d2 + d3)
    }
}