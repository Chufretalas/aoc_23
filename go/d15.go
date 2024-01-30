package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getHash(s string) int {
	current := 0
	for _, c := range s {
		current += int(c)
		current *= 17
		current %= 256
	}
	return current
}

func D15P1() {
	f, _ := os.Open("./inputs/d15.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	scan.Scan()

	seq := strings.Split(scan.Text(), ",")

	sum := 0

	for _, step := range seq {
		sum += getHash(step)
	}

	fmt.Printf("sum: %v\n", sum)
}

type Lens struct {
	label string
	fl    int
}

func D15P2() {
	f, _ := os.Open("./inputs/d15.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	scan.Scan()

	seq := strings.Split(scan.Text(), ",")

	boxes := make([][]Lens, 256)

	for _, step := range seq {
		if step[len(step)-1] == '-' {
			label := step[:len(step)-1]
			hash := getHash(label)
			boxes[hash] = slices.DeleteFunc(boxes[hash], func(l Lens) bool {
				return l.label == label
			})
			continue
		}

		split := strings.Split(step, "=")

		label := split[0]
		fl, _ := strconv.Atoi((split[1]))

		hash := getHash(label)

		newLens := Lens{label: label, fl: fl}

		alreadyInBox := false

		for idx, lens := range boxes[hash] {
			if lens.label == newLens.label {
				boxes[hash][idx] = newLens
				alreadyInBox = true
				break
			}
		}

		if !alreadyInBox {
			boxes[hash] = append(boxes[hash], Lens{label: label, fl: fl})
		}
	}

	sum := 0

	for boxIdx, box := range boxes {
		for lensIdx, lens := range box {
			sum += (boxIdx + 1) * (lensIdx + 1) * lens.fl
		}
	}

	fmt.Printf("sum: %v\n", sum)
}
