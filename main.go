package main

import (
	"fmt"
)

func main() {
	// #01 - Check Palindrom
	fmt.Println("#01 - Check Palindrom")
	wordArr := []string{"radar", "hello"}
	for _, str := range wordArr {
		isValid := checkPalindrom(str)
		fmt.Println("Output:", isValid)
		fmt.Println()
	}

	// #02 - Get Maximum Value
	fmt.Println("#02 - Get Maximum Value")
	initArray := []int{3, 5, 1, 9, 2}
	result := getMaxValue(initArray)
	fmt.Println("Output:", result)
	fmt.Println()

	// #03 - Print Triangle Pattern
	fmt.Println("#03 - Print Triangle Pattern")
	rowNumber := 5
	printTriangle(rowNumber)

	// #04 - Calculate Factorial
	fmt.Println("#04 - Calculate Factorial")
	inputNumber := 5
	fmt.Println("Input:", inputNumber)
	result = calcFactorial(inputNumber)
	fmt.Println("Output:", result)
}

// #01 - Check Palindrom
func checkPalindrom(word string) bool {
	fmt.Println("Input:", word)

	strLength := len(word)
	for i := 0; i < strLength/2; i++ {
		if word[i] != word[strLength-1-i] {
			return false
		}
	}
	return true
}

// #02 - Get Maximum Value
func getMaxValue(numArr []int) int {
	fmt.Println("Input:", numArr)

	maxValue := numArr[0]
	for _, el := range numArr {
		if el > maxValue {
			maxValue = el
		}
	}
	return maxValue
}

// #03 - Print Triangle Pattern
func printTriangle(rowNumber int) {
	fmt.Println("Input:", rowNumber)
	fmt.Println("Output:")

	for i := 1; i <= rowNumber; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
}

// #04 - Calculate Factorial
func calcFactorial(inputNumber int) int {
	if inputNumber == 1 || inputNumber == 0 {
		return 1
	} else {
		return inputNumber * calcFactorial(inputNumber-1)
	}
}
