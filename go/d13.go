package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func equalSlices(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func findReflections(m [][]string, c chan int) int {

	points := 0

	// find horizontal
	for i := 1; i < len(m[0]); i++ {
		isReflection := true
		for _, line := range m {
			s1 := make([]string, i)
			s2 := make([]string, len(line)-i)
			copy(s1, line[:i])
			copy(s2, line[i:])
			size := min(len(s1), len(s2))
			slices.Reverse(s1)
			if !equalSlices(s1[:size], s2[:size]) {
				isReflection = false
				break
			}
		}
		if isReflection {
			points += i
			// fmt.Println(i, " is a horizontal reflection point")
		}
	}

	// find vertical

	for i := 1; i < len(m); i++ {

		s1 := make([]string, 0)
		s2 := make([]string, 0)

		for _, line := range m[:i] {
			s1 = append(s1, strings.Join(line, ""))
		}

		for _, line := range m[i:] {
			s2 = append(s2, strings.Join(line, ""))
		}

		size := min(len(s1), len(s2))
		slices.Reverse(s1)

		if equalSlices(s1[:size], s2[:size]) {
			points += i * 100
			// fmt.Println(i, " is a vertical reflection point")
		}
	}
	c <- points
	return points
}

func D13P1() {
	f, _ := os.Open("./inputs/d13.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	res := 0

	maps := 0

	m := make([][]string, 0)
	c := make(chan int)
	for scan.Scan() {
		if scan.Text() == "" {
			maps++
			go findReflections(m, c)
			m = make([][]string, 0)
			continue
		}
		m = append(m, strings.Split(scan.Text(), ""))
	}
	go findReflections(m, c)
	maps++

	for i := 0; i < maps; i++ {
		res += <-c
	}

	fmt.Printf("res: %v\n", res)
}

// --------------------------------------------- PART 2 --------------------------------------------- //

// returns the number of differences
func equalSlices2(s1, s2 []string) int {

	diff := 0

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}

	return diff
}

func findReflections2(m [][]string, c chan int) int {

	points := 0

	// find horizontal
	for i := 1; i < len(m[0]); i++ {
		diff := 0
		for _, line := range m {
			s1 := make([]string, i)
			s2 := make([]string, len(line)-i)
			copy(s1, line[:i])
			copy(s2, line[i:])
			size := min(len(s1), len(s2))
			slices.Reverse(s1)
			diff += equalSlices2(s1[:size], s2[:size])
		}
		if diff == 1 {
			// fmt.Println(i, " is a horizontal reflection point")
		}
	}

	// find vertical

	for i := 1; i < len(m); i++ {

		s1 := make([]string, 0)
		s2 := make([]string, 0)

		for _, line := range m[:i] {
			s1 = append(s1, strings.Join(line, ""))
		}

		for _, line := range m[i:] {
			s2 = append(s2, strings.Join(line, ""))
		}

		size := min(len(s1), len(s2))
		slices.Reverse(s1)

		diff := 0

		for j := 0; j < size; j++ {
			diff += equalSlices2(strings.Split(s1[j], ""), strings.Split(s2[j], ""))
		}

		if diff == 1 {
			points += i * 100
			// fmt.Println(i, " is a vertical reflection point")
		}
	}
	c <- points
	return points
}

// works with the example at least
func D13P2() {
	f, _ := os.Open("./inputs/d13.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	res := 0

	maps := 0

	m := make([][]string, 0)
	c := make(chan int)
	for scan.Scan() {
		if scan.Text() == "" {
			maps++
			go findReflections2(m, c)
			m = make([][]string, 0)
			continue
		}
		m = append(m, strings.Split(scan.Text(), ""))
	}
	go findReflections2(m, c)
	maps++

	for i := 0; i < maps; i++ {
		res += <-c
	}

	fmt.Printf("res: %v\n", res)
}
