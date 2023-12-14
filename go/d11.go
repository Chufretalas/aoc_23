package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"

	u "github.com/Chufretalas/aoc_23/utils"
)

func D11P1() {
	f, _ := os.Open("./inputs/d11.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	// read input duplicating lines
	original := make([][]string, 0)
	for scan.Scan() {
		split := strings.Split(scan.Text(), "")
		empty := true
		for _, s := range split {
			if s != "." {
				empty = false
				break
			}
		}
		if empty {
			original = append(original, split[:])
		}
		original = append(original, split[:])
	}

	// duplicate columns
	m := make(u.Matrix2D, len(original))
	for c := range original[0] {
		empty := true
		for l := range original {
			if original[l][c] != "." {
				empty = false
				break
			}
		}
		for l := range original {
			if empty {
				m[l] = append(m[l], original[l][c])
			}
			m[l] = append(m[l], original[l][c])
		}
	}

	// find galaxies
	gals := make([]u.Coord2D, 0)
	for l := range m {
		for c := range m[l] {
			coord := u.Coord2D{L: l, C: c}
			if m.Get(coord) == "#" {
				gals = append(gals, coord)
			}
		}
	}

	sum := 0
	// calculate distaces
	for i1 := range gals {
		for i2 := i1 + 1; i2 < len(gals); i2++ {
			sum += int(math.Abs(float64(gals[i1].L-gals[i2].L)) + math.Abs(float64(gals[i1].C-gals[i2].C)))
		}
	}

	fmt.Printf("sum: %v\n", sum)
}

// --------------------------------------------- PART 2 --------------------------------------------- //

// in am going to indicate the lines and columns that are worth more
func D11P2() {
	f, _ := os.Open("./inputs/d11.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	bigLines := make([]int, 0)
	bigCols := make([]int, 0)
	m := make(u.Matrix2D, 0)
	lineIdx := 0
	for scan.Scan() {
		split := strings.Split(scan.Text(), "")
		empty := true
		for _, s := range split {
			if s != "." {
				empty = false
				break
			}
		}
		if empty {
			bigLines = append(bigLines, lineIdx)
		}
		m = append(m, split)
		lineIdx++
	}

	// duplicate columns
	for c := range m[0] {
		empty := true
		for l := range m {
			if m[l][c] != "." {
				empty = false
				break
			}
		}
		if empty {
			bigCols = append(bigCols, c)
		}
	}

	// find galaxies
	gals := make([]u.Coord2D, 0)
	for l := range m {
		for c := range m[l] {
			coord := u.Coord2D{L: l, C: c}
			if m.Get(coord) == "#" {
				gals = append(gals, coord)
			}
		}
	}

	sum := 0
	// calculate distaces
	for i1 := range gals {
		for i2 := i1 + 1; i2 < len(gals); i2++ {
			smaller := gals[i1].L
			bigger := gals[i2].L
			if smaller > bigger {
				temp := smaller
				smaller = bigger
				bigger = temp
			}
			for i := smaller; i < bigger; i++ {
				if slices.Contains(bigLines, i) {
					sum += 1000000
				} else {
					sum++
				}
			}

			smaller = gals[i1].C
			bigger = gals[i2].C
			if smaller > bigger {
				temp := smaller
				smaller = bigger
				bigger = temp
			}
			for i := smaller; i < bigger; i++ {
				if slices.Contains(bigCols, i) {
					sum += 1000000
				} else {
					sum++
				}
			}
		}
	}

	fmt.Printf("sum: %v\n", sum)
}
