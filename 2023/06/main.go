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

func atoi(s string) int {
	n, _ := strconv.Atoi(s)

	return n
}

type race struct {
	time     int
	distance int
}

type competition []race

func newCompetition(s []string) competition {
	c := competition{}

	// Time
	races := strings.Fields(strings.Split(s[0], ":")[1])
	for _, r := range races {
		c = append(c, race{time: atoi(r)})
	}

	// Distance
	races = strings.Fields(strings.Split(s[1], ":")[1])
	for i, r := range races {
		c[i].distance = atoi(r)
	}

	return c
}

func newCompetitionBig(s []string) competition {
	c := make(competition, 1)

	races := strings.Fields(strings.Split(s[0], ":")[1])
	val := ""
	for _, r := range races {
		val += r
	}
	c[0].time = atoi(val)

	// Distance
	races = strings.Fields(strings.Split(s[1], ":")[1])
	val = ""
	for _, r := range races {
		val += r
	}
	c[0].distance = atoi(val)

	return c
}

func (c competition) race(r int) []int {
	race := c[r]
	wins := []int{}

	for time := 1; time < race.time; time++ {
		dist := time * (race.time - time)
		if dist > race.distance {
			wins = append(wins, time)
		}
	}

	return wins

}

func main() {
	flag.Parse()

	part1, part2 := 1, 1
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)
	comp := newCompetition(lines)
	dbg("%v", comp)
	for i := range comp {
		ways := comp.race(i)
		part1 *= len(ways)
	}

	comp2 := newCompetitionBig(lines)
	dbg("%v", comp2)
	part2 = len(comp2.race(0))

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
