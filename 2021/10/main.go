package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sort"
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

type score struct {
	corrupt  int
	complete int
}

type synChecker struct {
	syn    map[rune]rune
	scores map[rune]score
	stack  []rune
}

func NewSynChecker() synChecker {
	return synChecker{
		syn: map[rune]rune{
			'(': ')',
			'[': ']',
			'{': '}',
			'<': '>',
		},
		scores: map[rune]score{
			')': {corrupt: 3, complete: 1},
			']': {corrupt: 57, complete: 2},
			'}': {corrupt: 1197, complete: 3},
			'>': {corrupt: 25137, complete: 4},
		},
		stack: make([]rune, 0),
	}
}

func (s *synChecker) Push(e rune) {
	s.stack = append(s.stack, e)
}

func (s *synChecker) Pop() rune {
	var ret rune
	ret, s.stack = s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]

	return ret
}

func (s synChecker) Empty() bool {
	return len(s.stack) == 0
}

func (s synChecker) Top() rune {
	return s.stack[len(s.stack)-1]
}

func (s *synChecker) checkLine(line string) int {
	dbg("Check line %s", line)
	for _, c := range line {
		dbg(" -> %c", c)
		if exp, ok := s.syn[c]; ok {
			dbg(" -> Push %c", exp)
			s.Push(exp)
		} else {
			expected := s.Pop()
			dbg(" -> Expected: %c", expected)
			if c != expected {
				dbg(" -> ERR returning score[%c]: %v!", c, s.scores[c].corrupt)
				return s.scores[c].corrupt
			}
		}
		dbg("Stack: %v", string(s.stack))
	}

	return 0
}

func (s synChecker) completeLine(line string) int {
	score := 0
	for i := len(s.stack) - 1; i >= 0; i-- {
		ch := s.stack[i]
		score *= 5
		score += s.scores[ch].complete
	}

	return score
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
	completeScores := make([]int, 0)
	for i := range lines {
		checker := NewSynChecker()
		score := checker.checkLine(lines[i])
		part1 += score
		if score == 0 {
			completeScores = append(completeScores, checker.completeLine(lines[i]))
		}
	}
	sort.Sort(sort.IntSlice(completeScores))
	part2 = completeScores[len(completeScores)/2]

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
