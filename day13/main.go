package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

type Vec struct {
	X, Y int
}

type Fold struct {
	Dim byte
	Val int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var dots []Vec
	var folds []Fold
	state := 0
	for sc.Scan() {
		if sc.Text() == "" {
			state++
			continue
		}
		switch state {
		case 0:
			ss := strings.Split(sc.Text(), ",")
			assert(len(ss) == 2)
			x, err := strconv.Atoi(ss[0])
			ck(err)
			y, err := strconv.Atoi(ss[1])
			ck(err)
			dots = append(dots, Vec{x, y})
		case 1:
			var dim byte
			var val int
			fmt.Sscanf(sc.Text(), "fold along %c=%d", &dim, &val)
			folds = append(folds, Fold{dim, val})
		default:
			panic(state)
		}
	}
	ck(sc.Err())

	dotsA := run(dots, folds[0:1])
	fmt.Printf("a) %d\n", countUnique(dotsA))

	dotsB := run(dots, folds)
	fmt.Println("b)")
	draw(dotsB)
}

func run(dots0 []Vec, folds []Fold) []Vec {
	dots := make([]Vec, len(dots0))
	copy(dots, dots0)
	for _, fold := range folds {
		val := fold.Val
		switch fold.Dim {
		case 'x':
			for i, d := range dots {
				assert(d.X != val)
				if d.X > val {
					dots[i].X = 2*val - d.X
				}
			}
		case 'y':
			for i, d := range dots {
				assert(d.Y != val)
				if d.Y > val {
					dots[i].Y = 2*val - d.Y
				}
			}
		default:
			panic(fold.Dim)
		}
	}
	return dots
}

func countUnique(vec []Vec) int {
	seen := make(map[Vec]bool)
	for _, u := range vec {
		seen[u] = true
	}
	return len(seen)
}

func draw(vec []Vec) {
	maxX, maxY := 0, 0
	for _, u := range vec {
		if u.X > maxX {
			maxX = u.X
		}
		if u.Y > maxY {
			maxY = u.Y
		}
	}

	Nx, Ny := maxX+1, maxY+1
	chart := make([]byte, Nx*Ny)
	for i := range chart {
		chart[i] = ' '
	}
	for _, u := range vec {
		chart[Nx*u.Y+u.X] = '#'
	}
	for y := 0; y < Ny; y++ {
		fmt.Println(string(chart[Nx*y : Nx*(y+1)]))
	}
}
