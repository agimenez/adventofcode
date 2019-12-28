package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var (
	debug int
)

func dbg(level int, fmt string, v ...interface{}) {
	if debug >= level {
		log.Printf(fmt+"\n", v...)
	}
}

func init() {
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.Parse()
}

func pattern(position int) []int {
	var basePattern = []int{0, 1, 0, -1}
	p := []int{}

	for _, n := range basePattern {
		for i := position; i > 0; i-- {
			p = append(p, n)
		}
	}

	return p
}

func FFTnthDigit(in string, pos int) rune {
	total := 0
	p := pattern(pos)
	dbg(1, "Pattern: %v", p)
	for i, d := range in {
		num := int(d - '0')
		dbg(2, " %d * %d", num, p[(i+1)%len(p)])
		total += num * p[(i+1)%len(p)]
	}

	dbg(1, "Pos %d, Total: %d", pos, total)
	t := fmt.Sprintf("%d", total)
	return rune(t[len(t)-1])
}

func FFTPhase(in string) string {
	var b strings.Builder
	for i := range in {
		b.WriteRune(FFTnthDigit(in, i+1))
	}

	return b.String()
}

func FFT(in string, phases int) string {
	var digits string
	for ; phases > 0; phases-- {
		digits = FFTPhase(in)
		in = digits
	}

	return digits
}

func main() {

	var in string
	fmt.Scan(&in)

	// test
	in = "12345678"
	result := FFT(in, 4)

	fmt.Printf("Part one: %#v\n", result)

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func mod(a, b int) int {
	return (a%b + b) % b
}
