package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
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

type program struct {
	mem  []int
	pc   int
	base int
}

func newProgram(p string) *program {

	pr := &program{
		mem:  []int{},
		pc:   0,
		base: 0,
	}

	pSlice := strings.Split(p, ",")
	for _, b := range pSlice {
		i, err := strconv.Atoi(b)
		if err != nil {
			log.Fatal(err)
		}

		pr.mem = append(pr.mem, i)
	}

	return pr
}

// ensureAddr ensures that a given address is reachable, by growing the memory slice
// if neccessary
func (p *program) ensureAddr(addr int) {
	extend := addr - len(p.mem) + 1
	if addr >= len(p.mem) {
		p.mem = append(p.mem, make([]int, extend)...)
	}
}

func (p *program) setMem(addr, val int) {
	p.ensureAddr(addr)

	dbg(2, "  (setMem) p.mem[%d] = %d", addr, val)
	p.mem[addr] = val
}

func (p *program) getMem(addr int) int {
	p.ensureAddr(addr)

	dbg(2, "  (get) p.mem[%d] = %d", addr, p.mem[addr])
	return p.mem[addr]
}
func (p *program) run(input <-chan int, output chan<- int) {

	for op := p.mem[p.pc]; op != 99; {
		dbg(4, "MEM = %v", p.mem)
		dbg(3, "pc = %d; op = %v", p.pc, op)

		opcode := op % 100
		dbg(3, "opcode = %v", opcode)

		switch opcode {
		case 1: // Add
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+4])
			a, b, c := p.fetchParameter(1), p.fetchParameter(2), p.getAddrIndex(3)
			dbg(3, " ADD %d %d -> %d", a, b, c)
			p.setMem(c, a+b)
			p.pc += 4
		case 2: // Mul
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+4])
			a, b, c := p.fetchParameter(1), p.fetchParameter(2), p.getAddrIndex(3)
			dbg(3, " MUL %d %d -> %d", a, b, c)
			p.setMem(c, a*b)
			p.pc += 4
		case 3: // In
			dbg(2, " INSTR = %v", p.mem[p.pc:p.pc+2])
			var in, dst int
			in, dst = <-input, p.getAddrIndex(1)
			dbg(2, " IN  %d -> mem[%d]", in, dst)
			p.setMem(dst, in)
			p.pc += 2
		case 4: // Out
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+2])
			src := p.fetchParameter(1)
			dbg(3, " OUT %d", src)

			output <- src
			p.pc += 2

		case 5: // JMP IF TRUE
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+3])
			tst, newpc := p.fetchParameter(1), p.fetchParameter(2)
			dbg(3, " JMP %d if %d != 0", newpc, tst)
			if tst != 0 {
				p.pc = newpc
			} else {
				p.pc += 3
			}

		case 6: // JMP IF FALSE
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+3])
			tst, newpc := p.fetchParameter(1), p.fetchParameter(2)
			dbg(3, " JMP %d if %d == 0", newpc, tst)
			if tst == 0 {
				p.pc = newpc
			} else {
				p.pc += 3
			}

		case 7: // LT
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+4])
			first, second, dst := p.fetchParameter(1), p.fetchParameter(2), p.getAddrIndex(3)
			dbg(3, " LT %d %d %d", first, second, dst)
			if first < second {
				p.setMem(dst, 1)
			} else {
				p.setMem(dst, 0)
			}

			p.pc += 4

		case 8: // EQ
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+4])
			first, second, dst := p.fetchParameter(1), p.fetchParameter(2), p.getAddrIndex(3)
			dbg(3, " EQ %d %d %d", first, second, dst)
			if first == second {
				p.setMem(dst, 1)
			} else {
				p.setMem(dst, 0)
			}
			p.pc += 4

		case 9: // RELBASE
			dbg(2, " INSTR = %v", p.mem[p.pc:p.pc+2])
			offset := p.fetchParameter(1)
			dbg(2, " RELBASE  %d", offset)

			p.base += offset
			p.pc += 2

		default:
			log.Fatalf("Bad opcode = %v", op)
		}
		dbg(4, " MEM = %v", p.mem)

		op = p.mem[p.pc]
	}
}

func (p *program) instructionMode(offset int) int {
	opcode := p.mem[p.pc]
	return opcode / int(math.Pow10(offset+1)) % 10
}

// This is for writing to mem operations (IN, etc)
func (p *program) getAddrIndex(n int) int {
	parameter := p.mem[p.pc+n]
	mode := p.instructionMode(n)

	if mode == 0 {
		return parameter
	} else if mode == 2 {
		return p.base + parameter
	} else {
		panic("unsupported immediate mode for writing")
	}
}

func (p *program) fetchParameter(n int) int {
	mode := p.instructionMode(n)
	parameter := p.mem[p.pc+n]
	dbg(2, "   param[%d](%d) mode: %d, base %d, memsize %d", n, parameter, mode, p.base, len(p.mem))

	if mode == 0 {
		// position mode
		val := p.getMem(parameter)
		dbg(2, "   (posmode) -> mem[%d] = %d", parameter, val)
		return val
	} else if mode == 2 {
		// relative mode
		val := p.getMem(p.base + parameter)
		dbg(2, "   (relmode) -> mem[%d+%d] = %d", p.base, parameter, val)
		return val
	}

	// immediate mode
	dbg(2, "   (immmode) -> = %d", parameter)
	return parameter
}

// Colours.
const (
	Black = 0
	White = 1
)

// Directions.
const (
	Up = iota
	Right
	Down
	Left
	DirLast
)

const (
	TurnLeft  = 0
	TurnRight = 1
)

type Robot struct {
	cpu     *program
	cam     chan int
	actions chan int
	dir     int
	cur     Point
	panels  map[Point]int
}

type Point struct {
	x, y int
}

var P0 = Point{0, 0}

func init() {
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.Parse()
}

func newRobot(code string) *Robot {
	return &Robot{
		cpu:     newProgram(code),
		cam:     make(chan int, 1),
		actions: make(chan int),
		dir:     Up,
		cur:     P0,
		panels:  make(map[Point]int),
	}
}

func (r *Robot) Run() {
	halt := make(chan bool)
	go func() {
		r.cpu.run(r.cam, r.actions)
		halt <- true
	}()

	for {
		dbg(1, "(%v) Sending %v to cam", r.cur, r.panels[r.cur])
		r.cam <- r.panels[r.cur]
		select {
		case color := <-r.actions:
			r.panels[r.cur] = color
			dbg(1, "-> got paint %d", color)

			dir := <-r.actions
			dbg(1, "-> got action %d", dir)
			r.DoTurn(dir)
			r.Forward()
		case <-halt:
			return
		}

	}
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (r *Robot) DoTurn(where int) {

	switch where {
	case TurnRight:
		r.dir = mod(r.dir+1, DirLast)
	case TurnLeft:
		r.dir = mod(r.dir-1, DirLast)
	default:
		panic("Unknown turn")
	}
	dbg(1, "Turn %d, new dir -> %d", where, r.dir)

}

func (r *Robot) Forward() {
	switch r.dir {
	case Up:
		r.cur.y++
	case Right:
		r.cur.x++
	case Down:
		r.cur.y--
	case Left:
		r.cur.x--
	}
	dbg(1, "Cur: %v", r.cur)
}

func main() {

	var in string
	fmt.Scan(&in)

	painter := newRobot(in)
	painter.Run()
	fmt.Printf("Painted panels: %d\n", len(painter.panels))

}
