package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var input []string

	for scanner.Scan() {
		v := scanner.Text()
		input = append(input, v)
	}

	file.Close()

	parts(input)
}

type point struct {
	X int
	Y int
}

func parts(input []string) {
	start := point{}
	startDir := 0
	dirs := []point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirChanges := map[rune][]int{}
	dirChanges['7'] = []int{1, -1, -1, 2}
	dirChanges['L'] = []int{-1, 0, 3, -1}
	dirChanges['F'] = []int{-1, -1, 1, 0}
	dirChanges['J'] = []int{3, 2, -1, -1}

	for x, iv := range input {
		for y, c := range iv {
			if c == 'S' {
				start = point{x, y}
				if input[x+1][y] == 'L' || input[x+1][y] == 'J' || input[x+1][y] == '|' {
					startDir = 1
				} else if input[x-1][y] == 'F' || input[x+1][y] == '7' || input[x+1][y] == '|' {
					startDir = 3
				} else {
					startDir = 0
				}
			}
		}
	}

	mem := map[point]int{}
	current := start
	currentDir := startDir
	count := 0

	for {
		if _, exists := mem[current]; exists {
			break
		}

		mem[current] = count
		next := point{current.X + dirs[currentDir].X, current.Y + dirs[currentDir].Y}
		if _, exists := dirChanges[rune(input[next.X][next.Y])]; exists {
			currentDir = dirChanges[rune(input[next.X][next.Y])][currentDir]
		}
		current = next
		count++
	}

	fmt.Println(count / 2)

	count2 := 0
	for x, iv := range input {
		crossed := 0
		for y, c := range iv {
			if _, exists := mem[point{x, y}]; exists {
				if c == '|' || c == 'L' || c == 'J' || c == 'S' {
					crossed++
				}
			} else {
				if crossed%2 == 1 {
					count2++
				}
			}
		}
	}
	fmt.Println(count2)
}
