package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	debug = false
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

func translate(s string) int {
	b := strings.Map(func(r rune) rune {
		switch r {
		case 'F', 'L':
			return '0'
		case 'B', 'R':
			return '1'
		}
		return -1
	}, s)

	dbg("b = %v", b)

	i, err := strconv.ParseInt(b, 2, 8)
	if err != nil {
		panic("malformed input")
	}

	return int(i)
}

func getSeatID(s string) int {
	r := translate(s[:7])
	c := translate(s[7:])
	dbg("Row: %v, Col: %v", r, c)

	return r*8 + c
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	max := 0
	for s.Scan() {
		l := s.Text()
		r := getSeatID(l)
		if r > max {
			max = r
		}

	}

	log.Printf("Part 1: %v\n", max)
	log.Printf("Part 2: \n")

}
