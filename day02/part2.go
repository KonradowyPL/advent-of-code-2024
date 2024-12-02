package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("full.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	safe := 0
loop:
	for scanner.Scan() {
		line := scanner.Text()
		raports := strings.Fields(line)
		valid := check(raports)
		if valid {
			safe++
			continue
		}

		for skipIndex := 0; skipIndex < len(raports); skipIndex++ {
			newArr := make([]string, 0, len(raports)-1)

			for i, v := range raports {
				if i == skipIndex {
					continue // Skip the nth element
				}
				newArr = append(newArr, v)

			}

			if check(newArr) {
				safe++
				continue loop
			}

		}
	}
	fmt.Println(safe)
}
func check(raports []string) bool {
	first, _ := strconv.Atoi(raports[0])
	second, _ := strconv.Atoi(raports[1])

	diff := first > second
	current, _ := strconv.Atoi(raports[0])
	for i, v := range raports {
		previous := current
		current, _ = strconv.Atoi(v)
		if i == 0 {
			continue
		}
		if previous == current {
			return false
		}

		if (previous > current) != diff {
			return false
		}
		change := math.Abs(float64(previous - current))
		if change > 3 {
			return false
		}
	}
	return true
}
