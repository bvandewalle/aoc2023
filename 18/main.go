package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	part1(input, true)
	part1(input, false)
}

type point struct {
	x int
	y int
}

func part1(input []string, part1 bool) {
	directions := []point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	count := 0
	perimeter := 0
	mem := []point{point{0, 0}}
	currentPoint := point{0, 0}

	for _, iv := range input {
		a := strings.Split(iv, " ")

		le, _ := strconv.Atoi(a[1])
		dir := -1

		if part1 {
			switch a[0] {
			case "R":
				dir = 0
			case "D":
				dir = 1
			case "L":
				dir = 2
			case "U":
				dir = 3
			}
		} else {
			hex := ""
			for i := 2; i <= 6; i++ {
				hex += string(a[2][i])
			}
			lee, _ := strconv.ParseInt(hex, 16, 64)
			le = int(lee)

			dir, _ = strconv.Atoi(string(a[2][7]))
		}

		nextPoint := point{currentPoint.x + (le * directions[dir].x), currentPoint.y + (le * directions[dir].y)}
		mem = append(mem, nextPoint)
		currentPoint = nextPoint
		perimeter += le
	}

	for i, v := range mem {
		if i == len(mem)-1 {
			break
		}
		count += (v.y + mem[i+1].y) * (v.x - mem[i+1].x)
	}

	fmt.Println(abs(perimeter/2) + abs(count/2) + 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
