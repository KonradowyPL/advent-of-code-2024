package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func f(char rune) bool {
	return !unicode.IsDigit(char)
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func combineInts(a, b int) int {
	// Find the number of digits in b
	multiplier := 1
	for temp := b; temp > 0; temp /= 10 {
		multiplier *= 10
	}

	// Combine the numbers
	return a*multiplier + b
}

func main() {
	file, _ := os.Open("full.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.FieldsFunc(line, f)
		var numbers []int

		for _, number := range fields {
			num, _ := strconv.Atoi(number)
			numbers = append(numbers, num)
		}

		target := numbers[0]
		for i := 0; i < powInt(3, len(numbers)); i++ {
			base3Representation := fmt.Sprintf("%0*s", len(numbers), strconv.FormatInt(int64(i), 3))
			accumulator := numbers[1]

			for index := 2; index < len(numbers); index++ {
				if base3Representation[index] == 48 {
					accumulator += numbers[index]
				} else if base3Representation[index] == 49 {
					accumulator *= numbers[index]
				} else if base3Representation[index] == 50 {
					accumulator = combineInts(accumulator, numbers[index])
				}
			}

			if accumulator == target {
				fmt.Println(accumulator, target)
				sum += target
				break
			}
		}
	}
	fmt.Println(sum)
}
