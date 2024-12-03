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

	rgx, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	instructions := rgx.FindAllString(string(b), -1)

	rgx, err = regexp.Compile(`\d{1,3},\d{1,3}`)
	if err != nil {
		panic(err)
	}

	total := 0

	for _, instruction := range instructions {
		ds := strings.Split(rgx.FindString(instruction), ",")
		if len(ds) != 2 {
			panic("ballsack")
		}

		int1, err := strconv.Atoi(ds[0])
		if err != nil {
			panic(err)
		}

		int2, err := strconv.Atoi(ds[1])
		if err != nil {
			panic(err)
		}

		total += int1 * int2

	}

	fmt.Println(total)
}
