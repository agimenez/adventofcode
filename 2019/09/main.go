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
	mem []int
	pc  int
}

func newProgram(p string) *program {

	pr := &program{
		mem: []int{},
		pc:  0,
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

func (p *program) run(input <-chan int, output chan<- int) {

	for op := p.mem[p.pc]; op != 99; {
		dbg(3, "MEM = %v", p.mem)
		dbg(3, "pc = %d; op = %v", p.pc, op)

		opcode := op % 100
		dbg(3, "opcode = %v", opcode)

		switch opcode {
		case 1: // Add
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+4])
			a, b, c := p.fetchParameter(1), p.fetchParameter(2), p.mem[p.pc+3]
			dbg(3, " ADD %d %d -> %d", a, b, c)
			p.mem[c] = a + b
			p.pc += 4
		case 2: // Mul
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+4])
			a, b, c := p.fetchParameter(1), p.fetchParameter(2), p.mem[p.pc+3]
			dbg(3, " MUL %d %d -> %d", a, b, c)
			p.mem[c] = a * b
			p.pc += 4
		case 3: // In
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+2])
			var in, dst int
			in, dst = <-input, p.mem[p.pc+1]
			dbg(3, " IN  %d -> %d", in, dst)
			p.mem[dst] = in
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
			first, second, dst := p.fetchParameter(1), p.fetchParameter(2), p.mem[p.pc+3]
			dbg(3, " LT %d %d %d", first, second, dst)
			if first < second {
				p.mem[dst] = 1
			} else {
				p.mem[dst] = 0
			}

			p.pc += 4

		case 8: // EQ
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+4])
			first, second, dst := p.fetchParameter(1), p.fetchParameter(2), p.mem[p.pc+3]
			dbg(3, " EQ %d %d %d", first, second, dst)
			if first == second {
				p.mem[dst] = 1
			} else {
				p.mem[dst] = 0
			}
			p.pc += 4

		default:
			log.Fatalf("Bad opcode = %v", op)
		}
		dbg(3, " MEM = %v", p.mem)

		op = p.mem[p.pc]
	}
}

func (p *program) fetchParameter(n int) int {
	opcode := p.mem[p.pc]
	parameter := p.mem[p.pc+n]
	mode := opcode / int(math.Pow10(n+1)) % 10

	dbg(3, "  (fetch) param[%d](%d) mode: %d", n, parameter, mode)
	if mode == 0 {
		return p.mem[parameter]
	}

	return parameter
}

func init() {
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.Parse()
}

func main() {

	var in string
	fmt.Scan(&in)

	log.Printf("input: %v", in)

}