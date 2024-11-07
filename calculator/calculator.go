package main

import (
	"fmt"
)
var revenue, expenses, taxRate float64
func main(){
	fmt.Println(calculator())
}

func calculator()(float64, float64, float64 ) {
	
    fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

    fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

    fmt.Print("TaxRate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1- taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}