package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	u "github.com/Chufretalas/aoc_23/utils"
)

type Beam struct {
	c u.Coord2D
	d u.Direction
}

func countEnergized(contraption *u.Matrix2D, startingPos u.Coord2D, startingDirec u.Direction) int {

	directionsMap := make(u.Matrix2D, 0) // map with directions previously traveled by a beam
	energyMap := make(u.Matrix2D, 0)     // map with add spots where a beam has visited

	for i := 0; i < len(*contraption); i++ {
		directionsMap = append(directionsMap, make([]string, len((*contraption))))
		energyMap = append(energyMap, make([]string, len((*contraption))))
	}

	beams := []Beam{{c: startingPos, d: startingDirec}}

	for len(beams) != 0 {
		obj := contraption.Get(beams[0].c)

		// beam got out of bounds
		if obj == "" {
			beams = slices.Delete(beams, 0, 1)
			continue
		}

		// a beam is redoing a path
		if strings.Contains(directionsMap.Get(beams[0].c), fmt.Sprintf("%v", beams[0].d)) {
			beams = slices.Delete(beams, 0, 1)
			continue
		}

		energyMap.Set(beams[0].c, "#")
		directionsMap.Set(beams[0].c, fmt.Sprintf("%v", beams[0].d))

		switch obj {
		case `/`:
			switch beams[0].d {
			case u.UP:
				beams[0].d = u.RIGHT
				beams[0].c.Move(u.RIGHT)
			case u.DOWN:
				beams[0].d = u.LEFT
				beams[0].c.Move(u.LEFT)
			case u.LEFT:
				beams[0].d = u.DOWN
				beams[0].c.Move(u.DOWN)
			case u.RIGHT:
				beams[0].d = u.UP
				beams[0].c.Move(u.UP)
			}
		case `\`:
			switch beams[0].d {
			case u.UP:
				beams[0].d = u.LEFT
				beams[0].c.Move(u.LEFT)
			case u.DOWN:
				beams[0].d = u.RIGHT
				beams[0].c.Move(u.RIGHT)
			case u.LEFT:
				beams[0].d = u.UP
				beams[0].c.Move(u.UP)
			case u.RIGHT:
				beams[0].d = u.DOWN
				beams[0].c.Move(u.DOWN)
			}
		case "-":
			if beams[0].d == u.UP || beams[0].d == u.DOWN {
				beams = append(beams, Beam{c: beams[0].c, d: u.LEFT}, Beam{c: beams[0].c, d: u.RIGHT})
				beams = slices.Delete(beams, 0, 1)
			} else {
				beams[0].c.Move(beams[0].d)
			}
		case "|":
			if beams[0].d == u.LEFT || beams[0].d == u.RIGHT {
				beams = append(beams, Beam{c: beams[0].c, d: u.UP}, Beam{c: beams[0].c, d: u.DOWN})
				beams = slices.Delete(beams, 0, 1)
			} else {
				beams[0].c.Move(beams[0].d)
			}
		default:
			directionsMap.Set(beams[0].c, fmt.Sprintf("%v%v", directionsMap.Get(beams[0].c), (beams[0].d)))
			beams[0].c.Move(beams[0].d)
		}
	}

	result := 0
	for _, l := range energyMap {
		for _, v := range l {
			if v == "#" {
				result++
			}
		}
	}
	return result
}

func D16P1() {
	f, _ := os.Open("./inputs/d16_test.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	contraption := make(u.Matrix2D, 0) // immutable original map

	for scan.Scan() {
		split := strings.Split(scan.Text(), "")
		contraption = append(contraption, split)
	}

	fmt.Println(countEnergized(&contraption, u.Coord2D{L: 0, C: 0}, u.RIGHT))
}

func D16P2() {
	f, _ := os.Open("./inputs/d16.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	contraption := make(u.Matrix2D, 0) // immutable original map

	for scan.Scan() {
		split := strings.Split(scan.Text(), "")
		contraption = append(contraption, split)
	}

	result := 0

	for i := 0; i < len(contraption[0]); i++ {
		if new := countEnergized(&contraption, u.Coord2D{L: 0, C: i}, u.DOWN); new > result {
			result = new
		}
		if new := countEnergized(&contraption, u.Coord2D{L: len(contraption) - 1, C: i}, u.UP); new > result {
			result = new
		}
	}

	for i := 0; i < len(contraption); i++ {
		if new := countEnergized(&contraption, u.Coord2D{L: i, C: 0}, u.RIGHT); new > result {
			result = new
		}
		if new := countEnergized(&contraption, u.Coord2D{L: i, C: len(contraption[0]) - 1}, u.LEFT); new > result {
			result = new
		}
	}

	fmt.Printf("result: %v\n", result)
}
