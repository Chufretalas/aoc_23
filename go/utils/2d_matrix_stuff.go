package utils

import "fmt"

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
