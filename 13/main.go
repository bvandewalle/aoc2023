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

	parts(input)
}

func parts(input []string) {
	count1, count2 := 0, 0

	currentPattern := []string{}
	for _, iv := range input {
		if iv == "" {
			a, b := testPattern(currentPattern)
			count1 += a
			count2 += b
			currentPattern = []string{}
		} else {
			currentPattern = append(currentPattern, strings.Clone(iv))
		}
	}
	a, b := testPattern(currentPattern)
	count1 += a
	count2 += b

	fmt.Println(count1)
	fmt.Println(count2)
}

func testPattern(pattern []string) (int, int) {
	count1 := 0
	count2 := 0

OUT1:
	for i := 1; i < len(pattern); i++ {
		diffCount := 0
		for j := 0; j < len(pattern); j++ {
			if i-j-1 >= 0 && i+j < len(pattern) {
				for y := 0; y < len(pattern[0]); y++ {
					if pattern[i-j-1][y] != pattern[i+j][y] {
						diffCount++
						if diffCount > 1 {
							continue OUT1
						}
					}
				}
			}
		}
		if diffCount == 0 {
			count1 += 100 * i
		}
		if diffCount == 1 {
			count2 += 100 * i
		}
	}

OUT2:
	for i := 1; i < len(pattern[0]); i++ {
		diffCount := 0
		for j := 0; j < len(pattern[0]); j++ {
			if i-j-1 >= 0 && i+j < len(pattern[0]) {
				for y := 0; y < len(pattern); y++ {
					if pattern[y][i-j-1] != pattern[y][i+j] {
						diffCount++
						if diffCount > 1 {
							continue OUT2
						}
					}
				}
			}
		}
		if diffCount == 0 {
			count1 += i
		}
		if diffCount == 1 {
			count2 += i
		}
	}
	return count1, count2
}
