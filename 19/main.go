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

type rule struct {
	c      string
	compar string
	num    int
	result string
}

func part1(input []string) {
	memRules := map[string][]rule{}
	memParts := []map[string]int{}
	secondPart := false
	count := 0

	for _, iv := range input {
		if iv == "" {
			secondPart = true
			continue
		}
		if !secondPart {
			ruleSet := []rule{}

			r := strings.Split(iv, "{")
			rr := strings.Split(r[1], "}")
			rrr := strings.Split(rr[0], ",")
			for i, ru := range rrr {
				a := strings.Split(ru, ":")

				if i == len(rrr)-1 {
					ruleSet = append(ruleSet, rule{
						compar: "L",
						result: ru,
					})
				} else if a[0][1] == '>' {
					aa := strings.Split(a[0], ">")
					vv, _ := strconv.Atoi(aa[1])
					ruleSet = append(ruleSet, rule{
						c:      aa[0],
						compar: ">",
						num:    vv,
						result: a[1],
					})
				} else {
					aa := strings.Split(a[0], "<")
					vv, _ := strconv.Atoi(aa[1])
					ruleSet = append(ruleSet, rule{
						c:      aa[0],
						compar: "<",
						num:    vv,
						result: a[1],
					})
				}
			}

			memRules[r[0]] = ruleSet
		} else {
			r := strings.Split(iv, "{")
			rr := strings.Split(r[1], "}")
			rrr := strings.Split(rr[0], ",")

			part := map[string]int{}
			for _, pa := range rrr {
				rrrr := strings.Split(pa, "=")
				vv, _ := strconv.Atoi(rrrr[1])
				part[rrrr[0]] = vv
			}
			memParts = append(memParts, part)
		}
	}

	for _, pa := range memParts {
		currentWorkflow := "in"
		for {
			if currentWorkflow == "A" {
				count += pa["x"] + pa["m"] + pa["a"] + pa["s"]
				break
			}
			if currentWorkflow == "R" {
				break
			}

			cw := memRules[currentWorkflow]

			result := ""
			for _, r := range cw {
				if r.compar == "L" {
					result = r.result
					break
				} else if r.compar == ">" {
					if pa[r.c] > r.num {
						result = r.result
						break
					}
				} else {
					if pa[r.c] < r.num {
						result = r.result
						break
					}
				}
			}
			currentWorkflow = result
		}
	}
	fmt.Println(count)
}

func opposite(r rule) rule {
	ret := rule{
		c:      r.c,
		result: r.result,
	}
	if r.compar == "<" {
		ret.compar = ">"
		ret.num = r.num - 1
	}
	if r.compar == ">" {
		ret.compar = "<"
		ret.num = r.num + 1
	}
	return ret
}

func part2Helper(memRules map[string][]rule, currentWorkflow string, ruleList []rule) [][]rule {
	if currentWorkflow == "R" {
		return [][]rule{}
	}
	if currentWorkflow == "A" {
		newRL := make([]rule, len(ruleList))
		copy(newRL, ruleList)
		return [][]rule{newRL}
	}

	cw := memRules[currentWorkflow]
	rulesOpposite := []rule{}

	toR := [][]rule{}
	for _, r := range cw {
		toR = append(toR, part2Helper(memRules, r.result, append(append(ruleList, rulesOpposite...), r))...)
		rulesOpposite = append(rulesOpposite, opposite(r))
	}
	return toR
}

func part2(input []string) {
	memRules := map[string][]rule{}
	secondPart := false
	count := 0

	for _, iv := range input {
		if iv == "" {
			secondPart = true
			continue
		}
		if !secondPart {
			ruleSet := []rule{}

			r := strings.Split(iv, "{")
			rr := strings.Split(r[1], "}")
			rrr := strings.Split(rr[0], ",")
			for i, ru := range rrr {
				a := strings.Split(ru, ":")

				if i == len(rrr)-1 {
					ruleSet = append(ruleSet, rule{
						compar: "L",
						result: ru,
					})
				} else if a[0][1] == '>' {
					aa := strings.Split(a[0], ">")
					vv, _ := strconv.Atoi(aa[1])
					ruleSet = append(ruleSet, rule{
						c:      aa[0],
						compar: ">",
						num:    vv,
						result: a[1],
					})
				} else {
					aa := strings.Split(a[0], "<")
					vv, _ := strconv.Atoi(aa[1])
					ruleSet = append(ruleSet, rule{
						c:      aa[0],
						compar: "<",
						num:    vv,
						result: a[1],
					})
				}
			}

			memRules[r[0]] = ruleSet
		}
	}

	m := part2Helper(memRules, "in", []rule{})

	for _, v := range m {
		ruleCount := 1
		mm := map[string][]int{}
		mm["x"] = []int{1, 4000}
		mm["m"] = []int{1, 4000}
		mm["a"] = []int{1, 4000}
		mm["s"] = []int{1, 4000}

		for _, vv := range v {
			if vv.compar == "<" {
				if mm[vv.c][1] >= vv.num {
					mm[vv.c][1] = vv.num - 1
				}
			}
			if vv.compar == ">" {
				if mm[vv.c][0] <= vv.num {
					mm[vv.c][0] = vv.num + 1
				}
			}
		}
		for _, d := range mm {
			ruleCount *= (d[1] - d[0]) + 1
		}
		count += ruleCount
	}

	fmt.Println(count)
}
