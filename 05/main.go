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
	var seen [1024]bool
	for s.Scan() {
		l := s.Text()
		r := getSeatID(l)
		seen[r] = true
		if r > max {
			max = r
		}

	}

	var mySeat int
	for i := 1; i < len(seen)-1; i++ {
		dbg("%v: %v", i, seen[i])
		if !seen[i] && seen[i-1] {
			mySeat = i
			break
		}
	}

	log.Printf("Part 1: %v\n", max)
	log.Printf("Part 2: %v\n", mySeat)

}
