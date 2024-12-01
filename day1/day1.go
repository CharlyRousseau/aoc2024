package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CalculateSimilarityScore() {
	left, right, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	total := calculateOccurences(left, right)

	fmt.Println("Total des différences absolues :", total)
}

func CalculateDistanceBetweenLists() {
	left, right, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	sort.Ints(left)
	sort.Ints(right)

	total := calculateDifference(left, right)

	fmt.Println("Total des différences absolues :", total)
}

func readInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line) // Séparer la ligne par les espaces

		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("Ligne mal formée : %s", line)
		}

		leftVal, err1 := strconv.Atoi(parts[0])
		rightVal, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("Erreur de conversion des valeurs : %s", line)
		}

		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}

func calculateDifference(left, right []int) int {
	total := 0
	for i := 0; i < len(left); i++ {
		total += int(math.Abs(float64(left[i] - right[i])))
	}
	return total
}

func calculateOccurences(left, right []int) int {
	total := 0
	count := 0

	for i := 0; i < len(left); i++ {
		count = 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				count += 1
			}
		}
		total += int(left[i]) * count
	}
	return total
}
