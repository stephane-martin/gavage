package main

import (
	"math/rand"
	"time"
)

type MonthGenerator struct {
	Year          int
	Month         int
	TotalLogLines int
	nbGenerated   int
	esbuf         []ESLogLine
	buf           []LogLine
	start         time.Time
	end           time.Time
	r             *rand.Rand
}

func NewMonthGenerator(year int, month int, nbLogLines int) *MonthGenerator {
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, time.Month(month)+1, 1, 0, 0, 0, -1, time.UTC)
	return &MonthGenerator{
		Year:          year,
		Month:         month,
		TotalLogLines: nbLogLines,
		nbGenerated:   0,
		esbuf:         make([]ESLogLine, 10000),
		buf:           make([]LogLine, 10000),
		start:         start,
		end:           end,
		r:             rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *MonthGenerator) Next() (res []ESLogLine) {
	if g.nbGenerated >= g.TotalLogLines {
		return nil
	}
	rest := g.TotalLogLines - g.nbGenerated
	if rest >= 10000 {
		rest = 10000
	}
	g.esbuf = g.esbuf[0:rest]
	g.buf = g.buf[0:rest]
	var i int
	for i = 0; i < rest; i++ {
		_ = g.buf[i].Random(g.start, g.end, g.r)
		g.buf[i].ToES(&g.esbuf[i])
	}
	g.nbGenerated += rest
	return g.esbuf
}

func (g *MonthGenerator) Percent() int {
	return int(float64(100*g.nbGenerated) / float64(g.TotalLogLines))
}
