package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	//part1(input)
	part2(input)
}

func part1(input []string) {
	count := -1

	a := strings.Split(strings.TrimPrefix(input[0], "seeds: "), " ")
	fmt.Println(a)

	for _, ss := range a {
		s, _ := strconv.Atoi(ss)
		fmt.Println(s)
		roundfinished := false

		for i := 2; i < len(input); i++ {
			if input[i] == "" {
				continue
			}
			if !unicode.IsDigit(rune(input[i][0])) {
				roundfinished = false
				continue
			}
			if roundfinished {
				continue
			}
			mm := strings.Split(input[i], " ")
			destination, _ := strconv.Atoi(mm[0])
			source, _ := strconv.Atoi(mm[1])
			l, _ := strconv.Atoi(mm[2])

			if source <= s && source+l > s {
				ds := s - source
				s = destination + ds
				roundfinished = true
			}
		}

		if count == -1 || s < count {
			count = s
		}
	}

	fmt.Println(count)
}

func part2(input []string) {
	count := -1

	a := strings.Split(strings.TrimPrefix(input[0], "seeds: "), " ")
	fmt.Println(a)

	for k := 0; k < len(a); k += 2 {
		s, _ := strconv.Atoi(a[k])
		sl, _ := strconv.Atoi(a[k+1])
		currentStart := s
		currentLength := sl
		currentDoneLength := 0

		roundfinished := false

		for currentDoneLength < sl {
			currentStart = s + currentDoneLength
			currentLength = sl - currentDoneLength
			fmt.Println("Next round ----")
			fmt.Println(currentStart)
			fmt.Println(currentLength)

			for i := 2; i < len(input); i++ {
				if input[i] == "" {
					continue
				}
				if !unicode.IsDigit(rune(input[i][0])) {
					roundfinished = false
					continue
				}
				if roundfinished {
					continue
				}
				mm := strings.Split(input[i], " ")
				destination, _ := strconv.Atoi(mm[0])
				source, _ := strconv.Atoi(mm[1])
				l, _ := strconv.Atoi(mm[2])

				if currentStart >= source && currentStart <= source+l-1 {
					ds := currentStart - source
					currentStart = destination + ds
					if currentLength > l {
						currentLength = l
					}
					roundfinished = true
					fmt.Println("A", currentStart, currentLength, currentDoneLength)
					continue
				}

				if source >= currentStart && source <= currentStart+currentLength-1 {
					currentLength = source - currentStart
					fmt.Println("B", currentStart, currentLength, currentDoneLength)
				}
			}
			if count == -1 || currentStart < count {
				count = currentStart
			}
			currentDoneLength += currentLength
		}

	}

	fmt.Println(count)
}
