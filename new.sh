#! /bin/sh

if [ -z "$1" ]; then
  echo "Error: No folder name provided."
  exit 1
fi

# Check if the folder exists
if [ -d "day$1" ]; then
  echo "Error: Folder 'day$1' already exists."
  exit 1
fi

mkdir "day$1"
cd "day$1"
touch partial.txt full.txt

# Define template
template=$(
  cat <<EOF
package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("partial.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
	}
}
EOF
)

# Create Go files with the template
echo "$template" >part1.go
echo "$template" >part2.go
