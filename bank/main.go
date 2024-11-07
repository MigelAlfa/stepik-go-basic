package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalance = "Balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalance)

	 if err != nil {
	 	return 0, errors.New("No balance file")
	 }

	balanceTxt := string(data)
	balance, err := strconv.ParseFloat(balanceTxt, 64)

	if err != nil {
		return 0, errors.New("Faild to parse storage file value")
	}

	return balance, err
}

func writeBalanceInFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(accountBalance, []byte(balanceTxt), 0644)

}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("___________")
	}
	fmt.Println("Welcome to GO bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Chack the balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choise int
		fmt.Print("Your coise is: ")
		fmt.Scan(&choise)

		if choise == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choise == 2 {
			fmt.Print("How much you whant to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount")
				//return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updeted. New amoutn is: ", accountBalance)
			writeBalanceInFile(accountBalance)

		} else if choise == 3 {
			fmt.Print("How much you whant to Withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if accountBalance < withdrawAmount {
				fmt.Println("You cann not withdraw more than you have")
				//returne
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updeted. New amoutn is: ", accountBalance)
			writeBalanceInFile(accountBalance)

		} else {
			fmt.Println("Exiting from app")
			break
		}

	}
	fmt.Println("Tank you for choosing our bank")
}
