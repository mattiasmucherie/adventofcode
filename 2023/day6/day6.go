package main

import (
	"bufio"
	"fmt"
	"os"
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

	scanner := bufio.NewScanner(file)
	raceTimes := make([]int, 0)
	raceRecords := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		splitOutput := strings.Split(line, ":")
		if splitOutput[0] == "Time" {
			numbers := strings.Trim(splitOutput[1], " ")
			races := strings.Split(numbers, " ")
			for _, race := range races {
				if race != " " && race != "" {
					num, _ := strconv.Atoi(race)
					raceTimes = append(raceTimes, num)
				}
			}
		}
		if splitOutput[0] == "Distance" {
			numbers := strings.Trim(splitOutput[1], " ")
			races := strings.Split(numbers, " ")
			for _, race := range races {
				if race != " " && race != "" {
					num, _ := strconv.Atoi(race)
					raceRecords = append(raceRecords, num)
				}
			}
		}
	}
	total := 1
	for i, race := range raceTimes {
		waysToWin := 0
		for j := 1; j <= race; j++ {
			speed := j
			timeLeft := race - j
			distanceTraveled := speed * timeLeft
			if distanceTraveled > raceRecords[i] {
				waysToWin++
			}
		}
		total *= waysToWin
	}

	fmt.Println("day6 part A answer is:", total)
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
	raceTimes := make([]string, 0)
	raceRecords := make([]string, 0)
	var raceTime int
	var raceRecord int
	for scanner.Scan() {
		line := scanner.Text()
		splitOutput := strings.Split(line, ":")
		if splitOutput[0] == "Time" {
			numbers := strings.Trim(splitOutput[1], " ")
			races := strings.Split(numbers, " ")
			for _, race := range races {
				if race != " " && race != "" {
					raceTimes = append(raceTimes, race)
				}
			}
		}
		if splitOutput[0] == "Distance" {
			numbers := strings.Trim(splitOutput[1], " ")
			races := strings.Split(numbers, " ")
			for _, race := range races {
				if race != " " && race != "" {
					raceRecords = append(raceRecords, race)

				}
			}
		}
	}
	test := strings.Join(raceTimes[:], "")
	raceTime, _ = strconv.Atoi(test)
	test1 := strings.Join(raceRecords[:], "")
	raceRecord, _ = strconv.Atoi(test1)

	waysToWin := 0
	for j := 1; j <= raceTime; j++ {
		speed := j
		timeLeft := raceTime - j
		distanceTraveled := speed * timeLeft
		if distanceTraveled > raceRecord {
			waysToWin++
		}
	}

	fmt.Println("day6 part B answer is:", waysToWin)
	duration := time.Since(start)
	fmt.Println(duration)
}
