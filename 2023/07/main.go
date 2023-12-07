package main

import (
	"cmp"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	debug bool
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
}

const (
	highCard = iota
	onePair
	twoPair
	threeKind
	fullHouse
	fourKind
	fiveKind
)

type hand struct {
	cards    [5]uint
	bid      int
	handType int
}

func cardToValue(c rune) uint {
	switch c {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	}

	return (uint(c - '0'))
}

func getHandType(h hand) int {
	reps := map[uint]int{}

	maxCount := 0
	for _, c := range h.cards {
		reps[c]++
		if reps[c] > maxCount {
			maxCount = reps[c]
		}
	}

	threeSeen := false
	pairSeen := false
	for _, count := range reps {
		if count == 5 {
			return fiveKind
		}

		if count == 4 {
			return fourKind
		}

		if count == 3 {
			// This could be either fullHouse or Three of a Kind
			threeSeen = true
		}

		if count == 2 {
			// already seen a double: must be two pair
			if pairSeen {
				return twoPair
			}

			pairSeen = true

		}
	}

	if threeSeen {
		if pairSeen {
			return fullHouse
		}

		return threeKind
	}

	if pairSeen {
		return onePair
	}

	return highCard
}

func parseHand(s string) hand {
	h := hand{}
	fields := strings.Fields(s)
	h.bid, _ = strconv.Atoi(fields[1])
	for i, c := range fields[0] {
		h.cards[i] = cardToValue(c)
	}

	h.handType = getHandType(h)
	dbg("Hand: %s, type: %v", s, h.handType)

	return h
}

func cmpHands(a, b hand) int {
	ret := cmp.Compare(a.handType, b.handType)
	if ret != 0 {
		return ret
	}

	return slices.Compare(a.cards[:], b.cards[:])

}

func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)
	hands := []hand{}
	for _, l := range lines {
		h := parseHand(l)
		hands = append(hands, h)
	}
	slices.SortFunc(hands, cmpHands)
	for rank, hand := range hands {
		dbg("rank %d, hand %v", rank, hand)
		part1 += (rank + 1) * hand.bid
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
