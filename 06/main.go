package main

import (
	"bufio"
	"log"
	"os"
)

const (
	debug = false
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

type questions map[rune]struct{}

func main() {

	s := bufio.NewScanner(os.Stdin)
	list := []questions{}

	q := questions{}
	for s.Scan() {
		l := s.Text()
		dbg("Line: %v\n", l)
		if l == "" {
			list = append(list, q)
			q = questions{}
			continue
		}

		for _, question := range l {
			q[question] = struct{}{}
		}

		dbg("List: %v\n", list)

	}
	list = append(list, q)

	part1, part2 := 0, 0
	for _, q := range list {
		part1 += len(q)
	}

	log.Printf("Part 1: %v", part1)
	log.Printf("Part 2: %v", part2)

}
