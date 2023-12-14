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
	count := 0

	for i := range input[0] {
		currentRow := len(input)
		for j := range input {
			if input[j][i] == 'O' {
				count += currentRow
				currentRow--
			}
			if input[j][i] == '#' {
				currentRow = len(input) - j - 1
			}
		}
	}

	fmt.Println(count)
}

func part2(input []string) {
	mem := map[string]int{}
	i := 0
	x := 0
	next := input
	for {
		next = tiltCycle(next)
		cc := Concat(next)
		if _, e := mem[cc]; e {
			x = mem[cc]
			break
		}

		mem[cc] = i
		i++
	}
	fmt.Println("Cycle", i-x)
	fmt.Println("Offset", x)
	fmt.Println("OtherSideOffset", (1000000000-x)%(i-x))
	fmt.Println("Total Cycles", (1000000000-x)%(i-x)+x)

	cycles := (1000000000-x)%(i-x) + x

	for h := 0; h < cycles; h++ {
		input = tiltCycle(input)
	}

	count := 0
	for i, v := range input {
		for _, c := range v {
			if c == 'O' {
				count += len(input) - i
			}
		}
	}

	fmt.Println(count)
}

func tiltCycle(input []string) []string {

	for a := 0; a < 2; a++ {
		n := []string{}
		for i := range input[0] {
			currentLine := ""
			currentRock := 0
			currentFilled := 0
			for j := range input {
				if input[j][i] == 'O' {
					currentRock++
				}
				if input[j][i] == '#' {
					for k := 0; k < currentRock; k++ {
						currentLine += "O"
						currentFilled++
					}
					currentRock = 0
					for k := currentFilled; k < j; k++ {
						currentLine += "."
						currentFilled++
					}
					currentLine += "#"
					currentFilled++
				}
			}

			for k := 0; k < currentRock; k++ {
				currentLine += "O"
				currentFilled++
			}
			currentRock = 0
			for k := currentFilled; k < len(input); k++ {
				currentLine += "."
				currentFilled++
			}

			n = append(n, currentLine)
		}

		input = n

	}

	for a := 2; a < 4; a++ {
		n := []string{}
		for i := range input[0] {
			currentLine := ""
			currentRock := 0
			currentFilled := 0
			for j := range input {
				if input[len(input)-j-1][i] == 'O' {
					currentRock++
				}
				if input[len(input)-j-1][i] == '#' {
					for k := 0; k < currentRock; k++ {
						currentLine += "O"
						currentFilled++
					}
					currentRock = 0
					for k := currentFilled; k < j; k++ {
						currentLine += "."
						currentFilled++
					}
					currentLine += "#"
					currentFilled++
				}
			}

			for k := 0; k < currentRock; k++ {
				currentLine += "O"
				currentFilled++
			}
			currentRock = 0
			for k := currentFilled; k < len(input); k++ {
				currentLine += "."
				currentFilled++
			}

			n = append(n, currentLine)
		}

		input = n

	}

	newInput := []string{}

	for i := len(input) - 1; i >= 0; i-- {
		newInput = append(newInput, Reverse(input[i]))
	}

	return newInput
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func Concat(s []string) string {
	n := ""
	for _, v := range s {
		n += v
	}
	return n
}
