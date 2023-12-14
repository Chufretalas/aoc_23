package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	u "github.com/Chufretalas/aoc_23/utils"
)

func D10P1() {
	f, _ := os.Open("./inputs/d10.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	m := u.Matrix2D{}
	start := u.Coord2D{L: -1, C: -1}
	current := u.Coord2D{L: -1, C: -1}
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
	guesses := []u.Direction{u.UP, u.DOWN, u.LEFT, u.RIGHT}
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
			case u.UP:
				switch next {
				case "|":
					dir = u.UP
				case "F":
					dir = u.RIGHT
				case "7":
					dir = u.LEFT
				default:
					current = start
					failed = true
				}
			case u.DOWN:
				switch next {
				case "|":
					dir = u.DOWN
				case "L":
					dir = u.RIGHT
				case "J":
					dir = u.LEFT
				default:
					current = start
					failed = true
				}
			case u.LEFT:
				switch next {
				case "-":
					dir = u.LEFT
				case "L":
					dir = u.UP
				case "F":
					dir = u.DOWN
				default:
					current = start
					failed = true
				}
			case u.RIGHT:
				switch next {
				case "-":
					dir = u.RIGHT
				case "7":
					dir = u.DOWN
				case "J":
					dir = u.UP
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

	m := u.Matrix2D{}
	start := u.Coord2D{L: -1, C: -1}
	current := u.Coord2D{L: -1, C: -1}
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

	var painted u.Matrix2D

	// finding the loop
	guesses := []u.Direction{u.UP, u.DOWN, u.LEFT, u.RIGHT}
	foundLoop := false
	for _, dir := range guesses {
		steps := 0
		failed := false
		painted = make(u.Matrix2D, len(m))
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
			case u.UP:
				switch next {
				case "|":
					dir = u.UP
				case "F":
					dir = u.RIGHT
				case "7":
					dir = u.LEFT
				default:
					current = start
					failed = true
				}
			case u.DOWN:
				switch next {
				case "|":
					dir = u.DOWN
				case "L":
					dir = u.RIGHT
				case "J":
					dir = u.LEFT
				default:
					current = start
					failed = true
				}
			case u.LEFT:
				switch next {
				case "-":
					dir = u.LEFT
				case "L":
					dir = u.UP
				case "F":
					dir = u.DOWN
				default:
					current = start
					failed = true
				}
			case u.RIGHT:
				switch next {
				case "-":
					dir = u.RIGHT
				case "7":
					dir = u.DOWN
				case "J":
					dir = u.UP
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
				coord := u.Coord2D{L: l, C: c}
				nexts := []u.Coord2D{coord.Next(u.UP), coord.Next(u.DOWN), coord.Next(u.LEFT), coord.Next(u.RIGHT)}
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
				painted.Set(u.Coord2D{L: l, C: c}, "🟡")
				enclosed++
			}
		}
	}

	fmt.Println(painted)
	fmt.Printf("enclosed: %v\n", enclosed)
}
