package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	file, _ := os.Open("partial.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	antenas := make(map[rune][][2]int)
	var line string
	y := -1
	for scanner.Scan() {
		y++
		line = scanner.Text()
		for x, char := range line {
			if unicode.IsLetter(char) || unicode.IsDigit(char) {
				antenas[char] = append(antenas[char], [2]int{x, y})
			}
		}
	}
	fmt.Println(antenas)
	x := len(line) - 1

	antinodes := make(map[[2]int]struct{})

	for _, antena := range antenas {
		for _, a := range antena {
			for _, b := range antena {
				if a == b {
					continue
				}

				xPos := 2*b[0] - a[0]
				yPos := 2*b[1] - a[1]

				if xPos < 0 || yPos < 0 || xPos > x || yPos > y {
					continue
				}
				antinodes[[2]int{xPos, yPos}] = struct{}{}
				fmt.Println("Found at", xPos, yPos)
			}
		}
	}
	fmt.Println(len(antinodes))
}
