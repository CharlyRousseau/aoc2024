package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RemoveDontSegments(s string) string {
	var filteredDos []string
	dos := strings.Split(s, "do()")
	for _, do := range dos {
		part, _, _ := strings.Cut(do, "don't()")
		filteredDos = append(filteredDos, part)
	}
	return strings.Join(filteredDos, "")
}

func ReadInput() (string, error) {

	input, err := os.ReadFile("input.txt")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(input)), nil
}

func CalculateValidMuls(input string) (int, error) {

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Find all matches
	matches := re.FindAllStringSubmatch(input, -1)

	fmt.Println(matches)
	total := 0
	for _, match := range matches {
		// Extract and convert the numbers
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		// Calculate their product and add to the total
		total += x * y
	}

	return total, nil
}
