package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
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

var scores = map[string]int{
	// Rock
	"A": 1,
	"X": 1,

	// Paper
	"B": 2,
	"Y": 2,

	// Scissors
	"C": 3,
	"Z": 3,
}

type pair struct {
	a, b string
}

var playScore = map[pair]int{
	// Draws
	{"A", "X"}: 3,
	{"B", "Y"}: 3,
	{"C", "Z"}: 3,

	// Wins
	{"A", "Y"}: 6,
	{"B", "Z"}: 6,
	{"C", "X"}: 6,
}

var derive = map[string]string{
	// Rock
	"A X": "A Z", // Need to lose (paper)
	"A Y": "A X", // Need to draw (rock)
	"A Z": "A Y", // Need to win (scisors)

	// Paper
	"B X": "B X", // Need to lose (rock)
	"B Y": "B Y", // Need to draw (paper)
	"B Z": "B Z", // Need to win (scissors)

	// Scissors
	"C X": "C Y", // Need to lose (paper)
	"C Y": "C Z", // Need to draw (scissors)
	"C Z": "C X", // Need to win (rock)
}

func play(in []string) int {
	//their := in[0]
	//mine := in[1]
	return playScore[pair{in[0], in[1]}]
}
func roundScore(in string) int {
	choices := strings.Split(in, " ")

	shapeScore := scores[choices[1]]
	outcomeScore := play(choices)
	dbg("Shapescore: %v, outcomescore: %v", shapeScore, outcomeScore)

	return shapeScore + outcomeScore
}

func followStrategy(in []string) int {
	total := 0
	for _, round := range in {
		total += roundScore(round)
	}

	return total
}

func deriveFromStrategy(in []string) int {
	total := 0
	for _, round := range in {
		dbg("Round: %v", round)
		play := derive[round]
		dbg("Derived play: %v", play)
		dbg("round score: %v", roundScore(play))
		total += roundScore(play)
	}

	return total
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	dbg("lines: %#v", lines)

	part1 = followStrategy(lines)
	part2 = deriveFromStrategy(lines)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
