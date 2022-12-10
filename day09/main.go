package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Heights [][]byte

func (h Heights) At(u Pos) byte {
	Nx, Ny := len(h[0]), len(h)
	if u.X < 0 || Nx <= u.X || u.Y < 0 || Ny <= u.Y {
		return 255
	}
	return h[u.Y][u.X]
}

func Neighbors(u Pos) [4]Pos {
	return [...]Pos{
		Pos{u.X - 1, u.Y},
		Pos{u.X + 1, u.Y},
		Pos{u.X, u.Y - 1},
		Pos{u.X, u.Y + 1},
	}
}

type Pos struct {
	X, Y int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var h Heights
	for sc.Scan() {
		bb := sc.Bytes()
		row := make([]byte, len(bb))
		for i, c := range bb {
			assert('0' <= c && c <= '9')
			row[i] = c - '0'
		}
		h = append(h, row)
	}
	ck(sc.Err())

	Nx, Ny := len(h[0]), len(h)
	var sinks []Pos
	{
		risk := 0
		for y := 0; y < Ny; y++ {
			for x := 0; x < Nx; x++ {
				u := Pos{x, y}
				val := h.At(u)
				lt := true
				for _, v := range Neighbors(u) {
					lt = lt && val < h.At(v)
				}
				if lt {
					risk += 1 + int(val)
					sinks = append(sinks, Pos{x, y})
				}
			}
		}
		fmt.Printf("a) %d\n", risk)
	}

	{
		sizes := make([]int, len(sinks))
		for i, sink := range sinks {
			q := []Pos{sink}
			seen := make(map[Pos]bool)
			seen[sink] = true
			size := 0
			for len(q) > 0 {
				u := q[0]
				q = q[1:]
				size++
				for _, v := range Neighbors(u) {
					if !seen[v] && h.At(v) < 9 {
						q = append(q, v)
						seen[v] = true
					}
				}
			}
			sizes[i] = size
		}

		sort.Slice(sizes, func(i, j int) bool {
			return sizes[i] > sizes[j]
		})

		fmt.Printf("b) %d\n", sizes[0]*sizes[1]*sizes[2])
	}
}
