package main

import (
	"bytes"
	"flag"
	"fmt"
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

func (s Stack) PushN(b []byte) Stack {
	return s
}

func (s Stack) Push(b byte) Stack {
	s = append(s, b)

	return s
}

func (s Stack) Top() byte {
	return s[len(s)-1]
}

func (s Stack) PopN(count int) (Stack, []byte) {
	return s, []byte{}
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

type Crane []Stack

func (c Crane) InsertCrate(stack int, crane byte) Crane {
	if len(c) <= stack {
		c2 := make(Crane, stack+1)
		copy(c2, c)
		c = c2
	}

	c[stack] = c[stack].Insert(crane)

	return c
}

func (c Crane) MoveCrates(count, from, to int) Crane {
	for ; count > 0; count-- {
		s, crate := c[from].Pop()
		c[from] = s
		c[to] = c[to].Push(crate)
	}

	return c
}

func (c Crane) TopCrates() string {
	var b bytes.Buffer
	for _, s := range c {
		b.WriteByte(s.Top())
	}

	return b.String()
}

func main() {
	flag.Parse()

	part1, part2 := "", ""
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	dbg("lines: %#v", lines)

	crane := Crane{}
	crane9k1 := Crane{}
	dbg("crane: %#v", crane)
	for _, line := range lines {
		dbg("line: %v", line)
		if strings.Contains(line, "[") {
			for i := 1; i < len(line); i += 4 {
				if line[i] != ' ' {
					crane = crane.InsertCrate(i/4, line[i])
					crane9k1 = crane9k1.InsertCrate(i/4, line[i])
				}
			}
			dbg("Crane: %q", crane)
		}

		if strings.Contains(line, "move") {
			var count, src, dst int
			fmt.Sscanf(line, "move %d from %d to %d", &count, &src, &dst)
			crane = crane.MoveCrates(count, src-1, dst-1)
			dbg("Crane: %q", crane)
		}
	}
	part1 = crane.TopCrates()
	part2 = crane9k1.TopCrates()

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
