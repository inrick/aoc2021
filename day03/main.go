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

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	ck(sc.Err())

	w, h := len(lines[0]), len(lines)
	ones := make([]int, w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ones[j] += bool2int(lines[i][j] == '1')
		}
	}
	gamma := 0
	for j := 0; j < w; j++ {
		gamma |= bool2int(2*ones[j] > h) << (w - j - 1)
	}
	epsilon := ((1 << w) - 1) & ^gamma

	fmt.Printf("a) %d\n", gamma*epsilon)

	oxygen := filter(lines, true)
	co2 := filter(lines, false)

	fmt.Printf("b) %d\n", oxygen*co2)
}

func filter(lines []string, keepMostCommon bool) int {
	h, w := len(lines), len(lines[0])
	removed, remaining := make([]bool, h), h
	for j := 0; j < w; j++ {
		ones := 0
		for i, rem := range removed {
			ones += bool2int(!rem && lines[i][j] == '1')
		}
		toRemove := bool2int(2*ones >= remaining) ^ bool2int(keepMostCommon)
		for i := range removed {
			if !removed[i] && lines[i][j] == byte(toRemove)+'0' {
				removed[i] = true
				remaining--
			}
		}
		if remaining == 1 {
			for k, b := range removed {
				if !b {
					return str2num(lines[k])
				}
			}
		}
	}
	panic(nil)
}

func str2num(str string) (n int) {
	w := len(str)
	for i := range str {
		n |= int(str[i]-'0') << (w - i - 1)
	}
	return n
}

func bool2int(b bool) int {
	var i int
	if b {
		i = 1
	}
	return i
}
