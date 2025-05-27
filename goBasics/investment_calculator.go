package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)

	return fv, frv
}

func main() {
	investmentAmount := getUserInput("Investment Amount: ")
	expectedReturnRate := getUserInput("Expected Return Rate: ")
	years := getUserInput("Years: ")

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value with Inflation: %.2f\n", futureRealValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedFRV)
}
