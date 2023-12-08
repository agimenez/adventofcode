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
	cards [5]rune
	joker rune
	bid   int
}

func (h hand) cardValue(i int) uint {
	if h.cards[i] == h.joker {
		return 1
	}

	switch h.cards[i] {
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

	return (uint(h.cards[i] - '0'))
}

func (h hand) getHandType() int {
	reps := map[rune]int{}

	maxCount := 0
	var maxCard rune
	for _, c := range h.cards {
		reps[c]++
		if c != h.joker && reps[c] > maxCount {
			maxCount = reps[c]
			maxCard = c
		}
	}

	if maxCard != h.joker {
		reps[maxCard] += reps[h.joker]
		reps[h.joker] = 0
	}

	threeSeen := false
	pairSeen := false
	for _, count := range reps {
		dbg(" -> Count: %v", count)
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

func parseHand(s string, joker rune) hand {
	h := hand{}
	fields := strings.Fields(s)
	h.bid, _ = strconv.Atoi(fields[1])
	for i, c := range fields[0] {
		h.cards[i] = c
	}

	h.joker = joker
	dbg("Hand: %s", s)

	return h
}

func cmpHands(a, b hand) int {
	ret := cmp.Compare(a.getHandType(), b.getHandType())
	if ret != 0 {
		return ret
	}

	for i := 0; i < len(a.cards); i++ {
		ret = cmp.Compare(a.cardValue(i), b.cardValue(i))
		if ret != 0 {
			return ret
		}
	}

	return 0

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
		h := parseHand(l, '0')
		hands = append(hands, h)
	}
	slices.SortFunc(hands, cmpHands)
	for rank, hand := range hands {
		dbg("rank %d, hand %v", rank, hand)
		part1 += (rank + 1) * hand.bid
	}

	hands = []hand{}
	for _, l := range lines {
		h := parseHand(l, 'J')
		hands = append(hands, h)
	}
	slices.SortFunc(hands, cmpHands)
	for rank, hand := range hands {
		dbg("rank %d, hand %s", rank, string(hand.cards[:]))
		part2 += (rank + 1) * hand.bid
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
