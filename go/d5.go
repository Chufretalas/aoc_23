package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// I added two extra empty lines to the end of the input file, just to make things easier

func D5P1() {
	file, _ := os.Open("./inputs/d5.txt")
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	scan.Scan()

	seeds_str := strings.Split(strings.Split(scan.Text(), ": ")[1], " ")

	seeds := make([]int, 0)
	has_converted := make([]bool, 0)

	for _, s := range seeds_str {
		v, _ := strconv.Atoi(s)
		seeds = append(seeds, v)
		has_converted = append(has_converted, false)
	}

	scan.Scan()
	scan.Scan()
	for scan.Scan() {
		if scan.Text() == "" {
			scan.Scan()
			for idx := range has_converted {
				has_converted[idx] = false
			}
			continue
		}

		dst := 0
		src := 0
		size := 0
		for idx, s := range strings.Split(scan.Text(), " ") {
			v, _ := strconv.Atoi(s)
			switch idx {
			case 0:
				dst = v
			case 1:
				src = v
			case 2:
				size = v
			}
		}

		for idx, seed := range seeds {
			if seed >= src && seed <= src+size-1 && !has_converted[idx] {
				offset := seed - src
				seeds[idx] = dst + offset
				has_converted[idx] = true
			}
		}
	}

	lowest := 9999999999999

	for _, seed := range seeds {
		if seed < lowest {
			lowest = seed
		}
	}

	fmt.Printf("lowest: %v\n", lowest)

}

// TODO: does not work*
// *it would work if I could leave my computer running for a week
func D5P2() {
	file, _ := os.Open("./inputs/d5.txt")
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	scan.Scan()

	seeds_str := strings.Split(strings.Split(scan.Text(), ": ")[1], " ")

	seeds := make([]int, 0)
	has_converted := make([]bool, 0)

	j := 0
	for len(seeds) != int(len(seeds_str)/2) {
		if j >= len(seeds_str) {
			break
		}

		start, _ := strconv.Atoi(seeds_str[j])
		size, _ := strconv.Atoi(seeds_str[j+1])
		for k := start; k < start+size; k++ {
			fmt.Printf("Doing %v / %v\r", k-start, size)
			if !slices.Contains(seeds, k) {
				seeds = append(seeds, k)
				has_converted = append(has_converted, false)
			}
		}
		fmt.Println("")
		j += 2
	}

	fmt.Println("finished getting the seeds")

	scan.Scan()
	scan.Scan()
	for scan.Scan() {
		if scan.Text() == "" {
			scan.Scan()
			for idx := range has_converted {
				has_converted[idx] = false
			}
			continue
		}

		dst := 0
		src := 0
		size := 0
		for idx, s := range strings.Split(scan.Text(), " ") {
			v, _ := strconv.Atoi(s)
			switch idx {
			case 0:
				dst = v
			case 1:
				src = v
			case 2:
				size = v
			}
		}

		for idx, seed := range seeds {
			if seed >= src && seed <= src+size-1 && !has_converted[idx] {
				offset := seed - src
				seeds[idx] = dst + offset
				has_converted[idx] = true
			}
		}
	}

	lowest := 9999999999999

	for _, seed := range seeds {
		if seed < lowest {
			lowest = seed
		}
	}

	fmt.Printf("lowest: %v\n", lowest)

}
