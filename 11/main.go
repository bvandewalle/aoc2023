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

	part1(input)
}

type point struct {
	X int
	Y int
}

func part1(input []string) {
	doubleLine := map[int]bool{}
	doubleCol := map[int]bool{}

	galaxies := []point{}

	for y, iv := range input {
		for x, v := range iv {
			if v != '.' {
				galaxies = append(galaxies, point{x, y})
			}
		}
	}

OUT1:
	for y, iv := range input {
		for _, v := range iv {
			if v != '.' {
				continue OUT1
			}
		}
		doubleLine[y] = true
	}

OUT2:
	for x, _ := range input[0] {
		for y, _ := range input {
			if input[y][x] != '.' {
				continue OUT2
			}
		}
		doubleCol[x] = true
	}

	for _, v := range []int{1, 1000000 - 1} {
		dist := 0
		for i, g1 := range galaxies {
			for j := i + 1; j < len(galaxies); j++ {
				g2 := galaxies[j]
				dist += abs(g1.X-g2.X) + abs(g1.Y-g2.Y)
				dist += addExpanded(g1.X, g2.X, doubleCol, v)
				dist += addExpanded(g1.Y, g2.Y, doubleLine, v)
			}
		}
		fmt.Println(dist)
	}
}

func addExpanded(a, b int, mem map[int]bool, amount int) int {
	if a > b {
		a, b = b, a
	}

	count := 0
	for i := 0; i <= (b - a); i++ {
		if _, e := mem[a+i]; e {
			count += amount
		}
	}
	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
