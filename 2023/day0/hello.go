package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    // Open the input file
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Initialize variables
    var maxCalories int
    currentCalories := 0

    // Create a scanner to read the file
    scanner := bufio.NewScanner(file)

    // Read file line by line
    for scanner.Scan() {
        line := scanner.Text()
        // Check if line is empty, indicating a new elf's data
        if line == "" {
            if currentCalories > maxCalories {
                maxCalories = currentCalories
            }
            currentCalories = 0
        } else {
            // Convert line to integer and add to current calories
            calories, err := strconv.Atoi(line)
            if err != nil {
                fmt.Println("Error converting string to int:", err)
                return
            }
            currentCalories += calories
        }
    }

    // Check last elf's calories
    if currentCalories > maxCalories {
        maxCalories = currentCalories
    }

    // Check for errors during Scan
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }

    // Output the result
    fmt.Printf("The Elf with the most calories is carrying %d calories.\n", maxCalories)
}

