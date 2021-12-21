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

type charCheck struct {
	c rune
	v int
}

type synChecker struct {
	syn    map[rune]rune
	scores map[rune]int
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
		scores: map[rune]int{
			')': 3,
			']': 57,
			'}': 1197,
			'>': 25137,
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

func (s synChecker) checkLine(line string) int {
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
				dbg(" -> ERR returning score[%c]: %v!", c, s.scores[c])
				return s.scores[c]
			}
		}
		dbg("Stack: %v", s.stack)
	}

	return 0
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
	checker := NewSynChecker()
	for i := range lines {
		part1 += checker.checkLine(lines[i])
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
