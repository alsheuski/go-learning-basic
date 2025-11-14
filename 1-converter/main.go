package main

import (
	"fmt"
	"os"
)

func main() {
	const usdToEur = 0.86
	const usdToRub = 81
	const eurToRub = 1 / usdToEur * usdToRub

	userInput := getUserInput()
}

func getUserInput() string {
	var userInput string

	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Print("Ошибка ввода данных: ", err)
		os.Exit(1)
	}

	return userInput
}

func convert(amount float64, from, to string) float64 {
	return amount
}
