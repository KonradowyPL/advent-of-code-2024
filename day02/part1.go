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

	for scanner.Scan() {
		line := scanner.Text()
		raports := strings.Fields(line)
		first, _ := strconv.Atoi(raports[0])
		second, _ := strconv.Atoi(raports[1])

		diff := first > second
		current, _ := strconv.Atoi(raports[0])
		isSafe := true
		for i, v := range raports {
			previous := current
			current, _ = strconv.Atoi(v)
			if i == 0 {
				continue
			}
			if previous == current {
				isSafe = false
			}

			if (previous > current) != diff {
				isSafe = false
			}
			change := math.Abs(float64(previous - current))
			if change > 3 {
				isSafe = false
			}
		}
		if isSafe {
			safe++
		} else {
			fmt.Println("failed:", line)

		}
	}
	fmt.Println(safe)
}
