package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	sum := 0

	file, _ := os.ReadFile("full.txt")
	text := string(file)
	pattern := regexp.MustCompile("(mul\\((\\d+),(\\d+)\\))|do\\(\\)|don't\\(\\)")
	re := regexp.MustCompile(`\d+`)
	matches := pattern.FindAllString(text, -1)
	enabled := true
	for _, v := range matches {

		if v == "do()" {
			enabled = true
			continue
		} else if v == "don't()" {
			enabled = false
			continue
		}
		if !enabled {
			continue
		}

		numbers := re.FindAllString(v, -1)
		a, _ := strconv.Atoi(numbers[0])
		b, _ := strconv.Atoi(numbers[1])

		sum += a * b
	}
	fmt.Println(sum)
}
