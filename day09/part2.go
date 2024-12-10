// THIS DOES NOT WORK :(
package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Part struct {
	size int
	id   int
}

func mergeParts(disk []Part) []Part {
	if len(disk) == 0 {
		return disk
	}

	var result []Part
	current := disk[0]

	for i := 1; i < len(disk); i++ {
		if disk[i].id == current.id { // Merge if IDs are the same
			current.size += disk[i].size
		} else {
			if current.size > 0 { // Only add if size > 0
				result = append(result, current)
			}
			current = disk[i] // Move to the next part
		}
	}

	// Add the last part if its size is > 0
	if current.size > 0 {
		result = append(result, current)
	}

	result = append(result, Part{0, -1})

	return result
}

func display(disk []Part) {
	for _, part := range disk {
		if part.id == -1 {
			fmt.Print(strings.Repeat(".", part.size))
		} else {
			fmt.Print(strings.Repeat(strconv.Itoa(part.id), part.size))
		}
	}
}

func main() {
	fmt.Println("PART 2 DOES NOT WORK!!!!")
	line, _ := os.ReadFile("full.txt")

	var disk []Part
	id := 0
	for index, char := range line {
		size, _ := strconv.Atoi(string(char))
		if index%2 == 0 {
			disk = append(disk, Part{size, id})
			id++
		} else {
			disk = append(disk, Part{size, -1})

		}
	}
	disk = mergeParts(disk)
	display(disk)
	fmt.Println()
	for i := len(disk) - 1; i >= 0; i-- {
		ele := disk[i]

		distance := 0
		for d := 0; d < i; d++ {
			if disk[d].id == ele.id {
				break
			}
			distance += disk[d].size

		}

		if ele.id == -1 {
			continue
		}

		emptyIdx := 0
		for !(emptyIdx == len(disk) || disk[emptyIdx].id == -1 && disk[emptyIdx].size >= ele.size) {
			emptyIdx++
		}
		if emptyIdx == len(disk) || i < emptyIdx {
			continue
		}
		cp := Part{ele.size, ele.id}

		disk[i].id = -1
		disk[emptyIdx].size -= ele.size
		slices.Insert(disk, emptyIdx, cp)
		disk = mergeParts(disk)

		//	display(disk)
		//fmt.Println()

		distance2 := 0
		for d := 0; d < i; d++ {
			if distance2 >= distance {
				i = d
				break
			}
			distance2 += disk[d].size

		}

	}
	disk = mergeParts(disk)
	checksum := 0
	index := 0
	for _, part := range disk {
		for i := 0; i < part.size; i++ {
			if part.id > 0 {
				checksum += index * part.id
			}
			index++
		}

	}
	fmt.Println(checksum)
}
