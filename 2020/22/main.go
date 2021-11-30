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
	flag.Parse()
}

func Play(players [][]int) []int {

	for {
		var c1, c2 int
		dbg("Player %d's deck: %v", 1, players[0])
		c1, players[0] = players[0][0], players[0][1:]
		dbg("Player %d's deck: %v", 2, players[1])
		c2, players[1] = players[1][0], players[1][1:]

		dbg("Player %d plays: %d", 1, c1)
		dbg("Player %d plays: %d", 2, c2)
		if c1 > c2 {
			dbg("Player 1 wins the round!")
			players[0] = append(players[0], c1, c2)
		} else {
			dbg("Player 2 wins the round!")
			players[1] = append(players[1], c2, c1)
		}
		//dbg("Player %d's deck: %v", 1, players[0])
		//dbg("Player %d's deck: %v", 2, players[1])

		if len(players[0]) == 0 {
			return players[1]
		} else if len(players[1]) == 0 {
			return players[0]
		}

	}

	return nil
}

func Score(player []int) int {
	score := 0
	for i, v := range player {
		score += v * (len(player) - i)
		dbg("%v * %v", v, len(player)-i)
	}

	return score
}

func main() {

	part1, part2 := 0, 0
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(in), "\n")

	players := [][]int{}
	var p []int
	for _, l := range lines {
		if l == "" {
			players = append(players, p)
			continue
		}

		if l[len(l)-1] == ':' {
			p = []int{}
			continue
		}

		num, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalf("Couldn't parse '%s'", l)
		}
		p = append(p, num)
	}
	dbg("players: %#v", players)
	part1 = Score(Play(players))

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
