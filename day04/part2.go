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
	matrices := [][4]int{
		{1, 0, 0, 1},   // Identity / No Change
		{0, -1, 1, 0},  // 90° Counterclockwise
		{-1, 0, 0, -1}, // 180°
		{0, 1, -1, 0},  // 270° Counterclockwise (or 90° Clockwise)
		{1, 0, 0, -1},  // Reflection across x-axis
		{-1, 0, 0, 1},  // Reflection across y-axis
		{0, 1, 1, 0},   // Reflection across y = x
		{0, -1, -1, 0}, // Reflection across y = -x
	}

	for _, ele := range matrices {
		if string(get(x, y, grid)) == "A" &&
			string(get(x-ele[0]-ele[1], y-ele[2]-ele[3], grid)) == "M" &&
			string(get(x-ele[0]+ele[1], y-ele[2]+ele[3], grid)) == "M" &&
			string(get(x+ele[0]+ele[1], y+ele[2]+ele[3], grid)) == "S" &&
			string(get(x+ele[0]-ele[1], y+ele[2]-ele[3], grid)) == "S" {
			fmt.Println("found at ", x, y)
			return 1
		}
	}

	return 0
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
