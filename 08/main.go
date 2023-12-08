package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	//part1(input)
	part2(input)
}

func part1(input []string) {
	count := 0

	inst := input[0]
	//current := 0
	mem := map[string][]string{}
	current := "AAA"

	for i := 2; i < len(input); i++ {
		a := strings.Split(input[i], " = (")
		b := strings.Split(a[1], ", ")
		c := strings.Split(b[1], ")")
		mem[a[0]] = []string{b[0], c[0]}
	}

	for {
		if current == "ZZZ" {
			break
		}
		if inst[count%len(inst)] == 'L' {
			current = mem[current][0]
		} else {
			current = mem[current][1]
		}
		count++
	}

	fmt.Println(mem)
	fmt.Println(count)
}

func part2(input []string) {
	count := 0

	inst := input[0]
	//current := 0
	mem := map[string][]string{}
	current := []string{}

	for i := 2; i < len(input); i++ {
		a := strings.Split(input[i], " = (")
		b := strings.Split(a[1], ", ")
		c := strings.Split(b[1], ")")
		mem[a[0]] = []string{b[0], c[0]}

		if strings.HasSuffix(a[0], "A") {
			current = append(current, a[0])
		}
	}

	fmt.Println(current)

	cycles := map[int]int{}

	for {
		for i, v := range current {
			if strings.HasSuffix(v, "Z") {
				if _, exists := cycles[i]; !exists {
					cycles[i] = count
				}
			}
		}
		if len(cycles) >= len(current) {
			break
		}

		newCurrent := []string{}
		if inst[count%len(inst)] == 'L' {
			for _, v := range current {
				newCurrent = append(newCurrent, mem[v][0])
			}
		} else {
			for _, v := range current {
				newCurrent = append(newCurrent, mem[v][1])
			}
		}
		count++
		current = newCurrent
	}

	fmt.Println(current)
	fmt.Println(cycles)
	fmt.Println(LCM(cycles[0], cycles[1], cycles[2], cycles[3], cycles[4], cycles[5]))
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
