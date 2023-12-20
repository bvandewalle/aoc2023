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

	part1(input, false)
	part1(input, true)
}

type pulse struct {
	high      bool
	recipient string
	sender    string
}

func part1(input []string, part2 bool) {
	memFlipFlop := map[string][]string{}
	memConjunction := map[string][]string{}
	memConjunctionInput := map[string][]string{}
	memBroadcast := []string{}

	stateFlipFlop := map[string]bool{}
	stateConjunction := map[string]map[string]bool{}

	for _, iv := range input {
		r := strings.Split(iv, " -> ")
		s := strings.Split(r[1], ", ")
		dest := []string{}
		for _, d := range s {
			dest = append(dest, d)
		}

		if r[0][0] == '%' {
			memFlipFlop[strings.TrimPrefix(r[0], "%")] = dest
		} else if r[0][0] == '&' {
			memConjunction[strings.TrimPrefix(r[0], "&")] = dest
		} else {
			memBroadcast = dest
		}
	}

	for k, _ := range memConjunction {
		deps := []string{}
		for kk, vv := range memFlipFlop {
			for _, vvv := range vv {
				if vvv == k {
					deps = append(deps, kk)
				}
			}
		}
		for kk, vv := range memConjunction {
			for _, vvv := range vv {
				if vvv == k {
					deps = append(deps, kk)
				}
			}
		}
		memConjunctionInput[k] = deps
		stateConjunction[k] = map[string]bool{}
	}

	queue := []pulse{}
	count := map[bool]int{}
	i := 0
	cycles := []int{}

OUT:
	for {
		if i == 1000 && !part2 {
			break
		}

		count[false]++
		for _, v := range memBroadcast {
			queue = append(queue, pulse{
				high:      false,
				recipient: v,
				sender:    "broadcast",
			})
		}
		for {
			if len(queue) == 0 {
				break
			}
			next := queue[0]
			count[next.high]++
			queue = queue[1:]

			if (next.recipient == "rx" || next.recipient == "jg" || next.recipient == "rh" || next.recipient == "jm" || next.recipient == "hf") && !next.high {
				cycles = append(cycles, i+1)
			}

			if len(cycles) == 4 {
				break OUT
			}

			if v, e := memFlipFlop[next.recipient]; e {
				if !next.high {
					for _, s := range v {
						queue = append(queue, pulse{
							high:      !stateFlipFlop[next.recipient],
							recipient: s,
							sender:    next.recipient,
						})
					}
					stateFlipFlop[next.recipient] = !stateFlipFlop[next.recipient]
				}
			} else if v, e := memConjunction[next.recipient]; e {
				stateConjunction[next.recipient][next.sender] = next.high
				highAll := true
				for _, vv := range memConjunctionInput[next.recipient] {
					if !stateConjunction[next.recipient][vv] {
						highAll = false
					}
				}
				for _, vv := range v {
					queue = append(queue, pulse{
						high:      !highAll,
						recipient: vv,
						sender:    next.recipient,
					})
				}
			}

		}
		i++
	}

	if !part2 {
		fmt.Println(count[false] * count[true])
	} else {
		fmt.Println(cycles[0] * cycles[1] * cycles[2] * cycles[3])
	}
}
