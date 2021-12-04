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

type num struct {
	row int
	col int
}

type board struct {
	nums    map[int]num
	rows    [5]int
	cols    [5]int
	playing bool
}

func NewBoard() *board {
	return &board{
		nums:    make(map[int]num),
		playing: true,
	}
}
func (b *board) checkNumber(n int) bool {
	pos, ok := b.nums[n]
	if !ok {
		return false
	}

	b.rows[pos.row]++
	b.cols[pos.col]++
	delete(b.nums, n)
	if b.rows[pos.row] == 5 || b.cols[pos.col] == 5 {
		return true
	}

	return false
}

func (b *board) getScore() int {
	total := 0
	for n := range b.nums {
		total += n
	}

	return total

}

type bingo struct {
	draw   []int
	boards []*board
}

func (b *bingo) addDraw(in string) {
	s := strings.Split(in, ",")
	b.draw = make([]int, len(s))

	for i := range s {
		n, _ := strconv.Atoi(s[i])
		b.draw[i] = n
	}

}

func (b *bingo) addBoards(in []string) {
	brd := NewBoard()
	y := 0

	for i := range in {
		if in[i] == "" {
			b.boards = append(b.boards, brd)
			brd = NewBoard()
			y = 0
			continue
		}
		dbg("Add nums %v", in[i])

		rowNums := strings.Fields(in[i])
		for col := range rowNums {
			val, _ := strconv.Atoi(rowNums[col])
			brd.nums[val] = num{row: y, col: col}
			dbg(" -> %d: %v", val, brd.nums[val])
		}
		y++
	}

	b.boards = append(b.boards, brd)
}

func (b *bingo) play() ([]*board, []int) {
	boards := []*board{}
	draws := []int{}
	for _, draw := range b.draw {
		for _, board := range b.boards {
			if !board.playing {
				continue
			}

			win := board.checkNumber(draw)
			if win {
				boards = append(boards, board)
				draws = append(draws, draw)
				board.playing = false
			}
		}
	}

	return boards, draws
}
func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]

	b := &bingo{}
	b.addDraw(lines[0])

	b.addBoards(lines[2:])
	dbg("Bingo: %+v (%d)", b, len(b.boards))
	boards, nums := b.play()
	part1 = boards[0].getScore() * nums[0]
	part2 = boards[len(boards)-1].getScore() * nums[len(nums)-1]

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
