package main

import "fmt"

func main() {

	input, _ := ReadInput()
	result, err := CalculateValidMuls(input)
	if err != nil {
		fmt.Println("Error calculating multiplications:", err)
		return
	}
	fmt.Println("Total of valid mul results:", result)

	input_without_dont := RemoveDontSegments(input)
	resultDont, err := CalculateValidMuls(input_without_dont)
	if err != nil {
		fmt.Println("Error calculating multiplications:", err)
		return
	}
	fmt.Println("Total of valid mul results:", resultDont)
}
