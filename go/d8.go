package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func D8P1() {
	f, _ := os.Open("./inputs/d8.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	scan.Scan()

	directions := strings.Split(scan.Text(), "")

	scan.Scan()

	mmap := make(map[string][]string)
	for scan.Scan() {

		split := strings.Split(scan.Text(), " = ")

		node := split[0]

		mmap[node] = strings.Split(strings.Trim(split[1], "()"), ", ")
	}

	pos := "AAA"
	moves := 0
	i := 0
	getLR := map[string]int{"L": 0, "R": 1}
	for pos != "ZZZ" {
		pos = mmap[pos][getLR[directions[i]]]
		moves++
		i++
		if i >= len(directions) {
			i = 0
		}
	}

	fmt.Printf("moves: %v\n", moves)
}

// --------------------------------------------- PART 2 --------------------------------------------- //

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(ints []int) int {
	result := ints[0] * ints[1] / GCD(ints[0], ints[1])

	if len(ints) > 2 {
		for _, v := range ints[2:] {
			result = LCM([]int{result, v})
		}
	}

	return result
}

func D8P2() {
	f, _ := os.Open("./inputs/d8.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	scan.Scan()

	directions := strings.Split(scan.Text(), "")

	scan.Scan()

	pos := make([]string, 0)
	mmap := make(map[string][]string)
	for scan.Scan() {

		split := strings.Split(scan.Text(), " = ")

		node := split[0]
		if strings.HasSuffix(node, "A") {
			pos = append(pos, node)
		}

		mmap[node] = strings.Split(strings.Trim(split[1], "()"), ", ")
	}

	getLR := map[string]int{"L": 0, "R": 1}
	solutions := make([]int, len(pos))

	for idx := range pos {
		moves := 0
		i := 0
		for !strings.HasSuffix(pos[idx], "Z") {
			pos[idx] = mmap[pos[idx]][getLR[directions[i]]]
			moves++
			i++
			if i >= len(directions) {
				i = 0
			}
		}
		solutions[idx] = moves
	}

	fmt.Printf("moves: %v\n", LCM(solutions))
}
