package main

import (
	"bufio"
	"fmt"
	"os"
)

func get(world [][]rune, x int, y int) rune {
	if y < 0 || x < 0 {
		return 0
	}

	if y >= len(world) {
		return 0
	}
	line := world[y]

	if x >= len(line) {
		return 0
	}

	return line[x]
}

func tryAddToBuffer(candidate [2]int, visited map[[2]int]struct{}, buffer *[][2]int) {
	if _, inVisited := visited[candidate]; !inVisited {
		addToBuffer := true
		for _, b := range *buffer {
			if b == candidate {
				addToBuffer = false
				break
			}
		}
		if addToBuffer {
			*buffer = append(*buffer, candidate)
		}
	}
}

func main() {
	file, _ := os.Open("full.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var world [][]rune

	height := -1
	for scanner.Scan() {
		height++
		line := scanner.Text()
		var worldLine []rune
		for _, character := range line {
			worldLine = append(worldLine, character)
		}
		world = append(world, worldLine)
	}

	sum := 0
	visited := make(map[[2]int]struct{})

	for y, line := range world {
		for x, char := range line {
			_, ok := visited[[2]int{x, y}]
			if ok {
				continue
			}

			var buffer = [][2]int{{x, y}}
			border := 0
			skipped := 0
			i := 0
			for i = 0; i < len(buffer); i++ {
				item := buffer[i]

				_, ok := visited[item]
				if ok {
					skipped++
					continue
				}

				visited[item] = struct{}{}

				if get(world, item[0]+1, item[1]) == char { // South
					candidate := [2]int{item[0] + 1, item[1]}
					tryAddToBuffer(candidate, visited, &buffer)
				} else if !(get(world, item[0]+1, item[1]+1) != char && get(world, item[0], item[1]+1) == char) {
					border++
				}

				if get(world, item[0]-1, item[1]) == char { // North
					candidate := [2]int{item[0] - 1, item[1]}
					tryAddToBuffer(candidate, visited, &buffer)
				} else if !(get(world, item[0]-1, item[1]-1) != char && get(world, item[0], item[1]-1) == char) {

					border++
				}

				if get(world, item[0], item[1]+1) == char { // East
					candidate := [2]int{item[0], item[1] + 1}
					tryAddToBuffer(candidate, visited, &buffer)
				} else if !(get(world, item[0]+1, item[1]+1) != char && get(world, item[0]+1, item[1]) == char) {

					border++
				}

				if get(world, item[0], item[1]-1) == char { // West
					candidate := [2]int{item[0], item[1] - 1}
					tryAddToBuffer(candidate, visited, &buffer)
				} else if !(get(world, item[0]-1, item[1]-1) != char && get(world, item[0]-1, item[1]) == char) {

					border++
				}

			}
			fmt.Println(i-skipped, border, string(char))
			sum += (i - skipped) * border
		}

	}
	fmt.Println(sum)
}
