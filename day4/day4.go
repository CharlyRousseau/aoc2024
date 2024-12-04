package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var directionsBis = [][2]int{
	{1, 1},   // top-left to bottom-right diagonal
	{-1, -1}, // bottom-left to top-right diagonal
	{1, -1},  // top-right to bottom-left diagonal
	{-1, 1},  // bottom-right to top-left diagonal
}

var directions = [][2]int{
	{0, 1},   // right
	{0, -1},  // left
	{1, 0},   // down
	{-1, 0},  // up
	{1, 1},   // down-right diagonal
	{-1, -1}, // up-left diagonal
	{1, -1},  // down-left diagonal
	{-1, 1},  // up-right diagonal
}

func checkDirection(grid [][]rune, word string, i, j, di, dj int) bool {
	for k := 0; k < len(word); k++ {

		x := i + k*di
		y := j + k*dj

		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			return false
		}
		if grid[x][y] != rune(word[k]) {
			return false
		}
	}
	return true
}

func countXMAS(grid [][]rune) int {
	word := "XMAS"
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for _, dir := range directions {
				if checkDirection(grid, word, i, j, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	return count
}

func checkMAS(grid [][]rune, word string, i, j, di, dj int) bool {
	for k := 0; k < len(word); k++ {
		x := i + k*di
		y := j + k*dj
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			return false
		}
		if grid[x][y] != rune(word[k]) {
			return false
		}
	}
	return true
}

func checkXMASPattern(grid [][]rune, i, j int) bool {
	// The center of the X must be 'A'
	if grid[i][j] != 'A' {
		return false
	}

	for _, dir := range directionsBis {
		di, dj := dir[0], dir[1]
		if !(checkMAS(grid, "MAS", i-di, j-dj, di, dj) || checkMAS(grid, "SAM", i-di, j-dj, di, dj)) {
			return false
		}
	}

	return true
}

func countXMASPattern(grid [][]rune) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if checkXMASPattern(grid, i, j) {
				count++
			}
		}
	}

	return count
}

func readGridFromFile(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line)) // Convert the line into a slice of runes
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func main() {
	grid, err := readGridFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	resultPart1 := countXMAS(grid)
	fmt.Println("Part 1: The word XMAS appears", resultPart1, "times.")
	resultPart2 := countXMASPattern(grid)
	fmt.Println("Part 2: The X-MAS pattern appears", resultPart2, "times.")
}
