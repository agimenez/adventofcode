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
		log.Printf(fmt, v...)
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

func (p *program) setMem(dst, val int) {
	extend := dst - len(p.mem) + 1
	if dst >= len(p.mem) {
		p.mem = append(p.mem, make([]int, extend)...)
	}

	dbg(2, "  (setMem) p.mem[%d] = %d", dst, val)
	p.mem[dst] = val
}

func (p *program) run(input <-chan int, output chan<- int) {

	for op := p.mem[p.pc]; op != 99; {
		dbg(3, "MEM = %v", p.mem)
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
		dbg(3, " MEM = %v", p.mem)

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
		dbg(2, "   (posmode) -> mem[%d] = %d", parameter, p.mem[parameter])
		return p.mem[parameter]
	} else if mode == 2 {
		// relative mode
		dbg(2, "   (relmode) -> mem[%d+%d] = %d", p.base, parameter, p.mem[p.base+parameter])
		return p.mem[p.base+parameter]
	}

	// immediate mode
	dbg(2, "   (immmode) -> = %d", parameter)
	return parameter
}

func init() {
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.Parse()
}

func main() {

	var in string
	fmt.Scan(&in)

	chanIn := make(chan int)
	chanOut := make(chan int)

	program := newProgram(in)

	go func() {
		program.run(chanIn, chanOut)
		close(chanOut)
	}()

	chanIn <- 1
	output := []string{}
	for {
		val, ok := <-chanOut
		if !ok {
			break
		}
		output = append(output, fmt.Sprintf("%d", val))
	}

	log.Printf("List of unknownOps: %s", strings.Join(output[:len(output)-1], ", "))
	log.Printf("BOOST keycode: %s", output[len(output)-1])

	// part 2
	program = newProgram(in)
	chanIn = make(chan int)
	chanOut = make(chan int)

	go func() {
		program.run(chanIn, chanOut)
		close(chanOut)
	}()
	chanIn <- 2
	coords := <-chanOut
	log.Printf("Coordinates: %d", coords)

}
