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

	parts(input)
}

func parts(input []string) {
	count1 := 0
	count2 := 0

	for _, iv := range input {
		nn := strings.Split(iv, " ")
		nums := []int{}
		for _, v := range nn {
			vv, _ := strconv.Atoi(v)
			nums = append(nums, vv)
		}

		firstN := []int{}
		lastN := []int{}

		for {
			finished := true
			newNums := []int{}

			for i, v := range nums {
				if v != 0 {
					finished = false
				}

				if i != 0 {
					newNums = append(newNums, nums[i]-nums[i-1])
				}
			}
			if finished {
				break
			}

			firstN = append(firstN, nums[0])
			lastN = append(lastN, nums[len(nums)-1])

			nums = newNums
		}

		n1 := 0
		n2 := 0
		for i := len(lastN) - 1; i >= 0; i-- {
			n1 += lastN[i]
			n2 = firstN[i] - n2

		}
		count1 += n1
		count2 += n2
	}

	fmt.Println(count1)
	fmt.Println(count2)
}
