package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countElementFrequency(arr []int) map[int]int {
	freqMap := make(map[int]int)
	for _, v := range arr {
		freqMap[v]++
	}
	return freqMap
}

func main() {
	dat, _ := os.Open("partial.txt")
	scanner := bufio.NewScanner(dat)

	var left []int
	var right []int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		l, _ := strconv.Atoi(fields[0])
		r, _ := strconv.Atoi(fields[1])
		left = append(left, l)
		right = append(right, r)
	}

	freq := countElementFrequency(right)

	var score int = 0

	for _, v := range left {
		amount := freq[v]
		score += amount * v
	}
	fmt.Println(score)
}
