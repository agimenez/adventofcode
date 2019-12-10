package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strings"
)

var (
	debug int
)

func dbg(level int, fmt string, v ...interface{}) {
	if debug >= level {
		log.Printf(fmt, v...)
	}
}

func init() {
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.Parse()
}

func main() {

	var in string
	fmt.Scan(&in)
	wide := 25
	tall := 6

	minz := math.MaxInt32
	res := 0
	for i, start, end := 0, 0, wide*tall; start < len(in); i, start, end = i+1, end, end+(wide*tall) {
		dbg(1, "%d: start = %d, end = %d", i, start, end)
		str := in[start:end]
		z := strings.Count(str, "0")
		dbg(2, "str = %s, z = %d (%d)", str, z, minz)
		if z < minz {
			o := strings.Count(str, "1")
			t := strings.Count(str, "2")
			dbg(2, "new max! o = %d, t = %d", o, t)

			minz = z
			res = o * t
		}

	}

	log.Printf("Max output: %v", res)

}
