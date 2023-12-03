package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode"
)

const inputFile = "input.txt"

func main() {
	partA()
	// partB()
}

func partA() {
	start := time.Now()
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		grid = append(grid, row)
	}
	sum := 0
	for i, l := range grid {
		var number string
		isNextTo := false
		for j, c := range l {
			if unicode.IsDigit(c) {
				number = number + string(c)
				if isNextToSpecialCharacter(grid, i, j) {
					isNextTo = true

				}
				if isNextTo && j == len(l)-1 {
					num, _ := strconv.Atoi(number)
					sum += num
					number = ""
					isNextTo = false
				}
			} else if isNextTo {
				num, _ := strconv.Atoi(number)
				sum += num
				number = ""
				isNextTo = false
			} else {
				number = ""
			}
		}

	}

	fmt.Println("day3 part A answer is:", sum)
	duration := time.Since(start)
	fmt.Println(duration)
}

func isNextToSpecialCharacter(grid [][]rune, x int, y int) bool {
	dir := [][]int{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}

	for _, direction := range dir {
		newX, newY := x+direction[0], y+direction[1]
		if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[newX]) {
			if !unicode.IsDigit(grid[newX][newY]) && grid[newX][newY] != '.' {
				return true
			}
		}
	}
	return false
}

func partB() {
	start := time.Now()
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}

	fmt.Println("day3 part A answer is:", nil)
	duration := time.Since(start)
	fmt.Println(duration)
}
