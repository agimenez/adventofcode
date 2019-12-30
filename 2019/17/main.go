package main

import (
	"fmt"
	"strings"

	"github.com/agimenez/adventofcode2019/intcode"
	"github.com/agimenez/adventofcode2019/utils"
)

type Robot struct {
	cpu    *intcode.Program
	input  chan int
	output chan int
	image  []string
}

type Point struct {
	x, y int
}

var P0 = Point{0, 0}

func newRobot(code string) *Robot {
	return &Robot{
		cpu:    intcode.NewProgram(code),
		input:  make(chan int),
		output: make(chan int),
		image:  []string{},
	}
}

func (r *Robot) Run() {
	go func() {
		r.cpu.Run(r.input, r.output)
		close(r.output)
	}()

	var b strings.Builder
	for {
		char, ok := <-r.output
		if !ok {
			break
		}
		utils.Dbg(1, "Char: %c (%d)", rune(char), rune(char))
		b.WriteRune(rune(char))
	}

	r.image = strings.Split(strings.TrimSpace(b.String()), "\n")

}

func (r *Robot) Paint() {
	for _, line := range r.image {
		fmt.Println(line)
	}
}

func (r *Robot) GetIntersections() []Point {
	intersections := []Point{}
	for y := 1; y < len(r.image)-1; y++ {
		for x := 1; x < len(r.image[y])-1; x++ {
			utils.Dbg(2, "Checking: {%d, %d}", y, x)
			if r.image[y][x] == '#' && r.IsIntersection(x, y) {
				utils.Dbg(1, " -> Int: {%d, %d}", y, x)
				intersections = append(intersections, Point{x, y})
			}
		}
	}

	return intersections
}

func (r *Robot) IsIntersection(x, y int) bool {
	return r.image[y-1][x] == '#' &&
		r.image[y+1][x] == '#' &&
		r.image[y][x-1] == '#' &&
		r.image[y][x+1] == '#'
}

func (r *Robot) SumAlignmentParameters() int {
	tot := 0
	for _, p := range r.GetIntersections() {
		tot += p.x * p.y
	}

	return tot
}

func (r *Robot) ReadLine() string {
	var b strings.Builder

	for {
		c := <-r.output
		//utils.Dbg(1, "Got %c (%d)", c, c)
		if c == '\n' {
			break
		}
		b.WriteRune(rune(c))
	}

	return b.String()
}

func (r *Robot) WriteLine(cmd string) {
	for _, c := range cmd {
		r.input <- int(c)
	}
	r.input <- '\n'
}

func (r *Robot) ReadCam() {
	for {
		l := r.ReadLine()

		// We don't return newlines in r.ReadLine(), so if an empty line comes from
		// the camera, we return an empty string instead of just a newline rune
		if len(l) == 0 {
			break
		}
		r.image = append(r.image, l)
	}
}

func DrawImage(img []string) {
	for i := range img {
		fmt.Println(img[i])
	}
}

func (r *Robot) DrawAreaAround(l Location) {
	if utils.Debug == 0 {
		return
	}

	fmt.Printf("len y = %d, len x = %d\n", len(r.image), len(r.image[0]))
	fmt.Printf("    x = %d (@ = %d)\n", l.coord.x-3, l.coord.x)
	for y := l.coord.y - 3; y < l.coord.y+3; y++ {
		fmt.Printf("%4d", y)
		for x := l.coord.x - 3; x < l.coord.x+3; x++ {
			if l.coord.x == x && l.coord.y == y {
				fmt.Print("@")
			} else {
				if y < 0 || y >= len(r.image) || x < 0 || x >= len(r.image[y]) {
					fmt.Print("X")
				} else {
					fmt.Printf("%c", r.image[y][x])
				}
			}
		}
		fmt.Println()
	}
}

type Location struct {
	coord Point
	dir   Direction
}

type Direction Point

var (
	dirN = Direction{0, -1}
	dirE = Direction{1, 0}
	dirS = Direction{0, 1}
	dirW = Direction{-1, 0}
)

func (r *Robot) Location() Location {
	for y := range r.image {
		for x := range r.image[y] {
			switch r.image[y][x] {
			case '^':
				return Location{Point{x, y}, dirN}
			case '>':
				return Location{Point{x, y}, dirE}
			case 'v':
				return Location{Point{x, y}, dirS}
			case '<':
				return Location{Point{x, y}, dirW}
			}
		}
	}

	return Location{}
}

func (l Location) Forward() Location {
	utils.Dbg(2, " -> PREFWD: %+v", l)
	l.coord.x += l.dir.x
	l.coord.y += l.dir.y
	utils.Dbg(2, " -> FWD: %+v", l)

	return l
}

func (l Location) Valid(img []string) bool {
	utils.Dbg(2, " -> Valid: %+v", l)
	if l.coord.y < 0 || l.coord.y >= len(img) {
		utils.Dbg(2, " -> y out of bounds")
		return false
	}

	if l.coord.x < 0 || l.coord.x >= len(img[0]) {
		utils.Dbg(2, " -> x out of bounds")
		return false
	}

	if img[l.coord.y][l.coord.x] != '#' {
		utils.Dbg(2, " -> img[%d][%d] = %c", l.coord.y, l.coord.x, img[l.coord.y][l.coord.x])
		return false
	}

	utils.Dbg(2, "   -> FOUND (%c)", img[l.coord.y][l.coord.x])
	return true
}

func (l Location) Right() Location {
	switch l.dir {
	case dirN:
		l.dir = dirE
	case dirE:
		l.dir = dirS
	case dirS:
		l.dir = dirW
	case dirW:
		l.dir = dirN
	default:
		panic("Can't get new direction")
	}
	utils.Dbg(2, " -> Right: New direction: %+v", l)

	return l
}

func (l Location) Left() Location {
	switch l.dir {
	case dirN:
		l.dir = dirW
	case dirE:
		l.dir = dirN
	case dirS:
		l.dir = dirE
	case dirW:
		l.dir = dirS
	default:
		panic("Can't get new direction")
	}
	utils.Dbg(2, " -> Left: New direction: %+v", l)

	return l
}

func (l Location) Turn(dir rune) Location {

	switch dir {
	case 'R':
		l = l.Right()
	case 'L':
		l = l.Left()
	}

	return l
}
func (r *Robot) FindDirection(l Location) rune {
	utils.Dbg(2, " FindDirection: test FWD")
	if l.Forward().Valid(r.image) {
		utils.Dbg(1, "l.Forward() = %v", l.Forward())
		return 'F'
	}

	utils.Dbg(2, " FindDirection: test Right")
	if l.Right().Forward().Valid(r.image) {
		return 'R'
	}

	utils.Dbg(2, " FindDirection: test Left")
	if l.Left().Forward().Valid(r.image) {
		return 'L'
	}

	// Halt
	return 'H'
}

func (r *Robot) RunScaffolding(l Location) string {
	var b strings.Builder

	utils.Dbg(1, "Starting search form location %v", l)
	fwd := 0
	turn := r.FindDirection(l)
	l = l.Turn(turn)
	b.WriteRune(turn)
	b.WriteRune(',')

	for {
		utils.Dbg(1, "Path: '%s', checking FWD", b.String())
		r.DrawAreaAround(l)
		for l.Forward().Valid(r.image) {
			l = l.Forward()
			utils.Dbg(1, "NEW location: %+v", l)
			fwd++
			r.DrawAreaAround(l)
		}
		b.WriteString(fmt.Sprintf("%d,", fwd))
		fwd = 0
		turn := r.FindDirection(l)
		if turn == 'H' {
			break
		}
		l = l.Turn(turn)
		b.WriteString(fmt.Sprintf("%c,", turn))
	}

	return b.String()
}
func (r *Robot) RunPartTwo() int {
	go func() {
		r.cpu.SetMem(0, 2) // HACK THE CODE!!!
		r.cpu.Run(r.input, r.output)
	}()

	// This part I did manually, discovering the path using Robot.RunScaffolding(), since right now I'm not feeling like implementing a
	// simple dictionary based compressor
	// R,12,L,10,L,10,L,6,L,12,R,12,L,4,R,12,L,10,L,10,L,6,L,12,R,12,L,4,L,12,R,12,L,6,L,6,L,12,R,12,L,4,L,12,R,12,L,6,R,12,L,10,L,10,L,12,R,12,L,6,L,12,R,12,L,6
	// Main: A,B,A,B,C,B,C,A,C,C
	// A=R,12,L,10,L,10
	// B=L,6,L,12,R,12,L,4
	// C=L,12,R,12,L,6

	inputs := []string{
		"A,B,A,B,C,B,C,A,C,C",
		"R,12,L,10,L,10",
		"L,6,L,12,R,12,L,4",
		"L,12,R,12,L,6",
		"n",
	}

	// This was not clear from the exercise spec, but it seems that the first output
	// is just a snapshot of the current camera status. Then it asks for the values,
	// using a prompt
	r.ReadCam()
	//DrawImage(r.image)
	p := r.Location()
	path := r.RunScaffolding(p)
	fmt.Printf("Part two path: %s\n", path)

	for _, cmd := range inputs {
		prompt := r.ReadLine()
		utils.Dbg(1, "> %s", prompt)
		r.WriteLine(cmd)
		utils.Dbg(1, "< %s", cmd)
	}

	for {
		c := <-r.output
		utils.Dbg(1, "Got output: %c (%d)", c, c)
		if c > 127 {
			return c
		}
	}
}
func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {

	var in string
	fmt.Scan(&in)

	r := newRobot(in)
	r.Run()
	//r.Paint()
	part1 := r.SumAlignmentParameters()
	fmt.Printf("Part one: %#v\n", part1)

	// reset program
	r = newRobot(in)
	dust := r.RunPartTwo()

	fmt.Printf("Part two: %#v\n", dust)

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

func (p Point) Min(p2 Point) Point {
	return Point{
		x: min(p.x, p2.x),
		y: min(p.y, p2.y),
	}
}

func (p Point) Max(p2 Point) Point {
	return Point{
		x: max(p.x, p2.x),
		y: max(p.y, p2.y),
	}
}
