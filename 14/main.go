package main

import (
	"flag"
	"fmt"
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

type Masker interface {
	Apply(v uint64) interface{}
}

type ValueMask map[int]int

func NewValueMasker(mask string) ValueMask {
	m := make(ValueMask)
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

func (m ValueMask) Apply(v uint64) interface{} {
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

type AddressMask struct {
	mask              string
	floating          uint64
	floatingPositions []int
}

func (m AddressMask) String() string {
	return fmt.Sprintf("\nm  = %s\nf  = %d\nfm = %#v\n\n", m.mask, m.floating, m.floatingPositions)
}

func NewAddressMasker(mask string) AddressMask {
	m := AddressMask{
		mask:     mask,
		floating: (1 << strings.Count(mask, "X")),
	}
	for i, bit := range m.mask {
		if bit == 'X' {
			m.floatingPositions = append([]int{len(m.mask) - i - 1}, m.floatingPositions...)
		}
	}

	dbg("AddressMasker: %v", m)
	return m
}

func (m AddressMask) Apply(a uint64) interface{} {
	arr := []uint64{}

	dbg("")
	dbg("m  = %036s", m.mask)
	dbg("vi = %036b", a)
	// first, apply basic 0/1 transformations
	for i, bit := range m.mask {
		pos := len(m.mask) - i - 1
		switch bit {
		case '1':
			a |= (0x01 << pos)
		}
	}
	dbg("v1 = %036b", a)
	dbg("")

	//Resolve the Xs
	for f := uint64(0); f < m.floating; f++ {
		tmp := a
		dbg("f  = %036b (floating #%d)", f, f)
		dbg("n  = %036b", tmp)
		maskSet := uint64(0)
		maskClear := ^maskSet

		for bit, pos := range m.floatingPositions {
			fval := f & (1 << bit)
			dbg("fv = %036b", fval)
			if fval > 0 {
				maskSet |= (1 << pos)
			} else {
				maskClear &= ^(1 << pos)
			}
		}
		dbg("ms = %036b", maskSet)
		dbg("mc = %036b", maskClear)
		tmp |= maskSet
		tmp &= maskClear
		dbg("t  = %036b", tmp)
		dbg("")
		arr = append(arr, tmp)
	}
	dbg("%v", arr)

	return arr
}

type Computer struct {
	mem  map[uint64]uint64
	mask Masker
}

func NewComputer(m Masker) *Computer {
	return &Computer{
		mem:  make(map[uint64]uint64),
		mask: m,
	}
}

func (c *Computer) Write(a uint64, v uint64) {
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
func (c *Computer) RunMaskValues(lines []string) {
	for i := range lines {
		l := strings.Split(lines[i], " = ")
		op, v := l[0], l[1]
		switch op[:4] {
		case "mask":
			c.mask = NewValueMasker(v)
		case "mem[":
			addr, err := strconv.Atoi(op[4 : len(op)-1])
			if err != nil {
				panic("Can't parse address")
			}

			val, err := strconv.Atoi(v)
			if err != nil {
				panic("Can't parse value")
			}

			v64 := c.mask.Apply(uint64(val)).(uint64)
			c.Write(uint64(addr), v64)
		}
	}

}

func (c *Computer) RunMaskAddresses(lines []string) {
	for i := range lines {
		l := strings.Split(lines[i], " = ")
		op, v := l[0], l[1]
		switch op[:4] {
		case "mask":
			c.mask = NewAddressMasker(v)
		case "mem[":
			addr, err := strconv.Atoi(op[4 : len(op)-1])
			if err != nil {
				panic("Can't parse address")
			}

			val, err := strconv.Atoi(v)
			if err != nil {
				panic("Can't parse value")
			}

			ad := c.mask.Apply(uint64(addr)).([]uint64)
			for i := range ad {
				c.Write(ad[i], uint64(val))
			}
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
	c := NewComputer(NewValueMasker(""))
	c.RunMaskValues(lines)
	part1 = c.SumMem()
	log.Printf("Part 1: %v\n", part1)

	c = NewComputer(NewAddressMasker(""))
	c.RunMaskAddresses(lines)
	part2 = c.SumMem()

	log.Printf("Part 2: %v\n", part2)

}
