package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

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

	slices.Sort(left)
	slices.Sort(right)

	var sum int = 0

	for i := 0; i < len(left); i++ {
		a := left[i]
		b := right[i]

		diff := int(math.Abs(float64(a - b)))
		sum += diff
	}

	fmt.Println(sum)
}
