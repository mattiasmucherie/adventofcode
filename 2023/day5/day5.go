package main

import (
	"bufio"
	"fmt"
	"os"
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

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

	}

	fmt.Println("day5 part A answer is:", nil)
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

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		fmt.Println(line)

	}

	fmt.Println("day5 part B answer is:", nil)
	duration := time.Since(start)
	fmt.Println(duration)
}
