package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var sch []string

var notsymbols []string

type Candidate struct {
	V     int // nart number value
	Line  int // index of the line the number was in
	Start int // index whre the number starts
	End   int // index where the number ends
}

func check(c Candidate) int {
	start := c.Start - 1
	if start == -1 {
		start = c.Start
	}
	end := c.End + 1
	if end == len(sch) {
		end = c.End
	}

	// check top
	if line := c.Line - 1; line != -1 {
		for _, char := range sch[line][start : end+1] {
			if !slices.Contains(notsymbols, string(char)) {
				return c.V
			}
		}
	}
	// check sides
	if start != c.Start {
		if !slices.Contains(notsymbols, string(sch[c.Line][start])) {
			return c.V
		}
	}
	if end != c.End {
		if !slices.Contains(notsymbols, string(sch[c.Line][end])) {
			return c.V
		}
	}

	// check bottom
	if line := c.Line + 1; line != len(sch) {
		for _, char := range sch[line][start : end+1] {
			if !slices.Contains(notsymbols, string(char)) {
				return c.V
			}
		}
	}

	return 0
}

func D3P1() {
	file, _ := os.Open("./inputs/d3.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	notsymbols = []string{".", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	sch = make([]string, 0)
	candidates := make([]Candidate, 0)
	line_idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		sch = append(sch, line)

		// get all number from every line
		start_idx := -1
		n_str := ""
		for idx, char := range line {
			s := string(char)
			_, err := strconv.Atoi(s)
			if err == nil {
				if start_idx == -1 {
					start_idx = idx
				}
				n_str += s
				if idx == len(line)-1 {
					v, _ := strconv.Atoi(n_str)
					candidates = append(candidates, Candidate{v, line_idx, start_idx, idx})
					start_idx = -1
					n_str = ""
					break
				}
				continue
			}

			if n_str != "" {
				v, _ := strconv.Atoi(n_str)
				candidates = append(candidates, Candidate{v, line_idx, start_idx, idx - 1})
				start_idx = -1
				n_str = ""
			}
		}
		line_idx++
	}

	sum := 0
	for _, candidate := range candidates {
		sum += check(candidate)
	}

	fmt.Println(sum)
}

// ---------------------------------------------- Part 2 ----------------------------------------------- //

type Candidate2 struct {
	V     int // nart number value
	Line  int // index of the line the number was in
	Start int // index whre the number starts
	End   int // index where the number ends
}

type Gear struct {
	line    int
	col     int
	numbers []int
}

var gears []Gear

func check2(c Candidate) {
	start := c.Start - 1
	if start == -1 {
		start = c.Start
	}
	end := c.End + 1
	if end == len(sch) {
		end = c.End
	}

	// check top
	if line := c.Line - 1; line != -1 {
		for idx, char := range sch[line][start : end+1] {
			if string(char) == "*" {
				for g_idx, gear := range gears {
					if gear.line == line && gear.col == start+idx {
						gears[g_idx].numbers = append(gears[g_idx].numbers, c.V)
					}
				}
			}
		}
	}
	// check sides
	if start != c.Start {
		if string(sch[c.Line][start]) == "*" {
			for g_idx, gear := range gears {
				if gear.line == c.Line && gear.col == start {
					gears[g_idx].numbers = append(gears[g_idx].numbers, c.V)
				}
			}
		}
	}
	if end != c.End {
		if string(sch[c.Line][end]) == "*" {
			for g_idx, gear := range gears {
				if gear.line == c.Line && gear.col == end {
					gears[g_idx].numbers = append(gears[g_idx].numbers, c.V)
				}
			}
		}
	}

	// check bottom
	if line := c.Line + 1; line != len(sch) {
		for idx, char := range sch[line][start : end+1] {
			if string(char) == "*" {
				for g_idx, gear := range gears {
					if gear.line == line && gear.col == start+idx {
						gears[g_idx].numbers = append(gears[g_idx].numbers, c.V)
					}
				}
			}
		}
	}
}

func D3P2() {
	file, _ := os.Open("./inputs/d3.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	notsymbols = []string{".", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	sch = make([]string, 0)
	candidates := make([]Candidate, 0)
	gears = make([]Gear, 0)
	line_idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		sch = append(sch, line)

		// get all number from every line
		start_idx := -1
		n_str := ""
		for idx, char := range line {
			s := string(char)

			if s == "*" {
				gears = append(gears, Gear{line_idx, idx, make([]int, 0)})
			}

			_, err := strconv.Atoi(s)
			if err == nil {
				if start_idx == -1 {
					start_idx = idx
				}
				n_str += s
				if idx == len(line)-1 {
					v, _ := strconv.Atoi(n_str)
					candidates = append(candidates, Candidate{v, line_idx, start_idx, idx})
					start_idx = -1
					n_str = ""
					break
				}
				continue
			}

			if n_str != "" {
				v, _ := strconv.Atoi(n_str)
				candidates = append(candidates, Candidate{v, line_idx, start_idx, idx - 1})
				start_idx = -1
				n_str = ""
			}
		}
		line_idx++
	}

	for _, candidate := range candidates {
		check2(candidate)
	}

	sum := 0
	for _, gear := range gears {
		if len(gear.numbers) == 2 {
			sum += gear.numbers[0] * gear.numbers[1]
		}
	}

	fmt.Printf("sum: %v\n", sum)
}
