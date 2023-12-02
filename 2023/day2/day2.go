package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	max := [3]int{14, 12, 13}
	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		gamePossible := true
		line := scanner.Text()
		parts := strings.Split(line, ":")
		pattern := `\d+`
		re := regexp.MustCompile(pattern)
		gameId, _ := strconv.Atoi(re.FindString(parts[0]))
		colorPattern := [3]string{`(\d+) blue`, `(\d+) red`, `(\d+) green`}

		for i, p := range colorPattern {
			colorRe := regexp.MustCompile(p)
			colorMatch := colorRe.FindAllStringSubmatch(line, -1)

			for _, c := range colorMatch {
				colorCount, _ := strconv.Atoi(c[1])

				if colorCount > max[i] {
					gamePossible = false
				}
			}
		}

		if gamePossible {
			sum += gameId
		}
	}

	fmt.Println("day2 part A answer is:", sum)
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

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		colorPattern := [3]string{`(\d+) blue`, `(\d+) red`, `(\d+) green`}
		mininumColor := [3]int{0, 0, 0}

		for i, p := range colorPattern {
			colorRe := regexp.MustCompile(p)
			colorMatch := colorRe.FindAllStringSubmatch(line, -1)

			for _, c := range colorMatch {
				colorCount, _ := strconv.Atoi(c[1])

				if colorCount > mininumColor[i] {
					mininumColor[i] = colorCount
				}
			}
		}

		product := 1
		for _, value := range mininumColor {
			product *= value
		}
		sum += product

	}

	fmt.Println("day2 part B answer is:", sum)
	duration := time.Since(start)
	fmt.Println(duration)
}
