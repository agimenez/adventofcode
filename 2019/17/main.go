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
		b.WriteRune(rune(c))
		if c == '\n' {
			break
		}
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
		if l[0] == '\n' {
			break
		}
		r.image = append(r.image, l)
	}
}

func DrawImage(img []string) {
	for i := range img {
		fmt.Print(img[i])
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
