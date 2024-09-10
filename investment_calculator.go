package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount, years float64 = 1000, 10

	expectedReturnRate := 5.5
	fmt.Print("Enter the investment Amount: ")
	fmt.Scanln(&investmentAmount)
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	newFutureValue, newFutureRealValue := StringFormate(futureValue, futureRealValue)
	fmt.Println(newFutureValue, "\n", newFutureRealValue)
}

func StringFormate(futureValue, futureRealValue float64) (fv, frv string) {
	fv = fmt.Sprintf("Future Value:  %.2f", futureValue)
	frv = fmt.Sprintf("Futre real value: %.1f", futureRealValue)
	return fv, frv
	// return
}
