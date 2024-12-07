package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("full.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	world := make(map[[2]int]struct{})
	visited := make(map[[2]int]struct{})

	guardX := -1
	guardY := -1
	guardDir := 0 // up

	var lline string
	y := -1
	for scanner.Scan() {
		y++
		line := scanner.Text()
		lline = line
		for x, character := range line {
			fmt.Println(character)

			if character == 35 {
				world[[2]int{x, y}] = struct{}{}
			} else if character == 94 {
				guardX = x
				guardY = y
			}
		}
	}

	for true {
		orgX := guardX
		orgY := guardY
		if guardDir == 0 {
			// up
			orgY--
		} else if guardDir == 1 {
			// right
			orgX++
		} else if guardDir == 2 {
			// down
			orgY++
		} else {
			// left
			orgX--
		}
		fmt.Println(guardX, guardY)

		_, ok := world[[2]int{orgX, orgY}]
		if ok {
			guardDir = (guardDir + 1) % 4
		} else {
			guardX = orgX
			guardY = orgY
		}

		visited[[2]int{guardX, guardY}] = struct{}{}

		if guardX < 0 || guardY < 0 || guardX >= len(lline) || guardY > y {
			fmt.Println("exiting", len(visited))
			return
		}
	}
}
