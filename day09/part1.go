package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	line, _ := os.ReadFile("partial.txt")

	var disk []int
	id := 0
	for index, char := range line {
		size, _ := strconv.Atoi(string(char))
		for i := 0; i < size; i++ {
			if index%2 == 0 {
				disk = append(disk, id)
			} else {
				disk = append(disk, -1)
			}

		}
		if index%2 == 0 {
			id++
		}
	}
	fmt.Println(disk)

	lower := 0
	higher := len(disk) - 1
	for lower < higher {
		if disk[lower] != -1 {
			lower++
			continue
		}
		if disk[higher] == -1 {
			higher--
			continue
		}
		disk[lower], disk[higher] = disk[higher], disk[lower]
	}
	checksum := 0
	for i, num := range disk {
		if num == -1 {
			continue
		}
		checksum += i * num
	}
	fmt.Println(checksum)
}
