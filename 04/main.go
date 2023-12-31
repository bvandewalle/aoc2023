package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	parts(input)
}

func parts(input []string) {
	count1 := 0
	count2 := 0
	cardcount := map[int]int{}
	cardwon := map[int][]int{}

	for i, iv := range input {
		m := map[int]int{}
		cardcount[i]++
		a := strings.Split(iv, " | ")
		aa := strings.Split(a[0], ": ")
		winning := strings.Split(aa[1], " ")
		got := strings.Split(a[1], " ")

		for _, v := range winning {
			c, _ := strconv.Atoi(v)
			m[c]++
		}

		for _, v := range got {
			c, _ := strconv.Atoi(v)
			m[c]++
		}

		matches := 0
		for k, v := range m {
			if v == 2 && k != 0 {
				matches++
			}
		}
		for j := 1; j <= matches; j++ {
			cardwon[i] = append(cardwon[i], i+j)
		}
		if matches > 0 {
			count1 += int(math.Pow(2, float64(matches-1)))
		}
	}

	for i := 0; i <= len(input); i++ {
		for _, vv := range cardwon[i] {
			cardcount[vv] += cardcount[i]
		}
	}
	for _, v := range cardcount {
		count2 += v
	}

	fmt.Println(count1)
	fmt.Println(count2)
}
