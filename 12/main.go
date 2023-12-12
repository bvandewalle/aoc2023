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
	parts(input, 1)
	parts(input, 5)
}

type entry struct {
	currentChar        int
	currentBlockLength int
	currentCheckSum    int
}

func helper(input string, checksum []int, currentChar int, currentBlockLength int, currentCheckSum int, mem map[entry]int) int {
	e := entry{
		currentChar:        currentChar,
		currentBlockLength: currentBlockLength,
		currentCheckSum:    currentCheckSum,
	}
	if v, exists := mem[e]; exists {
		return v
	}

	if currentChar >= len(input) {
		if currentBlockLength == 0 && currentCheckSum == len(checksum) {
			return 1
		} else {
			if currentCheckSum == len(checksum)-1 && checksum[currentCheckSum] == currentBlockLength {
				return 1
			}
		}
		return 0
	}

	dot := func() int {
		if currentBlockLength != 0 {
			if currentCheckSum < len(checksum) && checksum[currentCheckSum] == currentBlockLength {
				return helper(input, checksum, currentChar+1, 0, currentCheckSum+1, mem)
			}
			return 0
		}
		return helper(input, checksum, currentChar+1, 0, currentCheckSum, mem)
	}

	dash := func() int {
		return helper(input, checksum, currentChar+1, currentBlockLength+1, currentCheckSum, mem)
	}

	toReturn := 0
	switch input[currentChar] {
	case '.':
		toReturn = dot()
	case '#':
		toReturn = dash()
	case '?':
		toReturn = dot() + dash()
	default:
		fmt.Println("unreachable")
	}
	mem[e] = toReturn
	return toReturn
}

func parts(input []string, repeat int) {
	count := 0

	for _, iv := range input {
		s := strings.Split(iv, " ")
		a := strings.Split(s[1], ",")
		cheksum := []int{}
		for _, c := range a {
			vv, _ := strconv.Atoi(c)
			cheksum = append(cheksum, vv)
		}

		newInput := ""
		newChecksum := []int{}
		for i := 0; i < repeat; i++ {
			newInput = newInput + s[0]
			if i != 4 && repeat == 5 {
				newInput += "?"
			}
			newChecksum = append(newChecksum, cheksum...)
		}
		mem := map[entry]int{}
		count += helper(newInput, newChecksum, 0, 0, 0, mem)
	}

	fmt.Println(count)
}
