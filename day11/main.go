package main

import (
	"bufio"
	"fmt"
	"os"
)

func ck(err error) {
	if err != nil {
		panic(err)
	}
}

func assert(b bool) {
	if !b {
		panic("assert failed")
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var l Level
	for sc.Scan() {
		bb := sc.Bytes()
		row := make([]int, len(bb))
		for i, ch := range bb {
			assert('0' <= ch && ch <= '9')
			row[i] = int(ch - '0')
		}
		l = append(l, row)
	}
	ck(sc.Err())

	{
		nflashed := 0
		for step := 0; step < 100; step++ {
			nflashed += l.Step()
		}
		fmt.Printf("a) %d\n", nflashed)
	}

	{
		N := len(l[0]) * len(l)
		step := 0
		for {
			nflashed := l.Step()
			step++
			if nflashed == N {
				break
			}
		}
		// Since we already ran 100 steps for part a)
		fmt.Printf("b) %d\n", step+100)
	}
}

type Pos struct {
	X, Y int
}

type Level [][]int

func (l Level) At(p Pos) int {
	x, y := p.X, p.Y
	if x < 0 || y < 0 || len(l[0]) <= x || len(l) <= y {
		return -1
	}
	return l[y][x]
}

func (l Level) Step() int {
	var moves []Pos
	flashed := make(map[Pos]bool)

	for y := range l {
		for x := range l[y] {
			p := Pos{x, y}
			l[y][x]++
			if l.At(p) > 9 {
				moves = append(moves, p)
				flashed[p] = true
			}
		}
	}

	for len(moves) > 0 {
		p := moves[0]
		moves = moves[1:]
		for _, q := range Neighbors(p) {
			if l.At(q) != -1 {
				l[q.Y][q.X]++
			}
			if l.At(q) > 9 && !flashed[q] {
				moves = append(moves, q)
				flashed[q] = true
			}
		}
	}

	for p, _ := range flashed {
		l[p.Y][p.X] = 0
	}

	return len(flashed)
}

func Neighbors(p Pos) [8]Pos {
	x, y := p.X, p.Y
	return [...]Pos{
		Pos{x - 1, y},
		Pos{x - 1, y + 1},
		Pos{x, y + 1},
		Pos{x + 1, y + 1},
		Pos{x + 1, y},
		Pos{x + 1, y - 1},
		Pos{x, y - 1},
		Pos{x - 1, y - 1},
	}
}
