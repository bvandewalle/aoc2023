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
	part2(input)
}

func part1(input []string) {
	time := []int{44, 70, 70, 80}
	dist := []int{283, 1134, 1134, 1491}

	count := 1

	for i, iv := range time {
		score := 0
		for j := 0; j < iv; j++ {
			cdist := j * (iv - j)
			if cdist >= dist[i] {
				score++
			}
		}
		count *= score
	}

	fmt.Println(count)
}

func part2(input []string) {
	time := 44707080
	dist := 283113411341491

	count := 0

	for j := 0; j < time; j++ {
		cdist := j * (time - j)
		if cdist >= dist {
			count++
		}
	}

	fmt.Println(count)
}
