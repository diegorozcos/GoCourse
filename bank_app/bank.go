package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank.")
			return
		}
	}
}
