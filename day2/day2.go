package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var intParts [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)

		var intSlice []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Error converting '%s' to int: %v\n", part, err)
				continue
			}
			intSlice = append(intSlice, num)
		}
		intParts = append(intParts, intSlice)
	}

	return intParts, nil

}

func CountSafeReports() (int, error) {
	count := 0
	reports, err := readInput("input.txt")
	if err != nil {
		return 0, err
	}

	for _, report := range reports {
		if decreasingCase(report) || increasingCase(report) {
			count++
			continue
		}

	}

	fmt.Println(count)
	return count, nil
}

func decreasingCase(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if math.Abs(float64(report[i]-report[i+1])) > 3 || math.Abs(float64(report[i]-report[i+1])) < 1 || report[i]-report[i+1] > 0 {
			return false
		}
	}
	return true
}

func increasingCase(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if report[i]-report[i+1] > 3 || report[i]-report[i+1] < 1 {
			return false
		}
	}
	return true
}

func CountSafeReportsProblemDampener() (int, error) {
	count := 0

	// Read input reports
	reports, err := readInput("input.txt")
	if err != nil {
		return 0, err
	}

	for _, report := range reports {
		if decreasingCase(report) || increasingCase(report) {
			count++
			continue
		}

		for i := 0; i < len(report); i++ {
			//modifiedReport := append(report[:i], report[i+1:]...)
			modifiedReport := make([]int, len(report))
			copy(modifiedReport, report)

			_ = slices.Delete(modifiedReport, i, i+1)
			modifiedReport = modifiedReport[:len(modifiedReport)-1]
			if decreasingCase(modifiedReport) || increasingCase(modifiedReport) {
				count++
				break
			}
		}
	}

	fmt.Println("Safe reports count:", count)
	return count, nil
}
