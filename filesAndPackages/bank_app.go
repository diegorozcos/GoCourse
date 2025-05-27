package main

import (
	"fmt"

	"example.com/bank-app/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("------------------")
		// panic("Can't continue.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7:", randomdata.PhoneNumber())
	for {
		presentOptions()

		var choice int

		fmt.Print("Choose an option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var deposit float64
			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Deposit must be greater than 0")
				continue
			}
			accountBalance += deposit
			fmt.Println("Your new account balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			var withdraw float64
			fmt.Print("Enter the amount you want to withdraw: ")
			fmt.Scan(&withdraw)
			if withdraw <= 0 {
				fmt.Println("Deposit must be greater than 0")
				continue
			}
			if withdraw > accountBalance {
				fmt.Println("Invalid ammount. You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdraw
			fmt.Println("Your new account balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank.")
			return
		}
	}
}
