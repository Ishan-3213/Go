package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue, expenses, tax_rate float64
	var err error
	for {
		revenue, err = TakeInputes("Enter Reveue: ")

		if err != nil {
			fmt.Println("ERROR!!")
			fmt.Println(err)
			continue
		}
		expenses, err = TakeInputes("Enter expenses: ")

		if err != nil {
			fmt.Println("ERROR!!")
			fmt.Println(err)
			continue
		}

		tax_rate, err = TakeInputes("Enter Tax Rate: ")

		if err != nil {
			fmt.Println("ERROR!!")
			fmt.Println(err)
			continue
		}

		if err == nil {
			break
		}
	}

	earningBeforeTax, earningAfterTax, ratio := CalculatOutput(revenue, expenses, tax_rate)
	fmt.Printf("Earning Before: %v \nEarning After Tax: %v \nRatio: %0.2f", earningBeforeTax, earningAfterTax, ratio)

}

func CalculatOutput(revenue, expenses, tax_rate float64) (earningBeforeTax, earningAfterTax, ratio float64) {
	earningBeforeTax = revenue - expenses
	earningAfterTax = earningBeforeTax * (1 - tax_rate/100)
	ratio = earningBeforeTax / earningAfterTax

	data := fmt.Sprint("EBT: ", earningBeforeTax, "\nProfit: ", earningAfterTax, "\nRatio: ", ratio)

	err := os.WriteFile("temp_1.txt", []byte(data), 0644)

	if err != nil {
		fmt.Println(err, "ERROR !!! ")
	}
	return earningBeforeTax, earningAfterTax, ratio

}

func TakeInputes(info string) (op float64, err error) {
	fmt.Print(info)
	fmt.Scan(&op)

	if op <= 0 {
		return 0, errors.New("Invalid input is provided! Please enter non-negative/non-zero value")
	}

	return op, nil
}
