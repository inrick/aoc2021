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

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var crabs []int
	assert(sc.Scan())
	for _, s := range strings.Split(sc.Text(), ",") {
		n, err := strconv.Atoi(s)
		ck(err)
		crabs = append(crabs, n)
	}
	assert(!sc.Scan())
	ck(sc.Err())

	min, max := Min(crabs), Max(crabs)

	{
		m := 1 << 30
		for i := min; i <= max; i++ {
			sum := 0
			for _, n := range crabs {
				sum += Abs(n - i)
			}
			if sum < m {
				m = sum
			}
		}
		fmt.Printf("a) %d\n", m)
	}

	{
		m := 1<<63 - 1
		for i := min; i <= max; i++ {
			sum := 0
			for _, n := range crabs {
				d := Abs(n - i)
				sum += (d + 1) * d / 2
			}
			if sum < m {
				m = sum
			}
		}
		fmt.Printf("a) %d\n", m)
	}
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Min(nn []int) int {
	m := 1<<63 - 1
	for _, n := range nn {
		if n < m {
			m = n
		}
	}
	return m
}

func Max(nn []int) int {
	m := -(1 << 63)
	for _, n := range nn {
		if n > m {
			m = n
		}
	}
	return m
}
