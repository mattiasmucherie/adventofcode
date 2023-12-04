package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

const inputFile = "input.txt"

func main() {
	partA()
	partB()
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
	totalSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ":")[1]
		placeholder := strings.Split(numbers, "|")
		winning, own := placeholder[0], placeholder[1]
		winningArray := strings.Split(strings.Trim(winning, " "), " ")
		ownArray := strings.Split(strings.Trim(own, " "), " ")

		sum := 0
		for _, num := range ownArray {

			for _, win := range winningArray {
				if win == " " || win == "" || num == " " || num == "" {
					continue
				} else if num == win {
					sum++
				}
			}
		}

		if sum > 0 {
			totalSum += int(math.Pow(2, float64(sum-1)))
		}

	}

	fmt.Println("day 4 part A answer is:", totalSum)
	duration := time.Since(start)
	fmt.Println(duration)
}

func partB() {
	start := time.Now()
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	lineScanner := bufio.NewScanner(file)
	totalSum := 0
	lineCount := 0
	for lineScanner.Scan() {
		lineCount++
	}
	// Reset the FILE POINTER!!!!!
	file.Seek(0, 0)

	card := make([]int, lineCount)
	for i := range card {
		card[i] = 1
	}

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {

		line := scanner.Text()
		numbers := strings.Split(line, ":")[1]
		placeholder := strings.Split(numbers, "|")
		winning, own := placeholder[0], placeholder[1]
		winningArray := strings.Split(strings.Trim(winning, " "), " ")
		ownArray := strings.Split(strings.Trim(own, " "), " ")

		sum := 0
		for _, num := range ownArray {
			for _, win := range winningArray {
				if win == " " || win == "" || num == " " || num == "" {
					continue
				} else if num == win {
					sum++
				}
			}
		}

		if sum > 0 {
			for j := 1; j <= sum; j++ {
				card[j+i] += card[i]
			}
		}

	}
	for _, c := range card {
		totalSum += c
	}
	fmt.Println("day 4 part B answer is:", totalSum)
	duration := time.Since(start)
	fmt.Println(duration)
}
