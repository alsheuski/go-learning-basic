package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var funcList = map[string]func([]float64) float64{
	"AVG": calcAvg,
	"SUM": calcSum,
	"MED": calcMedian,
}

func main() {
	operation := getUserOperation()
	numbers := getUserDigits()
	result := calculate(operation, numbers)

	fmt.Print(result)
}

func calculate(operation string, numbers []float64) float64 {
	fn, ok := funcList[operation]
	if !ok {
		fmt.Println("Unknown openration: ", operation)
		os.Exit(1)
	}

	return fn(numbers)
}

func calcAvg(data []float64) float64 {
	sum := calcSum(data)
	return sum / float64(len(data))
}

func calcMedian(data []float64) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}
	dataCopy := make([]float64, n)
	copy(dataCopy, data)
	sort.Float64s(dataCopy)

	mid := n / 2
	if n%2 == 0 {
		return (dataCopy[mid-1] + dataCopy[mid]) / 2
	} else {
		return dataCopy[mid]
	}
}

func calcSum(numbers []float64) float64 {
	var sum float64
	for _, val := range numbers {
		sum += val
	}

	return sum
}

func getUserDigits() []float64 {
	fmt.Print("Введите числа через запятую: ")

	for {
		input, err := userInput()
		if err != nil {
			continue
		}

		if len(input) == 0 {
			continue
		}

		numbers, err := parseDigits(input)
		if err != nil {
			continue
		}

		return numbers
	}
}

func parseDigits(str string) ([]float64, error) {
	members := strings.Split(str, ",")
	digits := []float64{}

	for _, val := range members {
		member, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
		if err != nil {
			return nil, err
		}

		digits = append(digits, member)
	}

	return digits, nil
}

func getUserOperation() string {
	fmt.Print("Введите операцию (AVG - среднее, SUM - сумма, MED - медиана): ")

	for {
		input, err := userInput()
		if err != nil {
			continue
		}

		if input != "AVG" && input != "SUM" && input != "MED" {
			continue
		}

		return input
	}
}

func userInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
		return "", err
	}

	return strings.TrimRight(input, "\n"), nil
}
