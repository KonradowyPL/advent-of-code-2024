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
	pattern := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	re := regexp.MustCompile(`\d+`)
	matches := pattern.FindAllString(text, -1)
	for _, v := range matches {
		numbers := re.FindAllString(v, -1)
		a, _ := strconv.Atoi(numbers[0])
		b, _ := strconv.Atoi(numbers[1])

		sum += a * b
	}
	fmt.Println(sum)
}
