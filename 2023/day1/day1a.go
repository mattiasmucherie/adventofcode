package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const inputFile = "input.txt"

func main() {
	day1a()
	day1b()
}

func stringToDigit(l string) string {
	var r = strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	return r.Replace(l)
}

func stringToDigitReverse(l string) string {
	var r = strings.NewReplacer(
		"eno", "1",
		"owt", "2",
		"eerht", "3",
		"ruof", "4",
		"evif", "5",
		"xis", "6",
		"neves", "7",
		"thgie", "8",
		"enin", "9",
	)
	return r.Replace(reverse(l))
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func day1a() {
	// A
	start := time.Now()
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var builder strings.Builder
		for _, character := range line {
			if unicode.IsDigit(character) {
				builder.WriteRune(character)
			}
		}
		numberString := builder.String()
		var last byte
		first := numberString[0]
		last = numberString[len(numberString)-1]
		newString := string(first) + string(last)
		newNumber, _ := strconv.Atoi(newString)

		sum += newNumber
	}

	fmt.Println("day1a answer is:", sum)
	duration := time.Since(start)
	fmt.Println(duration)

}

func day1b() {
	// B

	start := time.Now()
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var builder strings.Builder

		newLine := stringToDigit(line)
		var first byte
		var last byte

		for _, character := range newLine {
			if unicode.IsDigit(character) {
				builder.WriteRune(character)
				first = builder.String()[0]
				break
			}
		}
		var revBuilder strings.Builder
		revNewLine := stringToDigitReverse(line)
		for _, character := range revNewLine {
			if unicode.IsDigit(character) {
				revBuilder.WriteRune(character)
				last = revBuilder.String()[0]
				break
			}
		}

		newString := string(first) + string(last)

		newNumber, _ := strconv.Atoi(newString)

		sum += newNumber
	}

	fmt.Println("day1b answer is:", sum)
	duration := time.Since(start)
	fmt.Println(duration)
}
