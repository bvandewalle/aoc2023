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
	a := strings.Split(input[0], ",")

	for _, v := range a {
		val := 0
		for _, c := range v {
			val += int(c)
			val *= 17
			val %= 256
		}
		count += val
	}
	fmt.Println(count)
}

type entry struct {
	label string
	value int
}

func part2(input []string) {
	count := 0
	a := strings.Split(input[0], ",")
	mem := [][]entry{}
	for i := 0; i < 256; i++ {
		mem = append(mem, []entry{})
	}

	for _, v := range a {
		a := ""
		aa := -1
		if strings.HasSuffix(v, "-") {
			b := strings.Split(v, "-")
			a = b[0]
		} else {
			b := strings.Split(v, "=")
			a = b[0]
			bb, _ := strconv.Atoi(b[1])
			aa = bb
		}

		box := 0
		for _, c := range a {
			box += int(c)
			box *= 17
			box %= 256
		}

		if aa == -1 {
			newLine := []entry{}
			for _, v := range mem[box] {
				if v.label != a {
					newLine = append(newLine, v)
				}
			}
			mem[box] = newLine
		} else {
			e := entry{
				label: a,
				value: aa,
			}

			found := false
			for i, v := range mem[box] {
				if v.label == a {
					mem[box][i] = e
					found = true
					break
				}
			}
			if !found {
				mem[box] = append(mem[box], e)
			}
		}
	}
	for i, v := range mem {
		for j, l := range v {
			count += (i + 1) * (j + 1) * l.value
		}
	}

	fmt.Println(count)
}
