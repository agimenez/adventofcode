package main

import (
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

func parseSeq(s string) []int {
	seq := []int{}
	values := strings.Fields(s)
	for _, v := range values {
		n, _ := strconv.Atoi(v)
		seq = append(seq, n)
	}

	return seq
}

func nextSeq(s []int, stopValue int) ([]int, bool) {
	next := make([]int, len(s)-1)
	last := false

	zeros := 0
	for i := 0; i < len(s)-1; i++ {
		next[i] = s[i+1] - s[i]
		if next[i] == 0 {
			zeros++
		}
	}

	last = (len(next) == zeros)

	return next, last

}

func allSequences(seq []int, stopValue int) [][]int {
	dbg("INITIAL SEQUENCE: %v", seq)
	seqs := [][]int{}
	seqs = append(seqs, seq)

	last := false
	for !last {
		seq, last = nextSeq(seq, stopValue)
		seqs = append(seqs, seq)
		dbg(" -> NEXT: %v (last: %v)", seq, last)
	}

	return seqs
}

func predictNextValues(seqs [][]int) [][]int {
	dbg("PREDICT")
	for i := len(seqs) - 2; i >= 0; i-- {
		dbg("  SEQ: %v", seqs[i])

		nextV := seqs[i][len(seqs[i])-1] + seqs[i+1][len(seqs[i+1])-1]
		seqs[i] = append(seqs[i], nextV)
		dbg("   -> %v", seqs[i])
	}

	return seqs
}

func predictPreviousValues(seqs [][]int) [][]int {
	dbg("PREDICT")
	for i := len(seqs) - 2; i >= 0; i-- {
		dbg("  SEQ: %v", seqs[i])

		nextV := seqs[i][0] - seqs[i+1][0]
		seqs[i] = slices.Insert(seqs[i], 0, nextV)
		dbg("   -> %v", seqs[i])
	}

	return seqs
}

func solve1(s string) []int {
	seq := parseSeq(s)

	seqs := allSequences(seq, 0)
	dbg("ALLSEQS: %v", seqs)
	seqs = predictNextValues(seqs)
	dbg("PREDICTEDSEQS: %v", seqs)

	return seqs[0]
}

func solve2(s string) []int {
	seq := parseSeq(s)

	seqs := allSequences(seq, 0)
	dbg("ALLSEQS: %v", seqs)
	seqs = predictPreviousValues(seqs)
	dbg("PREDICTEDSEQS: %v", seqs)

	return seqs[0]
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
	for _, l := range lines {
		sequence := solve1(l)
		part1 += sequence[len(sequence)-1]
		sequence = solve2(l)
		part2 += sequence[0]
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
