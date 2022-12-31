package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	. "github.com/agimenez/adventofcode/utils"
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

func findVisible(f [][]int) map[Point]bool {
	out := make(map[Point]bool)

	var maxHeight int
	// Check down/up
	for x := 0; x < len(f); x++ {
		maxHeight = -1
		for y := 0; y < len(f[x]); y++ {
			if f[x][y] > maxHeight {
				out[Point{x, y}] = true
				maxHeight = f[x][y]
			}
		}

		maxHeight = -1
		for y := len(f[x]) - 1; y >= 0; y-- {
			if f[x][y] > maxHeight {
				out[Point{x, y}] = true
				maxHeight = f[x][y]
			}
		}

	}

	// Check left/right
	for y := 0; y < len(f[0]); y++ {
		maxHeight = -1
		for x := 0; x < len(f); x++ {
			if f[x][y] > maxHeight {
				out[Point{x, y}] = true
				maxHeight = f[x][y]
			}
		}

		maxHeight = -1
		for x := len(f) - 1; x >= 0; x-- {
			if f[x][y] > maxHeight {
				out[Point{x, y}] = true
				maxHeight = f[x][y]
			}
		}

	}
	return out
}

type moveFunc func(y, x int) (int, int)

func dirScore(f [][]int, y, x int, move moveFunc) int {
	cury, curx := y, x
	moves := 0
	for cury > 0 && cury < len(f)-1 && curx > 0 && curx < len(f)-1 {
		cury, curx = move(cury, curx)
		moves++
		dbg("     dirScore: (%v,%v)[%d] -> (%v,%v)[%d]", y, x, f[y][x], cury, curx, f[cury][curx])
		if f[cury][curx] >= f[y][x] {
			break
		}

	}

	dbg("     [found] dirScore: (%v,%v) -> (%v,%v), score: %d", y, x, cury, curx, moves)
	return moves
}

func getTreeScore(f [][]int, y, x int) int {
	dbg("Tree (%v, %v) [%v]", y, x, f[y][x])
	up := dirScore(f, y, x, func(y, x int) (int, int) {
		return y - 1, x
	})
	dbg(" - UP: %d", up)

	left := dirScore(f, y, x, func(y, x int) (int, int) {
		return y, x - 1
	})
	dbg(" - LEFT: %d", left)

	right := dirScore(f, y, x, func(y, x int) (int, int) {
		return y, x + 1
	})
	dbg(" - RIGHT: %d", right)

	down := dirScore(f, y, x, func(y, x int) (int, int) {
		return y + 1, x
	})
	dbg(" - DOWN: %d", down)

	return up * down * left * right
}

func getMaxScenicScore(f [][]int) int {
	maxScore := 0

	dbg("len f: %d/%d", len(f), len(f[0]))
	for x := 0; x < len(f); x++ {
		for y := 0; y < len(f); y++ {
			dbg("getTreeScore (%d,%d)", x, y)
			cur := getTreeScore(f, x, y)
			if cur > maxScore {
				maxScore = cur
			}
		}
	}

	return maxScore
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

	forest := make([][]int, len(lines))

	for row, line := range lines {
		forest[row] = make([]int, len(line))
		for col, v := range line {
			forest[row][col], _ = strconv.Atoi(string(v))
		}

	}
	dbg("forest: %v", forest)
	visible := findVisible(forest)
	dbg("visible: %v", visible)
	part1 = len(visible)

	part2 = getMaxScenicScore(forest)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
