package main

import (
	"bufio"
	"fmt"
	"os"
)

func get(x int, y int, grid [][]byte) byte {
	if y >= len(grid) || y < 0 || x < 0 {
		return 0
	}
	row := grid[y]
	if x >= len(row) {
		return 0
	}
	return row[x]
}

func isAt(x int, y int, grid [][]byte) int {
	count := 0
	var word []byte = []byte("XMAS")
	if get(x, y, grid) != word[0] {
	}

	for vx := -1; vx <= 1; vx++ {
		for vy := -1; vy <= 1; vy++ {
			if vx == 0 && vy == 0 {
				continue
			}

			for i := 0; i < len(word); i++ {
				nx, ny := x+i*vx, y+i*vy
				if get(nx, ny, grid) != word[i] {
					break
				}
				if i == len(word)-1 {
					count++
				}
			}

		}
	}

	return count
}

func main() {
	file, _ := os.Open("full.txt")
	defer file.Close()
	var grid [][]byte
	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]byte, len(line))
		copy(row, line)
		grid = append(grid, row)
	}

	for y, row := range grid {
		for x, _ := range row {
			sum += isAt(x, y, grid)
		}
	}
	fmt.Println(sum)
}
