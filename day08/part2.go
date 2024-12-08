package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// Function to calculate the GCD of two numbers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	file, _ := os.Open("full.txt")
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
		for i1, a := range antena {
			for i2, b := range antena {
				if i1 == i2 {
					continue
				}
				dx := a[0] - b[0]
				dy := a[1] - b[1]

				if dx != 0 && dy != 0 {

					l := gcd(dx, dy)
					dx = dx / l
					dy = dy / l
				}

				xPos := b[0]
				yPos := b[1]

				for !(xPos < 0 || yPos < 0 || xPos > x || yPos > y) {
					antinodes[[2]int{xPos, yPos}] = struct{}{}
					xPos += dx
					yPos += dy
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}
