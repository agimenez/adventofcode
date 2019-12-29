package intcode

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/agimenez/adventofcode2019/utils"
)

type Program struct {
	mem  []int
	pc   int
	base int
}

func NewProgram(p string) *Program {

	pr := &Program{
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
func (p *Program) ensureAddr(addr int) {
	extend := addr - len(p.mem) + 1
	if addr >= len(p.mem) {
		p.mem = append(p.mem, make([]int, extend)...)
	}
}

func (p *Program) SetMem(addr, val int) {
	p.ensureAddr(addr)

	utils.Dbg(3, "  (SetMem) p.mem[%d] = %d", addr, val)
	p.mem[addr] = val
}

func (p *Program) getMem(addr int) int {
	p.ensureAddr(addr)

	utils.Dbg(3, "  (get) p.mem[%d] = %d", addr, p.mem[addr])
	return p.mem[addr]
}
func (p *Program) Run(input <-chan int, output chan<- int) {

	p.pc = 0
	for op := p.mem[p.pc]; op != 99; {
		utils.Dbg(4, "MEM = %v", p.mem)
		utils.Dbg(2, "pc = %d; op = %v", p.pc, op)

		opcode := op % 100
		utils.Dbg(2, "opcode = %v", opcode)

		switch opcode {
		case 1: // Add
			utils.Dbg(2, " INSTR = %v", p.mem[p.pc:p.pc+4])
			a, b, c := p.fetchParameter(1), p.fetchParameter(2), p.getAddrIndex(3)
			utils.Dbg(2, " ADD %d %d -> %d", a, b, c)
			p.SetMem(c, a+b)
			p.pc += 4
		case 2: // Mul
			utils.Dbg(2, " INSTR = %v", p.mem[p.pc:p.pc+4])
			a, b, c := p.fetchParameter(1), p.fetchParameter(2), p.getAddrIndex(3)
			utils.Dbg(2, " MUL %d %d -> %d", a, b, c)
			p.SetMem(c, a*b)
			p.pc += 4
		case 3: // In
			utils.Dbg(2, " INSTR = %v", p.mem[p.pc:p.pc+2])
			var in, dst int
			in, dst = <-input, p.getAddrIndex(1)
			utils.Dbg(2, " IN  %d -> mem[%d]", in, dst)
			p.SetMem(dst, in)
			p.pc += 2
		case 4: // Out
			utils.Dbg(2, " INSTR = %v", p.mem[p.pc:p.pc+2])
			src := p.fetchParameter(1)
			utils.Dbg(2, " OUT %d", src)

			output <- src
			p.pc += 2

		case 5: // JMP IF TRUE
			utils.Dbg(2, " INSTR = %v", p.mem[p.pc:p.pc+3])
			tst, newpc := p.fetchParameter(1), p.fetchParameter(2)
			utils.Dbg(2, " JMP %d if %d != 0", newpc, tst)
			if tst != 0 {
				p.pc = newpc
			} else {
				p.pc += 3
			}

		case 6: // JMP IF FALSE
			utils.Dbg(2, " INSTR = %v", p.mem[p.pc:p.pc+3])
			tst, newpc := p.fetchParameter(1), p.fetchParameter(2)
			utils.Dbg(2, " JMP %d if %d == 0", newpc, tst)
			if tst == 0 {
				p.pc = newpc
			} else {
				p.pc += 3
			}

		case 7: // LT
			utils.Dbg(2, " INSTR = %v", p.mem[p.pc:p.pc+4])
			first, second, dst := p.fetchParameter(1), p.fetchParameter(2), p.getAddrIndex(3)
			utils.Dbg(2, " LT %d %d %d", first, second, dst)
			if first < second {
				p.SetMem(dst, 1)
			} else {
				p.SetMem(dst, 0)
			}

			p.pc += 4

		case 8: // EQ
			utils.Dbg(2, " INSTR = %v", p.mem[p.pc:p.pc+4])
			first, second, dst := p.fetchParameter(1), p.fetchParameter(2), p.getAddrIndex(3)
			utils.Dbg(2, " EQ %d %d %d", first, second, dst)
			if first == second {
				p.SetMem(dst, 1)
			} else {
				p.SetMem(dst, 0)
			}
			p.pc += 4

		case 9: // RELBASE
			utils.Dbg(2, " INSTR = %v", p.mem[p.pc:p.pc+2])
			offset := p.fetchParameter(1)
			utils.Dbg(2, " RELBASE  %d", offset)

			p.base += offset
			p.pc += 2

		default:
			log.Fatalf("Bad opcode = %v", op)
		}
		utils.Dbg(4, " MEM = %v", p.mem)

		op = p.mem[p.pc]
	}
}

func (p *Program) instructionMode(offset int) int {
	opcode := p.mem[p.pc]
	return opcode / int(math.Pow10(offset+1)) % 10
}

// This is for writing to mem operations (IN, etc)
func (p *Program) getAddrIndex(n int) int {
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

func (p *Program) fetchParameter(n int) int {
	mode := p.instructionMode(n)
	parameter := p.mem[p.pc+n]
	utils.Dbg(3, "   param[%d](%d) mode: %d, base %d, memsize %d", n, parameter, mode, p.base, len(p.mem))

	if mode == 0 {
		// position mode
		val := p.getMem(parameter)
		utils.Dbg(3, "   (posmode) -> mem[%d] = %d", parameter, val)
		return val
	} else if mode == 2 {
		// relative mode
		val := p.getMem(p.base + parameter)
		utils.Dbg(3, "   (relmode) -> mem[%d+%d] = %d", p.base, parameter, val)
		return val
	}

	// immediate mode
	utils.Dbg(3, "   (immmode) -> = %d", parameter)
	return parameter
}
