package main

import (
	"fmt"
	"os"
	"time"
)

const inputFile = "test_a.txt"

func main() {
	partA()
}

func partA() {
	start := time.Now()
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// day 2 part A

	fmt.Println("day2 part A answer is:", nil)
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

	// day 2 part B

	fmt.Println("day2 part B answer is:", nil)
	duration := time.Since(start)
	fmt.Println(duration)
}
