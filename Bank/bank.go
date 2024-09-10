package main

import (
	"fmt"
	"os"
	"strconv"
)

const fileName = "temp.txt"

func getBalance() float64 {
	data, _ := os.ReadFile(fileName)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}
func main() {
	var accountBalance = getBalance()
	fmt.Println("Welcome to Go Bank!")
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 2 {
			var deposit float64
			fmt.Println("Enter the amount to add: ")
			fmt.Scan(&deposit)
			accountBalance += deposit
			writeBalanceToFile(accountBalance)
		} else {
			return
		}

	}

}
