package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Coord2D struct {
	L int
	C int
}

func (c *Coord2D) Set(line, col int) {
	c.L = line
	c.C = col
}

func (c Coord2D) Next(d Direction) Coord2D {
	switch d {
	case UP:
		return Coord2D{c.L - 1, c.C}
	case DOWN:
		return Coord2D{c.L + 1, c.C}
	case LEFT:
		return Coord2D{c.L, c.C - 1}
	case RIGHT:
		return Coord2D{c.L, c.C + 1}
	}
	panic("??????????")
}

func (c *Coord2D) Move(d Direction) {
	switch d {
	case UP:
		c.L--
	case DOWN:
		c.L++
	case LEFT:
		c.C--
	case RIGHT:
		c.C++
	}
}

type Matrix2D [][]string

func (m Matrix2D) Get(c Coord2D) string {
	if c.C < 0 || c.C >= len(m[0]) || c.L < 0 || c.L >= len(m) {
		return ""
	}
	return m[c.L][c.C]
}

func (m Matrix2D) Set(c Coord2D, s string) {
	if c.C < 0 || c.C >= len(m[0]) || c.L < 0 || c.L >= len(m) {
		return
	}
	m[c.L][c.C] = s
}

func (m Matrix2D) Copy() Matrix2D {
	c := make(Matrix2D, len(m))
	for i := range m {
		c[i] = m[i][0:len(m[i])]
	}
	return c
}

func (m Matrix2D) String() string {
	final := ""
	for _, line := range m {
		for _, s := range line {
			final += fmt.Sprintf("%v", s)
		}
		final += "\n"
	}
	return final
}

func D10P1() {
	f, _ := os.Open("./inputs/d10.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	m := Matrix2D{}
	start := Coord2D{-1, -1}
	current := Coord2D{-1, -1}
	for scan.Scan() {
		m = append(m, strings.Split(scan.Text(), ""))
	}

	//find S
	for l, line := range m {
		for c, pipe := range line {
			if pipe == "S" {
				start.Set(l, c)
				current.Set(l, c)
				break
			}
		}
		if start.L != -1 {
			break
		}
	}

	// finding the loop
	guesses := []Direction{UP, DOWN, LEFT, RIGHT}
	for _, dir := range guesses {
		steps := 0
		failed := false
		for {
			current.Move(dir)
			steps++
			if current.C < 0 || current.C >= len(m[0]) || current.L < 0 || current.L >= len(m) {
				current = start
				break
			}
			next := m.Get(current)
			if next == "S" {
				// fmt.Printf("steps: %v\n", steps)
				fmt.Printf("Farthest: %v\n", steps/2)
				os.Exit(0)
			}
			switch dir {
			case UP:
				switch next {
				case "|":
					dir = UP
				case "F":
					dir = RIGHT
				case "7":
					dir = LEFT
				default:
					current = start
					failed = true
				}
			case DOWN:
				switch next {
				case "|":
					dir = DOWN
				case "L":
					dir = RIGHT
				case "J":
					dir = LEFT
				default:
					current = start
					failed = true
				}
			case LEFT:
				switch next {
				case "-":
					dir = LEFT
				case "L":
					dir = UP
				case "F":
					dir = DOWN
				default:
					current = start
					failed = true
				}
			case RIGHT:
				switch next {
				case "-":
					dir = RIGHT
				case "7":
					dir = DOWN
				case "J":
					dir = UP
				default:
					current = start
					failed = true
				}
			}
			if failed {
				break
			}
		}
	}
}

// --------------------------------------------- PART 2 --------------------------------------------- //

// Not ideia what is wrong with this one
func D10P2() {
	f, _ := os.Open("./inputs/d10.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	m := Matrix2D{}
	start := Coord2D{-1, -1}
	current := Coord2D{-1, -1}
	for scan.Scan() {
		m = append(m, strings.Split(scan.Text(), ""))
	}

	//find S
	for l, line := range m {
		for c, pipe := range line {
			if pipe == "S" {
				start.Set(l, c)
				current.Set(l, c)
				break
			}
		}
		if start.L != -1 {
			break
		}
	}

	var painted Matrix2D

	// finding the loop
	guesses := []Direction{UP, DOWN, LEFT, RIGHT}
	foundLoop := false
	for _, dir := range guesses {
		steps := 0
		failed := false
		painted = make(Matrix2D, len(m))
		for i := range m {
			painted[i] = make([]string, len(m[i]))
			copy(painted[i], m[i])
		}
		for {
			current.Move(dir)
			steps++
			if current.C < 0 || current.C >= len(m[0]) || current.L < 0 || current.L >= len(m) {
				current = start
				break
			}
			painted.Set(current, "🔵")
			next := m.Get(current)
			if next == "S" {
				foundLoop = true
				break
			}
			switch dir {
			case UP:
				switch next {
				case "|":
					dir = UP
				case "F":
					dir = RIGHT
				case "7":
					dir = LEFT
				default:
					current = start
					failed = true
				}
			case DOWN:
				switch next {
				case "|":
					dir = DOWN
				case "L":
					dir = RIGHT
				case "J":
					dir = LEFT
				default:
					current = start
					failed = true
				}
			case LEFT:
				switch next {
				case "-":
					dir = LEFT
				case "L":
					dir = UP
				case "F":
					dir = DOWN
				default:
					current = start
					failed = true
				}
			case RIGHT:
				switch next {
				case "-":
					dir = RIGHT
				case "7":
					dir = DOWN
				case "J":
					dir = UP
				default:
					current = start
					failed = true
				}
			}
			if failed {
				break
			}
		}
		if foundLoop {
			break
		}
	}

	// paint the map
	//left to right
	for l, line := range painted {
		for c, s := range line {
			if s == "🔵" {
				break
			}
			painted[l][c] = "🔴"
		}
	}
	// right to left
	for l := 0; l < len(painted); l++ {
		for c := len(painted[l]) - 1; c >= 0; c-- {
			s := painted[l][c]
			if s == "🔵" {
				break
			}
			painted[l][c] = "🔴"
		}
	}
	// top to bottom
	for c := 0; c < len(painted[0]); c++ {
		for l := 0; l < len(painted); l++ {
			s := painted[l][c]
			if s == "🔵" {
				break
			}
			painted[l][c] = "🔴"
		}
	}
	// bottom to top
	for c := 0; c < len(painted[0]); c++ {
		for l := len(painted) - 1; l >= 0; l-- {
			s := painted[l][c]
			if s == "🔵" {
				break
			}
			painted[l][c] = "🔴"
		}
	}

	// spread the red
	changes := -1
	for changes != 0 {
		changes = 0
		for l, line := range painted {
			for c, current := range line {
				if current != "🔴" {
					continue
				}
				coord := Coord2D{l, c}
				nexts := []Coord2D{coord.Next(UP), coord.Next(DOWN), coord.Next(LEFT), coord.Next(RIGHT)}
				for _, next := range nexts {
					s := painted.Get(next)
					if s == "" {
						continue
					}
					if s != "🔵" && s != "🔴" {
						changes++
						painted.Set(next, "🔴")
					}
				}
			}
		}
	}

	// count the remaining titles

	enclosed := 0
	for l, line := range painted {
		for c, s := range line {
			if s != "🔵" && s != "🔴" {
				painted.Set(Coord2D{l, c}, "🟡")
				enclosed++
			}
		}
	}

	fmt.Println(painted)
	fmt.Printf("enclosed: %v\n", enclosed)
}
