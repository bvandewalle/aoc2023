package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
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

func parts(input []string) {
	count1, count2 := 0, 0
	current := ""
	currentSymbolAdjacent := false
	currentMultAdjacent := false
	multx, multy := -1, -1
	m := map[int][]int{}

	for x, iv := range input {
		for y, c := range iv {
			if unicode.IsDigit(c) {
				current += string(c)
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						if x+dx >= 0 && x+dx < len(input) && y+dy >= 0 && y+dy < len(input[0]) {
							if !unicode.IsDigit(rune(input[x+dx][y+dy])) && input[x+dx][y+dy] != '.' {
								currentSymbolAdjacent = true
							}
							if input[x+dx][y+dy] == '*' {
								currentMultAdjacent = true
								multx = x + dx
								multy = y + dy
							}
						}
					}
				}
			} else {
				if current != "" {
					if currentSymbolAdjacent {
						v, _ := strconv.Atoi(current)
						count1 += v
					}
					if currentMultAdjacent {
						v, _ := strconv.Atoi(current)
						m[multx*len(input[0])+multy] = append(m[multx*len(input[0])+multy], v)
					}
				}
				current = ""
				currentSymbolAdjacent = false
				currentMultAdjacent = false
			}
		}
	}

	for _, v := range m {
		if len(v) == 2 {
			t := 1
			for _, vv := range v {
				t *= vv
			}
			count2 += t
		}
	}

	fmt.Println(count1)
	fmt.Println(count2)
}
