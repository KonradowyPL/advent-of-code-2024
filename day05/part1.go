package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func IndexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1 // Value not found
}

func split(c rune) bool {
	return !unicode.IsNumber(c)
}

func main() {
	file, _ := os.Open("full.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var roules [][2]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		fields := strings.FieldsFunc(line, split)

		key, _ := strconv.Atoi(fields[0])
		value, _ := strconv.Atoi(fields[1])

		dat := [2]int{key, value}
		roules = append(roules, dat)
	}

	sum := 0

manual:
	for scanner.Scan() {
		line := scanner.Text()
		_pages := strings.FieldsFunc(line, split)
		var pages []int

		for _, page := range _pages {
			num, _ := strconv.Atoi(page)

			pages = append(pages, num)
		}

		for _, roule := range roules {
			a := IndexOf(pages, roule[0])
			b := IndexOf(pages, roule[1])
			if a == -1 || b == -1 {
				continue
			}
			if a > b {
				continue manual
			}
		}
		sum += pages[len(pages)/2]
	}
	fmt.Println(sum)
}
