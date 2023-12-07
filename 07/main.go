package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/icza/abcsort"
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
	sorter := abcsort.New("123456789TJQKA")
	count := 0
	mem := map[string]int{}
	memClassified := map[int][]string{}

	for _, iv := range input {
		pair := strings.Split(iv, " ")
		value, _ := strconv.Atoi(pair[1])
		mem[pair[0]] = value

		matches := map[rune]int{}
		for _, c := range pair[0] {
			matches[c]++
		}
		m := []int{}
		for _, v := range matches {
			m = append(m, v)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(m)))
		if len(m) == 1 {
			memClassified[m[0]*5] = append(memClassified[m[0]*5], pair[0])
		} else {
			memClassified[m[0]*5+m[1]] = append(memClassified[m[0]*5+m[1]], pair[0])
		}
	}

	keys := make([]int, 0)
	for k, _ := range memClassified {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	i := 1
	for _, k := range keys {
		sorter.Strings(memClassified[k])
		for _, s := range memClassified[k] {
			count += i * mem[s]
			i++
		}
	}
	fmt.Println(count)
}

func part2(input []string) {
	sorter := abcsort.New("J123456789TQKA")
	count := 0
	mem := map[string]int{}
	memClassified := map[int][]string{}

	for _, iv := range input {
		pair := strings.Split(iv, " ")
		value, _ := strconv.Atoi(pair[1])
		mem[pair[0]] = value

		matches := map[rune]int{}
		jockers := 0
		for _, c := range pair[0] {
			if c == 'J' {
				jockers++
			} else {
				matches[c]++
			}
		}
		m := []int{}
		for _, v := range matches {
			m = append(m, v)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(m)))
		if len(m) >= 1 {
			m[0] += jockers
		} else {
			m = append(m, jockers)
		}

		if len(m) == 1 {
			memClassified[m[0]*5] = append(memClassified[m[0]*5], pair[0])
		} else {
			memClassified[m[0]*5+m[1]] = append(memClassified[m[0]*5+m[1]], pair[0])
		}
	}

	keys := make([]int, 0)
	for k, _ := range memClassified {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	i := 1
	for _, k := range keys {
		sorter.Strings(memClassified[k])
		for _, s := range memClassified[k] {
			count += i * mem[s]
			i++
		}
	}
	fmt.Println(count)
}
