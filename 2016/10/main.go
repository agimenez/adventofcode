package main

import (
	"flag"
	"fmt"
	. "github.com/agimenez/adventofcode/utils"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	debug bool
)

func dbg(f string, v ...interface{}) {
	if debug {
		fmt.Printf(f+"\n", v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
}
func main() {
	flag.Parse()

	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	part1, part2, dur1, dur2 := solve(lines)
	log.Printf("Part 1 (%v): %v\n", dur1, part1)
	log.Printf("Part 2 (%v): %v\n", dur2, part2)

}

func solve(lines []string) (int, int, time.Duration, time.Duration) {
	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 := solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 := solve2(lines)
	dur[1] = time.Since(now)

	return part1, part2, dur[0], dur[1]

}

type Factory struct {
	bots    map[int]*Bot
	outputs map[int]chan int
	outBins map[int]int
}

func NewFactory() Factory {
	return Factory{
		bots:    map[int]*Bot{},
		outputs: map[int]chan int{},
	}
}

type Bot struct {
	Id int

	In        chan int
	Low, High chan int
}

func NewBot(id int) *Bot {
	return &Bot{
		Id:   id,
		In:   make(chan int, 2),
		Low:  nil,
		High: nil,
	}
}

func (f Factory) SendValue(v int, botId int) {
	bot := f.GetOrCreateBot(botId)

	bot.In <- v
}

func (f Factory) GetOrCreateBot(botId int) *Bot {
	bot, ok := f.bots[botId]
	if !ok {
		bot = NewBot(botId)
		f.bots[botId] = bot
	}

	return bot
}

func (f Factory) GetOrCreateOutput(outputId int) chan int {
	dbg("Creating output %v", outputId)
	out, ok := f.outputs[outputId]
	if !ok {
		out = make(chan int)
		f.outputs[outputId] = out
	}

	return out
}

func (f Factory) GetOutput(id int) (int, bool) {
	if v, ok := f.outBins[id]; ok {
		return v, ok
	}
	select {

	case v := <-f.outputs[id]:
		dbg("[OUT: %3d] Got value: %v", v)
		f.outBins[id] = v
		return v, true

	default:
		return 0, false

	}
}

func (f Factory) Connect(srcId int, outLowType string, outLowId int, outHighType string, outHighId int) {
	srcBot := f.GetOrCreateBot(srcId)

	if outLowType == "bot" {
		dstBot := f.GetOrCreateBot(outLowId)
		srcBot.Low = dstBot.In
	} else {
		dstOut := f.GetOrCreateOutput(outLowId)
		srcBot.Low = dstOut
	}

	if outHighType == "bot" {
		dstBot := f.GetOrCreateBot(outHighId)
		srcBot.High = dstBot.In
	} else {
		dstOut := f.GetOrCreateOutput(outHighId)
		srcBot.High = dstOut
	}

	go func(b *Bot) {
		dbg("[BOT %3d] Bot waiting for inputs...", b.Id)
		v1, v2 := <-b.In, <-b.In
		high := Max(v1, v2)
		low := Min(v1, v2)
		dbg("[BOT %3d] Got values: %v, %v", b.Id, high, low)

		if high == 61 && low == 17 {
			fmt.Println("RESULT: ", b.Id)
		}

		dbg("[BOT %3d] Writing to low...", b.Id)
		b.Low <- low
		dbg("[BOT %3d] Writing to high...", b.Id)
		b.High <- high
		dbg("[BOT %3d] DONE!", b.Id)
	}(srcBot)

}

func (f Factory) processInstruction(s string) {
	dbg("INSTRUCTION: %q", s)
	valueRE := regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
	wireRE := regexp.MustCompile(`bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d)`)

	vals := valueRE.FindStringSubmatch(s)
	if vals != nil {
		botId := ToInt(vals[2])
		value := ToInt(vals[1])
		f.SendValue(value, botId)

		return
	}

	wires := wireRE.FindStringSubmatch(s)
	if wires != nil {
		srcBot := ToInt(wires[1])

		botLowType := wires[2]   // bot/output for low value
		dstLo := ToInt(wires[3]) // bot/output number

		botHighType := wires[4]  // bot/output for high value
		dstHi := ToInt(wires[5]) // Id of destination for high value

		f.Connect(srcBot, botLowType, dstLo, botHighType, dstHi)

		return
	}

	panic("Unknown instruction: " + s)

}

func solve1(s []string) int {
	res := 0

	f := NewFactory()

	for _, l := range s {
		f.processInstruction(l)
		dbg("%+v", f)
	}
	for i := range f.outputs {
		dbg("Output %v: %v", i, <-f.outputs[i])
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
