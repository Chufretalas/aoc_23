package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseNumbers(s string) []int {
	parsed := make([]int, 0)
	for _, c := range strings.Split(s, " ") {
		v, err := strconv.Atoi(c)
		if err != nil {
			continue
		}
		parsed = append(parsed, v)
	}
	return parsed
}

func D4P1() {
	file, err := os.Open("./inputs/d4.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	scan := bufio.NewScanner(file)

	sum := 0
	for scan.Scan() {
		split := strings.Split(strings.Split(scan.Text(), ":")[1], " | ")
		winners := parseNumbers(split[0])
		have := parseNumbers(split[1])
		points := 0
		for _, n := range have {
			if slices.Contains(winners, n) {
				if points == 0 {
					points = 1
					continue
				}
				points *= 2
			}
		}
		sum += points
	}
	fmt.Printf("sum: %v\n", sum)
}

var (
	matches   map[int]int
	instances map[int]int
)

func processCard(card int) {
	for i := card + 1; i <= card+matches[card]; i++ {
		instances[i] += 1
		processCard(i)
	}
}

func D4P2() {
	file, err := os.Open("./inputs/d4.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	scan := bufio.NewScanner(file)

	matches = make(map[int]int)
	instances = make(map[int]int)
	cardNumber := 1
	for scan.Scan() {
		split := strings.Split(strings.Split(scan.Text(), ":")[1], " | ")
		winners := parseNumbers(split[0])
		have := parseNumbers(split[1])

		for _, n := range have {
			if slices.Contains(winners, n) {
				matches[cardNumber]++
			}
		}
		instances[cardNumber] = 1
		cardNumber++
	}

	for i := 1; i <= len(instances); i++ {
		processCard(i)
	}

	sum := 0
	for _, v := range instances {
		sum += v
	}

	fmt.Printf("sum: %v\n", sum)
}
