package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("../data.txt")
	if err != nil {
		panic(err)
	}

	totalPartOne, err := parseAndCalculatePartOne(b)
	if err != nil {
		panic(err)
	}

	totalPartTwo, err := parseAndCalculatePartTwo(b)
	if err != nil {
		panic(err)
	}

	fmt.Println(totalPartOne, totalPartTwo)
}

func parseAndCalculatePartOne(b []byte) (int, error) {
	rgx := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	instructions := rgx.FindAllString(string(b), -1)

	rgx = regexp.MustCompile(`\d{1,3},\d{1,3}`)

	total := 0

	for _, instruction := range instructions {
		ds := strings.Split(rgx.FindString(instruction), ",")
		if len(ds) != 2 {
			return 0, fmt.Errorf("ballsack")
		}

		int1, err := strconv.Atoi(ds[0])
		if err != nil {
			return 0, err
		}

		int2, err := strconv.Atoi(ds[1])
		if err != nil {
			return 0, err
		}

		total += int1 * int2
	}

	return total, nil
}

func parseAndCalculatePartTwo(b []byte) (int, error) {
	rgx := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`)

	instructions := rgx.FindAllString(string(b), -1)

	rgx = regexp.MustCompile(`\d{1,3},\d{1,3}`)

	total := 0
	flag := true

	for _, instruction := range instructions {
		if instruction == "do()" {
			flag = true
		} else if instruction == "don't()" {
			flag = false
		} else {
			if flag {

				ds := strings.Split(rgx.FindString(instruction), ",")
				if len(ds) != 2 {
					return 0, fmt.Errorf("ballsack")
				}

				int1, err := strconv.Atoi(ds[0])
				if err != nil {
					return 0, err
				}

				int2, err := strconv.Atoi(ds[1])
				if err != nil {
					return 0, err
				}

				total += int1 * int2
			}
		}
	}

	return total, nil
}
