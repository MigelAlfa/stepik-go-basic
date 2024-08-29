package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    _ = scanner.Scan()
    line1 := scanner.Text()
    _ = scanner.Scan()
    line2 := scanner.Text()
    _ = scanner.Scan()
    line3 := scanner.Text()
    _ = scanner.Scan()
    line4 := scanner.Text()
    
    fmt.Println(line2 +line1+ line3+line1 + line4)
    
}