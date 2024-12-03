package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	b, err := os.ReadFile("../data.txt")
	if err != nil {
		panic(err)
	}

	rgx, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	instructions := rgx.FindAllString(string(b), -1)

	fmt.Println(instructions)
}
