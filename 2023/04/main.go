package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
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

func listFromString(s string) []int {
	l := []int{}

	for _, num := range strings.Split(s, " ") {
		n, err := strconv.Atoi(num)
		if err != nil {
			continue
		}

		l = append(l, n)
	}

	return l
}
func mapList(s string) map[int]bool {
	nums := listFromString(s)
	m := map[int]bool{}
	for _, n := range nums {
		m[n] = true
	}

	return m
}

func getCardScore(card string) (int, int) {
	numbers := strings.Split(card, ": ")[1]
	parts := strings.Split(numbers, " | ")
	winners := mapList(parts[0])
	played := listFromString(parts[1])

	wins := 0
	points := 0
	//dbg("winners: %v", winners)
	for _, n := range played {
		//dbg("  -> Playing %v", n)
		if _, ok := winners[n]; ok {
			wins++
			//dbg("  -> WINNER! -> %v", wins)
		}
	}

	return (1 << uint(wins-1)), wins
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
	current := 1
	cards := map[int]int{}

	for _, card := range lines {
		cards[current] = 1 + cards[current]
		points, wins := getCardScore(card)
		part1 += points

		dbg("Card %v has %v wins", current, wins)
		for i := 1; i <= wins; i++ {
			cards[current+i] += cards[current]
			dbg("  -> card %v -> %v", current+i, cards[current+i])
		}

		dbg("Current: %v, cards: %v", current, cards)
		current++
	}

	for _, v := range cards {
		part2 += v
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
