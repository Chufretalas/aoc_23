package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func D6P1() {
	file, _ := os.Open("./inputs/d6.txt")
	defer file.Close()

	scan := bufio.NewScanner(file)

	reg := regexp.MustCompile(`(\d+)`)

	scan.Scan()
	times_str := reg.FindAllString(scan.Text(), -1)
	scan.Scan()
	dists_str := reg.FindAllString(scan.Text(), -1)

	result := 1
	for i := 0; i < len(times_str); i++ {
		time, _ := strconv.Atoi(times_str[i])
		dist, _ := strconv.Atoi(dists_str[i])
		ways2win := 0
		for j := 0; j < time; j++ {
			will_travel := j * (time - j)
			if will_travel > dist {
				ways2win++
				continue
			}

			if ways2win > 0 { // this means that you stopped winning and holding down the button for any more time will only get worse
				break
			}
		}

		result *= ways2win
	}
	fmt.Printf("result: %v\n", result)
}

func findWinners(ch chan int, start, end, time, dist int) {
	ways2win := 0
	for j := start; j < end; j++ {
		will_travel := j * (time - j)
		if will_travel > dist {
			ways2win++
			continue
		}

		if ways2win > 0 {
			break
		}
	}
	ch <- ways2win
}

func D6P2() {
	file, _ := os.Open("./inputs/d6.txt")
	defer file.Close()

	scan := bufio.NewScanner(file)

	scan.Scan()
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(scan.Text(), ": ")[1], " ", ""))
	scan.Scan()
	dist, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(scan.Text(), ": ")[1], " ", ""))

	ch := make(chan int)
	div := time / 4

	for i := 0; i <= 3; i++ {
		fmt.Printf("start: %v, end: %v\n", 0+div*i, div*(i+1))
		go findWinners(ch, 0+div*i, div*(i+1), time, dist)
	}
	go findWinners(ch, div*4, time, time, dist)

	ways2win := 0

	for i := 0; i <= 4; i++ {
		ways2win += <-ch
	}

	fmt.Printf("ways2win: %v\n", ways2win)
}
