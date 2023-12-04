package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Reveal struct {
	R int // red
	G int // green
	B int // blue
}

func ParseReveal(c chan Reveal, text string) {
	rev := Reveal{R: 0, G: 0, B: 0}
	for _, t := range strings.Split(strings.Trim(text, " "), ", ") {
		split := strings.Split(t, " ")

		n, _ := strconv.Atoi(split[0])

		switch split[1] {
		case "red":
			rev.R = n

		case "green":
			rev.G = n

		case "blue":
			rev.B = n
		}
	}
	c <- rev
}

func D2P1() {
	file, _ := os.Open("./inputs/d2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	sum := 0

	id := 1
	for scanner.Scan() {
		c := make(chan Reveal)
		reveals_str := strings.Split(strings.Split(scanner.Text(), ":")[1], ";")

		// there's no reason to go async, but I wanted to have some fun
		for _, text := range reveals_str {
			go ParseReveal(c, text)
		}

		failed := false
		for i := 0; i < len(reveals_str); i++ {
			reveal := <-c

			if reveal.R > 12 || reveal.G > 13 || reveal.B > 14 {
				failed = true
				break
			}
		}

		if failed {
			id++
			continue
		}

		sum += id
		id++
	}
	fmt.Println(sum)
}

func D2P2() {
	file, _ := os.Open("./inputs/d2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {
		c := make(chan Reveal)
		reveals_str := strings.Split(strings.Split(scanner.Text(), ":")[1], ";")

		for _, text := range reveals_str {
			go ParseReveal(c, text)
		}

		minimum := Reveal{R: 0, G: 0, B: 0}
		for i := 0; i < len(reveals_str); i++ {
			reveal := <-c

			if minimum.R < reveal.R {
				minimum.R = reveal.R
			}

			if minimum.G < reveal.G {
				minimum.G = reveal.G
			}

			if minimum.B < reveal.B {
				minimum.B = reveal.B
			}
		}
		sum += minimum.R * minimum.G * minimum.B
	}
	fmt.Println(sum)
}
