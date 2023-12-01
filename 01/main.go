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

	part1(input)
	part2(input)
}

func part1(input []string) {
	count := 0

	for _, iv := range input {
		s := ""
		for _, c := range iv {
			if unicode.IsDigit(c) {
				s += string(c)
				break
			}

		}

		for i := len(iv) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(iv[i])) {
				s += string(iv[i])
				break
			}

		}

		v, _ := strconv.Atoi(s)
		count += v
	}

	fmt.Println(count)
}

func part2(input []string) {
	new := []string{}
	for _, iv := range input {
		iv = strings.Replace(iv, "zero", "zero0zero", -1)
		iv = strings.Replace(iv, "one", "one1one", -1)
		iv = strings.Replace(iv, "two", "two2two", -1)
		iv = strings.Replace(iv, "three", "three3three", -1)
		iv = strings.Replace(iv, "four", "four4four", -1)
		iv = strings.Replace(iv, "five", "five5five", -1)
		iv = strings.Replace(iv, "six", "six6six", -1)
		iv = strings.Replace(iv, "seven", "seven7seven", -1)
		iv = strings.Replace(iv, "eight", "eight8eight", -1)
		iv = strings.Replace(iv, "nine", "nine9nine", -1)

		new = append(new, iv)
	}
	part1(new)
}
