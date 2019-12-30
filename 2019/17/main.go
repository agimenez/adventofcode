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
	return Location{Point{l.coord.x + l.dir.x, l.coord.y + l.dir.y}, l.dir}
}

func (l Location) Valid(img []string) bool {
	if l.coord.y >= len(img) {
		return false
	}

	if l.coord.x >= len(img[0]) {
		return false
	}

	if img[l.coord.y][l.coord.x] != '#' {
		return false
	}

	return true
}

func (l Location) Right() Location {
	nl := l
	switch l.dir {
	case dirN:
		l.dir = dirE
	case dirE:
		l.dir = dirS
	case dirS:
		l.dir = dirW
	case dirW:
		l.dir = dirN
	}

	return nl
}

func (l Location) Left() Location {
	nl := l
	switch l.dir {
	case dirN:
		l.dir = dirW
	case dirE:
		l.dir = dirN
	case dirS:
		l.dir = dirE
	case dirW:
		l.dir = dirS
	}

	return nl
}
func (r *Robot) FindDirection(l Location) rune {
	if l.Forward().Valid(r.image) {
		return 'F'
	}

	if l.Right().Forward().Valid(r.image) {
		return 'R'
	}

	if l.Left().Forward().Valid(r.image) {
		return 'L'
	}

	// should not get here
	return 'F'
}

func (r *Robot) RunScaffolding(l Location) string {
	var b strings.Builder

	fwd := 0
	turn := r.FindDirection(l)
	b.WriteRune(turn)
	for {
		for l.Forward().Valid(r.image) {
			l = l.Forward()
			fwd++
		}
		b.WriteString(fmt.Sprintf("%d", fwd))
		fwd = 0
		turn := r.FindDirection(l)
		b.WriteRune(turn)
	}
}
func (r *Robot) RunPartTwo() int {
	go func() {
		r.cpu.SetMem(0, 2) // HACK THE CODE!!!
		r.cpu.Run(r.input, r.output)
	}()

	inputs := []string{
		"A,B,C",
		"R,12,L,12",
		"R,12,L,12",
		"R,12,L,12",
		"n",
	}

	// This was not clear from the exercise spec, but it seems that the first output
	// is just a snapshot of the current camera status. Then it asks for the values,
	// using a prompt
	r.ReadCam()
	p := r.Location()
	r.RunScaffolding(p)

	for _, cmd := range inputs {
		prompt := r.ReadLine()
		utils.Dbg(1, "> %s", prompt)
		r.WriteLine(cmd)
		utils.Dbg(1, "< %s", cmd)
	}

	return <-r.output
}
func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {

	var in string
	fmt.Scan(&in)

	r := newRobot(in)
	//r.Run()
	r.Paint()
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
