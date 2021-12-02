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

type submarine struct {
	position int
	depth    int
	aim      int
}

func (s *submarine) command(cmd string) {
	parts := strings.Split(cmd, " ")
	val, _ := strconv.Atoi(parts[1])
	switch parts[0] {
	case "forward":
		s.forward(val)
	case "up":
		s.up(val)
	case "down":
		s.down(val)
	}
}

func (s *submarine) forward(val int) {
	s.position += val
	s.depth += s.aim * val
}

func (s *submarine) up(val int) {
	s.aim -= val
}

func (s *submarine) down(val int) {
	s.aim += val
}

func (s *submarine) positions() (int, int) {
	return s.position, s.depth
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

	s := submarine{}
	for i := range lines {
		s.command(lines[i])
	}

	pos, depth := s.positions()
	part1 = pos * depth

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
