package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

type point struct {
	x int
	y int
}

type state struct {
	x         int
	y         int
	direction int
	dirLength int
}

type stateBFS struct {
	x         int
	y         int
	direction int
	dirLength int
	score     int
}

func helperBFS(input []string, mem map[state]int, part2 bool) {
	directions := []point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	fifo := []stateBFS{
		stateBFS{
			x:         0,
			y:         0,
			direction: 0,
			dirLength: 0,
			score:     0,
		},
		stateBFS{
			x:         0,
			y:         0,
			direction: 1,
			dirLength: 0,
			score:     0,
		},
	}

OUT:
	for {
		if len(fifo) == 0 {
			return
		}

		currentState := fifo[0]
		fifo = fifo[1:]

		if currentState.y < 0 || currentState.y >= len(input[0]) {
			continue
		}
		if currentState.x < 0 || currentState.x >= len(input) {
			continue
		}

		ss, _ := strconv.Atoi(string(input[currentState.y][currentState.x]))
		newScore := currentState.score + ss

		toCheckMem := []state{}
		toCheckMem = append(toCheckMem, state{
			x:         currentState.x,
			y:         currentState.y,
			direction: currentState.direction,
			dirLength: currentState.dirLength,
		})

		if !part2 {
			if currentState.dirLength == 2 {
				toCheckMem = append(toCheckMem, state{
					x:         currentState.x,
					y:         currentState.y,
					direction: currentState.direction,
					dirLength: 1,
				})
			}
			if currentState.dirLength >= 1 {
				toCheckMem = append(toCheckMem, state{
					x:         currentState.x,
					y:         currentState.y,
					direction: currentState.direction,
					dirLength: 0,
				})
			}
		}

		for _, sss := range toCheckMem {
			if s, e := mem[sss]; e {
				if s <= newScore {
					continue OUT
				}
			}
		}

		mem[toCheckMem[0]] = newScore

		dl := (4 + (currentState.direction - 1)) % 4
		ndl := point{currentState.x + directions[dl].x, currentState.y + directions[dl].y}
		dr := (4 + (currentState.direction + 1)) % 4
		ndr := point{currentState.x + directions[dr].x, currentState.y + directions[dr].y}
		ndd := point{currentState.x + directions[currentState.direction].x, currentState.y + directions[currentState.direction].y}

		if !part2 {

			if currentState.dirLength < 2 {
				fifo = append(fifo, stateBFS{
					x:         ndd.x,
					y:         ndd.y,
					direction: currentState.direction,
					dirLength: currentState.dirLength + 1,
					score:     newScore,
				})
			}

			fifo = append(fifo, stateBFS{
				x:         ndl.x,
				y:         ndl.y,
				direction: dl,
				dirLength: 0,
				score:     newScore,
			})
			fifo = append(fifo, stateBFS{
				x:         ndr.x,
				y:         ndr.y,
				direction: dr,
				dirLength: 0,
				score:     newScore,
			})
		} else {
			if currentState.dirLength < 9 {
				fifo = append(fifo, stateBFS{
					x:         ndd.x,
					y:         ndd.y,
					direction: currentState.direction,
					dirLength: currentState.dirLength + 1,
					score:     newScore,
				})
			}
			if currentState.dirLength >= 3 {
				fifo = append(fifo, stateBFS{
					x:         ndl.x,
					y:         ndl.y,
					direction: dl,
					dirLength: 0,
					score:     newScore,
				})
				fifo = append(fifo, stateBFS{
					x:         ndr.x,
					y:         ndr.y,
					direction: dr,
					dirLength: 0,
					score:     newScore,
				})
			}
		}
	}

}

func part1(input []string) {
	count := -1
	mem := map[state]int{}

	helperBFS(input, mem, false)

	for k, v := range mem {
		if k.x == len(input)-1 && k.y == len(input)-1 {
			if v < count || count == -1 {
				count = v
			}
		}
	}

	f, _ := strconv.Atoi(string(input[0][0]))

	fmt.Println(count - f)
}

func part2(input []string) {
	count := -1
	mem := map[state]int{}

	helperBFS(input, mem, true)

	for k, v := range mem {
		if k.x == len(input)-1 && k.y == len(input)-1 {
			if k.dirLength >= 3 {
				if v < count || count == -1 {
					count = v
				}
			}
		}
	}

	f, _ := strconv.Atoi(string(input[0][0]))

	fmt.Println(count - f)
}
