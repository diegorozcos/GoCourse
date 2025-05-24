package main

import "fmt"

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt float64, eat float64, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - taxRate/100)
	ratio = ebt / eat

	return ebt, eat, ratio
}

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, eat, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println("EBT:", ebt)
	fmt.Println("Profit:", eat)
	fmt.Printf("Ratio: %.2f", ratio)
}
