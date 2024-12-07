package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func f(char rune) bool {
	return !unicode.IsDigit(char)
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
		for i := 0; i < 1<<len(numbers); i++ {
			accumulator := numbers[1]

			for index := 2; index < len(numbers); index++ {
				if (i>>index)&1 == 1 {
					accumulator += numbers[index]
				} else {
					accumulator *= numbers[index]
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
