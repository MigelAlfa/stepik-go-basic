package main
import (
	"fmt"
     "math"
	)

func main() {
    var b float64
    fmt.Scan(&b)
    
    fmt.Println(b / math.Pow(2, 13))    
}