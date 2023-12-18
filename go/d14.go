package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	u "github.com/Chufretalas/aoc_23/utils"
)

func D14P1() {
	f, _ := os.Open("./inputs/d14.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	m := u.Matrix2D{}
	for scan.Scan() {
		m = append(m, strings.Split(scan.Text(), ""))
	}

	for j := range m[0] {
		for i := range m {
			initial := u.Coord2D{L: i, C: j}
			if m.Get(initial) == "O" {
				final := u.Coord2D{L: i, C: j}
				for {
					next := final.Next(u.UP)
					if next.L < 0 || m.Get(next) == "#" || m.Get(next) == "O" {
						break
					}
					final = next
				}
				m.Set(initial, ".")
				m.Set(final, "O")
			}
		}
	}

	sum := 0

	for i := range m {
		for j := range m[i] {
			if m[i][j] == "O" {
				sum += len(m) - i
			}
		}
	}

	fmt.Printf("sum: %v\n", sum)
}

// --------------------------------------------- PART 2 --------------------------------------------- //

func moveBolder(m *u.Matrix2D, initial u.Coord2D, direction u.Direction) {
	if m.Get(initial) == "O" {
		final := initial
		for {
			next := final.Next(direction)
			if next.L < 0 || next.L >= len(*m) || next.C < 0 || next.C >= len((*m)[0]) || m.Get(next) == "#" || m.Get(next) == "O" {
				break
			}
			final = next
		}
		m.Set(initial, ".")
		m.Set(final, "O")
	}
}

// the waitgroups did not work
func D14P2() {
	f, _ := os.Open("./inputs/d14.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	m := u.Matrix2D{}
	for scan.Scan() {
		m = append(m, strings.Split(scan.Text(), ""))
	}

	var wg sync.WaitGroup
	for k := 0; k < 100; k++ {
		for _, direction := range []u.Direction{u.UP, u.LEFT, u.DOWN, u.RIGHT} {
			// fmt.Println(m)
			// tilt
			switch direction {
			case u.UP:
				for j := range m[0] {
					wg.Add(1)
					go func() {
						defer wg.Done()
						for i := range m {
							moveBolder(&m, u.Coord2D{L: i, C: j}, direction)
						}
					}()
				}
				wg.Wait()
			case u.LEFT:
				for i := range m {
					for j := range m[i] {
						moveBolder(&m, u.Coord2D{L: i, C: j}, direction)
					}
				}
			case u.DOWN:
				for j := range m[0] {
					for i := len(m) - 1; i >= 0; i-- {
						moveBolder(&m, u.Coord2D{L: i, C: j}, direction)
					}
				}
			case u.RIGHT:
				for i := range m {
					for j := len(m[0]) - 1; j >= 0; j-- {
						moveBolder(&m, u.Coord2D{L: i, C: j}, direction)
					}
				}

			}

		}
	}
	sum := 0

	for i := range m {
		for j := range m[i] {
			if m[i][j] == "O" {
				sum += len(m) - i
			}
		}
	}

	// fmt.Print(m)

	fmt.Printf("sum: %v\n", sum)
}
