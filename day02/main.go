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

func Add(u, v Vec) Vec {
	return Vec{u.X + v.X, u.Y + v.Y}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var instrs []Vec
	for sc.Scan() {
		ss := strings.Split(sc.Text(), " ")
		assert(len(ss) == 2)
		magn, err := strconv.Atoi(ss[1])
		ck(err)
		var instr Vec
		switch ss[0] {
		case "forward":
			instr = Vec{magn, 0}
		case "down":
			instr = Vec{0, magn}
		case "up":
			instr = Vec{0, -magn}
		default:
			panic(ss[0])
		}
		instrs = append(instrs, instr)
	}
	ck(sc.Err())

	var pos Vec
	for _, instr := range instrs {
		pos = Add(pos, instr)
	}
	fmt.Printf("a) %d\n", pos.X*pos.Y)

	aim := 0
	pos = Vec{0, 0}
	for _, instr := range instrs {
		switch {
		case instr.X == 0: // up/down
			aim += instr.Y
		case instr.Y == 0: // forward
			pos = Add(pos, instr)
			pos.Y += aim * instr.X
		default:
			panic(instr)
		}
	}
	fmt.Printf("b) %d\n", pos.X*pos.Y)
}
