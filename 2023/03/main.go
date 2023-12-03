package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/agimenez/adventofcode/utils"
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

type symbol struct {
	coord utils.Point
	val   rune
}

func isSymbol(c rune) bool {
	return !(unicode.IsDigit(c) || c == '.')
}

func adjSymbol(l []string, x int, y int) bool {

	//dbg("==== Check adjacent of {%v, %v} = %c", x, y, l[y][x])
	for dx := x - 1; dx <= x+1; dx++ {
		for dy := y - 1; dy <= y+1; dy++ {
			if (dx == x && dy == y) || dx < 0 || dy < 0 || dx >= len(l[y]) || dy >= len(l) {
				continue
			}
			//dbg("Checking {%v,%v} = %c", dx, dy, l[dy][dx])

			if isSymbol(rune(l[dy][dx])) {
				//dbg(" -> YES")
				return true
			}
		}
	}

	return false
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
	//dbg("lines: %#v", lines)
	for y, l := range lines {
		partNo := ""
		isPart := false
		for x, c := range l {
			if unicode.IsDigit(c) {
				partNo += string(c)

				if !isPart {
					isPart = adjSymbol(lines, x, y)
				}
			}

			// finished number, or end of line
			if !unicode.IsDigit(c) || x >= len(l)-1 {
				if isPart {
					dbg("Adding %v,%v = %v", x, y, partNo)
					v, _ := strconv.Atoi(partNo)
					part1 += v
				}
				partNo = ""
				isPart = false
			}

		}
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
