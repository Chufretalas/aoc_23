package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func D1P1() {
	file, _ := os.Open("./inputs/d1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {
		cv_str := "" // calibration_value

		for i := 0; i < len(scanner.Text()); i++ {
			candidate := string(scanner.Text()[i])
			_, err := strconv.Atoi(candidate)
			if err == nil {
				cv_str += candidate
				break
			}
		}

		for i := len(scanner.Text()) - 1; i >= 0; i-- {
			candidate := string(scanner.Text()[i])
			_, err := strconv.Atoi(candidate)
			if err == nil {
				cv_str += candidate
				break
			}
		}

		cv, err := strconv.Atoi(cv_str)

		if err != nil {
			panic(err.Error())
		}

		sum += cv
	}

	fmt.Printf("sum: %v\n", sum)
}

var n_strs = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9"}

func D1P2() {
	file, _ := os.Open("./inputs/d1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {

		// find spelled numbers
		first_idx := len(scanner.Text())
		first_number := ""

		last_idx := -1
		last_number := ""
		last_number_spelled := ""

		for key, n_str := range n_strs {
			if idx := strings.Index(scanner.Text(), key); idx != -1 {
				if idx < first_idx {
					first_idx = idx
					first_number = n_str
				}
			}

			if idx := strings.LastIndex(scanner.Text(), key); idx != -1 {
				if idx > last_idx {
					last_idx = idx
					last_number = n_str
					last_number_spelled = key
				}
			}
		}

		// finding normal numbers
		for _, char := range scanner.Text()[:first_idx] {
			candidate := string(char)
			_, err := strconv.Atoi(candidate)
			if err == nil {
				first_number = candidate
				break
			}
		}

		for i := len(scanner.Text()) - 1; i >= last_idx+len(last_number_spelled); i-- {
			candidate := string(scanner.Text()[i])
			_, err := strconv.Atoi(candidate)
			if err == nil {
				last_number = candidate
				break
			}
		}

		cv, err := strconv.Atoi(first_number + last_number)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(cv)

		sum += cv
	}

	fmt.Printf("sum: %v\n", sum)
}
