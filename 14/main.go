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

type Masker map[int]int

func NewMasker(mask string) Masker {
	m := make(map[int]int)
	trans := map[rune]int{
		'1': 1,
		'0': 0,
	}

	for i, c := range mask {
		if c == 'X' {
			continue
		}

		m[(len(mask)-1)-i] = trans[c]

	}

	return m
}

func (m Masker) Apply(v uint64) uint64 {
	dbg("vi = %036b", v)
	for pos, bit := range m {
		dbg("  pos %v, bit %v", pos, bit)
		switch bit {
		case 1:
			v |= (0x01 << pos)
		case 0:
			v &^= (0x01 << pos)
		}
	}
	dbg("vf = %036b\n", v)

	return v
}

type Computer struct {
	mem  map[int]uint64
	mask Masker
}

func NewComputer() *Computer {
	return &Computer{
		mem:  make(map[int]uint64),
		mask: NewMasker(""),
	}
}

func (c *Computer) Write(a int, v uint64) {
	v = c.mask.Apply(v)
	c.mem[a] = v
	dbg("mem[%d] = %036b", a, v)
}

func (c *Computer) SumMem() uint64 {
	var res uint64

	for i := range c.mem {
		res += c.mem[i]
	}

	return res
}
func (c *Computer) Run(lines []string) {
	for i := range lines {
		l := strings.Split(lines[i], " = ")
		op, v := l[0], l[1]
		switch op[:4] {
		case "mask":
			c.mask = NewMasker(v)
		case "mem[":
			addr, err := strconv.Atoi(op[4 : len(op)-1])
			if err != nil {
				panic("Can't parse address")
			}

			val, err := strconv.Atoi(v)
			if err != nil {
				panic("Can't parse value")
			}

			c.Write(addr, uint64(val))
		}
	}

}

func main() {

	part1, part2 := uint64(0), uint64(0)
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	maskString := strings.Split(lines[0], " ")[2]
	dbg("Mask: %#v", maskString)
	c := NewComputer()
	c.Run(lines)
	dbg("computer: %#v", c)
	part1 = c.SumMem()

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
