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

type Line struct {
	x1, y1 int
	x2, y2 int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var lines []Line
	for sc.Scan() {
		var xa, ya, xb, yb int
		fmt.Sscanf(sc.Text(), "%d,%d -> %d,%d", &xa, &ya, &xb, &yb)
		// Arrange so that lines go from left to right
		var l Line
		if xa < xb {
			l.x1 = xa
			l.x2 = xb
			l.y1 = ya
			l.y2 = yb
		} else {
			l.x1 = xb
			l.x2 = xa
			l.y1 = yb
			l.y2 = ya
		}
		lines = append(lines, l)
	}
	ck(sc.Err())

	var xn, yn int
	for _, l := range lines {
		xn = Max(xn, l.x2)
		yn = Max(yn, l.y1)
		yn = Max(yn, l.y2)
	}

	{
		field := make([][]int, yn+1)
		for i := range field {
			field[i] = make([]int, xn+1)
		}

		for _, l := range lines {
			switch {
			case l.x1 == l.x2:
				ya, yb := Min(l.y1, l.y2), Max(l.y1, l.y2)
				for y := ya; y <= yb; y++ {
					field[y][l.x1]++
				}
			case l.y1 == l.y2:
				for x := l.x1; x <= l.x2; x++ {
					field[l.y1][x]++
				}
			default:
				// Ignore other lines
			}
		}

		overlap := 0
		for i := range field {
			for j := range field[i] {
				if n := field[i][j]; n > 1 {
					overlap++
				}
			}
		}

		fmt.Printf("a) %d\n", overlap)
	}

	{
		field := make([][]int, yn+1)
		for i := range field {
			field[i] = make([]int, xn+1)
		}

		for _, l := range lines {
			switch {
			case l.x1 == l.x2:
				ya, yb := Min(l.y1, l.y2), Max(l.y1, l.y2)
				for y := ya; y <= yb; y++ {
					field[y][l.x1]++
				}
			case l.y1 == l.y2:
				for x := l.x1; x <= l.x2; x++ {
					field[l.y1][x]++
				}
			default:
				// Anticipated a bit too much, the problem turned out to only need
				// lines with k = ±1.
				kn := l.y2 - l.y1
				kd := l.x2 - l.x1
				div := gcd(kn, kd)
				kn /= div
				kd /= div
				if kn < 0 {
					for x, y := l.x1, l.y1; x <= l.x2 && y >= l.y2; x, y = x+kd, y+kn {
						field[y][x]++
					}
				} else {
					for x, y := l.x1, l.y1; x <= l.x2 && y <= l.y2; x, y = x+kd, y+kn {
						field[y][x]++
					}
				}
			}
		}

		overlap := 0
		for i := range field {
			for j := range field[i] {
				if n := field[i][j]; n > 1 {
					overlap++
				}
			}
		}

		fmt.Printf("b) %d\n", overlap)
	}
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
