package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isAllZeroes(arr []int) bool {
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

func D9P1() {
	f, _ := os.Open("./inputs/d9.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	sum := 0
	for scan.Scan() {
		inpt_str := strings.Split(scan.Text(), " ")

		history := make([][]int, 1)
		for _, s := range inpt_str {
			v, _ := strconv.Atoi(s)
			history[0] = append(history[0], v)
		}

		i := 0
		for !isAllZeroes(history[len(history)-1]) {
			newLine := make([]int, 0, len(history[i])-1)
			for j := 0; j < len(history[i])-1; j++ {
				newLine = append(newLine, history[i][j+1]-history[i][j])
			}
			history = append(history, newLine)
			i++
		}

		slices.Reverse(history)

		history[0] = append(history[0], 0)

		for i := 1; i < len(history); i++ {
			history[i] = append(history[i], history[i-1][len(history[i-1])-1]+history[i][len(history[i])-1])
		}

		next := history[len(history)-1][len(history[len(history)-1])-1]

		// fmt.Printf("history: %v\n", history)
		// fmt.Printf("next: %v\n", next)
		sum += next
	}
	fmt.Printf("sum: %v\n", sum)
}

// --------------------------------------------- PART 2 --------------------------------------------- //

func D9P2() {
	f, _ := os.Open("./inputs/d9.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	sum := 0
	for scan.Scan() {
		inpt_str := strings.Split(scan.Text(), " ")

		history := make([][]int, 1)
		for _, s := range inpt_str {
			v, _ := strconv.Atoi(s)
			history[0] = append(history[0], v)
		}

		i := 0
		for !isAllZeroes(history[len(history)-1]) {
			newLine := make([]int, 0, len(history[i])-1)
			for j := 0; j < len(history[i])-1; j++ {
				newLine = append(newLine, history[i][j+1]-history[i][j])
			}
			history = append(history, newLine)
			i++
		}

		slices.Reverse(history)

		for idx := range history {
			slices.Reverse(history[idx])
		}

		history[0] = append(history[0], 0)

		for i := 1; i < len(history); i++ {
			history[i] = append(history[i], history[i][len(history[i])-1]-history[i-1][len(history[i-1])-1])
		}

		next := history[len(history)-1][len(history[len(history)-1])-1]

		// fmt.Printf("history: %v\n", history)
		// fmt.Printf("next: %v\n", next)
		sum += next
	}
	fmt.Printf("sum: %v\n", sum)
}
