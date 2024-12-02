package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("../data.txt")
	if err != nil {
		panic(err)
	}

	reports := parseReportData(b)

	var safeAmount int
	var safeAmountWithReactor int

	for _, level := range reports {
		if safe := checkLevelPartOne(level); safe {
			safeAmount++
		}
		if safe := checkLevelPartTwo(level); safe {
			safeAmountWithReactor++
		}
	}

	fmt.Println(safeAmount, safeAmountWithReactor)
}

func parseReportData(b []byte) [][]int {
	levels := strings.Split(strings.TrimSpace(string(b)), "\n")
	reports := make([][]int, len(levels))

	for i, line := range levels {
		fields := strings.Fields(line)
		intLevel := make([]int, len(fields))

		for j, field := range fields {
			int, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			intLevel[j] = int
		}

		reports[i] = intLevel
	}

	return reports
}

func checkLevelPartOne(level []int) bool {
	diff := level[1] - level[0]

	if diff == 0 || math.Abs(float64(diff)) > 3 {
		return false
	}

	isIncreasing := diff > 0

	for i := 1; i < len(level)-1; i++ {
		diff := level[i+1] - level[i]

		if diff == 0 || math.Abs(float64(diff)) > 3 {
			return false
		}

		if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			return false
		}
	}

	return true
}

func checkLevelPartTwo(level []int) bool {
	for i := range level {
		newLevel := append([]int{}, level[:i]...)
		newLevel = append(newLevel, level[i+1:]...)

		if checkLevelPartOne(newLevel) {
			return true
		}
	}

	return false
}
