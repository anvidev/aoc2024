package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("../data.txt")
	if err != nil {
		panic(err)
	}

	left, right := splitAndSortData(b)
	distance := calculateDistance(left, right)
	similarityScore := calculateSimularityScore(left, right)

	fmt.Printf("total distance: %d - similarity score: %d", distance, similarityScore)
}

func splitAndSortData(b []byte) ([]int, []int) {
	numbers := strings.Split(strings.TrimSpace(string(b)), "\n")
	var left, right []int

	for _, line := range numbers {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			panic(fmt.Errorf("invalid line format"))
		}

		leftInt, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		rightInt, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	return left, right
}

func calculateDistance(left, right []int) int {
	if len(left) != len(right) {
		panic(fmt.Errorf("left and right is unequal in length"))
	}

	sort.Ints(left)
	sort.Ints(right)

	total := 0

	for i := range left {
		total += int(math.Abs(float64(left[i]) - float64(right[i])))
	}

	return total
}

func calculateSimularityScore(left, right []int) int {
	freq := make(map[int]int)

	for i := range right {
		freq[right[i]]++
	}

	total := 0
	for _, num := range left {
		total += num * freq[num]
	}

	return total
}
