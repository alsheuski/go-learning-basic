package main

import (
	"fmt"
	"strconv"
)

func main() {
	currencyMap := make(map[string]float64, 3)
	currencyMap["USDEUR"] = 0.86
	currencyMap["USDRUB"] = 81
	currencyMap["EURRUB"] = currencyMap["USDRUB"] / currencyMap["USDEUR"]
	currencyMap["EURUSD"] = 1 / currencyMap["USDEUR"]
	currencyMap["RUBEUR"] = 1 / currencyMap["EURRUB"]
	currencyMap["RUBUSD"] = 1 / currencyMap["USDRUB"]

	initialCurrency := getInitialCurrency()
	moneyAmountToConvert := getMoneyAmount()
	targetCurrency := getTargetCurrency(initialCurrency)

	conversionResult := convert(moneyAmountToConvert, initialCurrency, targetCurrency, currencyMap)

	fmt.Printf("Итоговая сумма %v: %.2f\n", targetCurrency, conversionResult)
}

func getTargetCurrency(pickedCurrency string) string {
	var available string

	switch pickedCurrency {
	case "EUR":
		available = "RUB, USD"
	case "RUB":
		available = "EUR, USD"
	default:
		available = "EUR, RUB"
	}

	for {
		text := fmt.Sprintf("Введите валюту для обмена (%v): ", available)
		currency, err := getUserInputString(text)
		if err != nil {
			continue
		}

		if currency == pickedCurrency {
			continue
		}

		if currency != "EUR" && currency != "RUB" && currency != "USD" {
			continue
		}

		return currency
	}
}

func getMoneyAmount() float64 {
	for {
		moneyAmountString, err := getUserInputString("Введите сумму: ")
		if err != nil {
			continue
		}

		moneyAmount, err := strconv.ParseFloat(moneyAmountString, 64)
		if err != nil {
			continue
		}

		if moneyAmount <= 0 {
			continue
		}

		return moneyAmount

	}
}

func getInitialCurrency() string {
	for {
		initialCurrency, err := getUserInputString("Введите валюту для обмена (EUR, RUB, USD): ")
		if err != nil {
			continue
		}

		if initialCurrency != "EUR" && initialCurrency != "RUB" && initialCurrency != "USD" {
			continue
		}

		return initialCurrency
	}
}

func getUserInputString(text string) (string, error) {
	var userInput string

	fmt.Print(text)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Print("Ошибка ввода данных: ", err)
		return "", err
	}

	return userInput, nil
}

func convert(amount float64, from, to string, currencyMap map[string]float64) float64 {
	return amount * currencyMap[from+to]
}
