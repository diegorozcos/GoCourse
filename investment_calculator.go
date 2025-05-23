package main

import (
	"fmt"
)

func main() {
	/*const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future Val:", futureValue)
	fmt.Println("Real Val:", futureRealValue)*/

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebtValue := revenue - expenses
	eatValue := ebtValue * (1 - taxRate/100)
	ratioValue := ebtValue / eatValue

	fmt.Println("EBT:", ebtValue)
	fmt.Println("EAT:", eatValue)
	fmt.Println("Ratio:", ratioValue)
}
