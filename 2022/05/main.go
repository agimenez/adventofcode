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
}

type Stack []byte

func (s Stack) Push(b byte) Stack {
	s = append(s, b)

	return s
}

func (s Stack) Top() byte {
	return s[len(s)-1]
}

func (s Stack) Pop() (Stack, byte) {
	if len(s) == 0 {
		return Stack{}, 0
	}

	return s[:len(s)-1], s[len(s)-1]
}

func (s Stack) Insert(b byte) Stack {
	s = append([]byte{b}, s...)

	return s
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
	dbg("lines: %#v", lines)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
