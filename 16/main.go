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

type point struct {
	x int
	y int
}

func helper(input []string, visited map[point][]int, currentPoint point, currentDirection int) {
	directions := []point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	if v, e := visited[currentPoint]; e {
		for _, vv := range v {
			if vv == currentDirection {
				return
			}
		}
	}
	visited[currentPoint] = append(visited[currentPoint], currentDirection)

	nextPoint := point{
		x: currentPoint.x + directions[currentDirection].x,
		y: currentPoint.y + directions[currentDirection].y,
	}

	if nextPoint.y < 0 || nextPoint.y >= len(input[0]) {
		return
	}
	if nextPoint.x < 0 || nextPoint.x >= len(input) {
		return
	}

	s := input[nextPoint.x][nextPoint.y]

	switch currentDirection {
	case 0:
		switch s {
		case '|':
			helper(input, visited, nextPoint, 1)
			helper(input, visited, nextPoint, 3)
		case '\\':
			helper(input, visited, nextPoint, 1)
		case '/':
			helper(input, visited, nextPoint, 3)
		default:
			helper(input, visited, nextPoint, currentDirection)
		}
	case 1:
		switch s {
		case '-':
			helper(input, visited, nextPoint, 0)
			helper(input, visited, nextPoint, 2)
		case '\\':
			helper(input, visited, nextPoint, 0)
		case '/':
			helper(input, visited, nextPoint, 2)
		default:
			helper(input, visited, nextPoint, currentDirection)
		}
	case 2:
		switch s {
		case '|':
			helper(input, visited, nextPoint, 1)
			helper(input, visited, nextPoint, 3)
		case '\\':
			helper(input, visited, nextPoint, 3)
		case '/':
			helper(input, visited, nextPoint, 1)
		default:
			helper(input, visited, nextPoint, currentDirection)
		}
	case 3:
		switch s {
		case '-':
			helper(input, visited, nextPoint, 0)
			helper(input, visited, nextPoint, 2)
		case '\\':
			helper(input, visited, nextPoint, 2)
		case '/':
			helper(input, visited, nextPoint, 0)
		default:
			helper(input, visited, nextPoint, currentDirection)
		}
	}

}

func part1(input []string) {
	currentPoint := point{0, -1}
	currentDirection := 0
	visited := map[point][]int{}

	helper(input, visited, currentPoint, currentDirection)

	fmt.Println(len(visited) - 1)
}

func part2(input []string) {
	count := 0

	currentPoint := point{}
	currentDirection := 0

	for x := 0; x < 4; x++ {
		for i := range input {
			switch x {
			case 0:
				currentPoint = point{i, -1}
				currentDirection = 0
			case 1:
				currentPoint = point{i, len(input)}
				currentDirection = 2
			case 2:
				currentPoint = point{-1, i}
				currentDirection = 1
			case 3:
				currentPoint = point{len(input), i}
				currentDirection = 3
			}

			visited := map[point][]int{}
			helper(input, visited, currentPoint, currentDirection)
			if len(visited) > count {
				count = len(visited)
			}
		}
	}

	fmt.Println(count - 1)
}
