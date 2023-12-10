package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type TYPE int

const (
	HIGHCARD TYPE = iota
	ONEPAIR
	TWOPAIR
	THREEOFAKIND
	FULLHOUSE
	FOUROFAKIND
	FIVEOFAKIND
)

type Hand struct {
	Cards string
	Type  TYPE
}

var c2v_1 = map[string]int{
	"2": 0,
	"3": 1,
	"4": 2,
	"5": 3,
	"6": 4,
	"7": 5,
	"8": 6,
	"9": 7,
	"T": 8,
	"J": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

type Play struct {
	Hand Hand
	Bid  int
}

func getType_1(cards string) TYPE {
	uniques := make([]rune, 0, 5)

	for _, c := range cards {
		if !slices.Contains(uniques, c) {
			uniques = append(uniques, c)
		}
	}

	counts := make([]int, len(uniques))

	for idx, r := range uniques {
		for _, c := range cards {
			if r == c {
				counts[idx]++
			}
		}
	}

	for _, c := range counts {
		if c == 5 {
			return FIVEOFAKIND
		}
	}

	for _, c := range counts {
		if c == 4 {
			return FOUROFAKIND
		}
	}

	if len(uniques) == 2 {
		return FULLHOUSE
	}

	for _, c := range counts {
		if c == 3 {
			return THREEOFAKIND
		}
	}

	twos := 0
	for _, c := range counts {
		if c == 2 {
			twos++
		}
	}

	if twos == 2 {
		return TWOPAIR
	}

	if twos == 1 {
		return ONEPAIR
	}

	return HIGHCARD
}

func Compare_1(p1, p2 Play) int {
	if p1.Hand.Type != p2.Hand.Type {
		return cmp.Compare(p1.Hand.Type, p2.Hand.Type)
	}

	for i := 0; i < 5; i++ {
		res := cmp.Compare(c2v_1[string(p1.Hand.Cards[i])], c2v_1[string(p2.Hand.Cards[i])])
		if res != 0 {
			return res
		}
	}

	panic("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
}

func D7P1() {
	f, _ := os.Open("./inputs/d7.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	plays := make([]Play, 0, 1000)
	for scan.Scan() {
		split := strings.Split(scan.Text(), " ")
		cards := split[0]
		bid, _ := strconv.Atoi(split[1])
		plays = append(plays, Play{Hand: Hand{Cards: cards, Type: getType_1(cards)}, Bid: bid})
	}

	slices.SortFunc(plays, func(p1, p2 Play) int {
		return Compare_1(p1, p2)
	})

	res := 0

	for idx, p := range plays {
		res += p.Bid * (idx + 1)
	}

	fmt.Printf("res: %v\n", res)
}

// --------------------------------------------- PART 2 --------------------------------------------- //

var c2v_2 = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

func getType_2(cards string) TYPE {

	copy := cards

	notJokers := make([]rune, 0, 5)

	for _, c := range cards {
		if !slices.Contains(notJokers, c) && string(c) != "J" {
			notJokers = append(notJokers, c)
		}
	}

	if len(notJokers) != 0 {
		popular := string(notJokers[0])
		popular_count := 0

		for _, u := range notJokers {
			count := 0
			for _, c := range cards {
				if u == c {
					count++
				}
			}
			if count > popular_count {
				popular_count = count
				popular = string(u)
			}
		}
		copy = strings.ReplaceAll(copy, "J", popular)
		// fmt.Printf("cards: %v\n", cards)
		// fmt.Printf("copy: %v\n", copy)
	}

	return getType_1(copy)
}

func Compare_2(p1, p2 Play) int {
	if p1.Hand.Type != p2.Hand.Type {
		return cmp.Compare(p1.Hand.Type, p2.Hand.Type)
	}

	for i := 0; i < 5; i++ {
		res := cmp.Compare(c2v_2[string(p1.Hand.Cards[i])], c2v_2[string(p2.Hand.Cards[i])])
		if res != 0 {
			return res
		}
	}

	panic("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
}

func D7P2() {
	f, _ := os.Open("./inputs/d7.txt")
	defer f.Close()

	scan := bufio.NewScanner(f)

	plays := make([]Play, 0, 1000)
	for scan.Scan() {
		split := strings.Split(scan.Text(), " ")
		cards := split[0]
		bid, _ := strconv.Atoi(split[1])
		plays = append(plays, Play{Hand: Hand{Cards: cards, Type: getType_2(cards)}, Bid: bid})
	}

	slices.SortFunc(plays, func(p1, p2 Play) int {
		return Compare_2(p1, p2)
	})

	res := 0

	for idx, p := range plays {
		res += p.Bid * (idx + 1)
	}

	fmt.Printf("res: %v\n", res)
}
