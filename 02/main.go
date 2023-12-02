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

	part1(input)
	part2(input)
}

func part1(input []string) {
	count := 0

FIRST:
	for _, iv := range input {
		a := strings.Split(iv, ": ")
		s := strings.Split(a[1], "; ")

		for _, v := range s {
			// Games
			blue := 0
			red := 0
			green := 0
			c := strings.Split(v, ", ")
			for _, cc := range c {
				d := strings.Split(cc, " ")
				x, _ := strconv.Atoi(d[0])

				if d[1] == "blue" {
					blue = x
				}
				if d[1] == "green" {
					green = x
				}
				if d[1] == "red" {
					red = x
				}
			}

			if !(red <= 12 && green <= 13 && blue <= 14) {
				continue FIRST
			}
		}
		aa := strings.Split(a[0], "Game ")
		aaa, _ := strconv.Atoi(aa[1])
		count += aaa
	}

	fmt.Println(count)
}

func part2(input []string) {
	count := 0

	for _, iv := range input {
		a := strings.Split(iv, ": ")
		s := strings.Split(a[1], "; ")
		// Games
		blue := 0
		red := 0
		green := 0
		for _, v := range s {
			c := strings.Split(v, ", ")
			for _, cc := range c {
				d := strings.Split(cc, " ")
				x, _ := strconv.Atoi(d[0])

				if d[1] == "blue" {
					if x > blue {
						blue = x
					}
				}
				if d[1] == "green" {
					if x > green {
						green = x
					}
				}
				if d[1] == "red" {
					if x > red {
						red = x
					}
				}
			}
		}
		aa := blue * green * red
		count += aa
	}

	fmt.Println(count)
}
