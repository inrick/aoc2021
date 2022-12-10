package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
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

// n   segments   #
//
// 0   abcefg     6
// 1   cf         2
// 2   acdeg      5
// 3   acdfg      5
// 4   bcdf       4
// 5   abdfg      5
// 6   abdefg     6
// 7   acf        3
// 8   abcdefg    7
// 9   abcdfg     6

type Display struct {
	I [10]byte
	O [4]byte
}

func Pack(s string) byte {
	var b byte
	for _, ch := range []byte(s) {
		assert('a' <= ch && ch <= 'g')
		b |= 1 << (ch - 'a')
	}
	return b
}

func (d *Display) Decode() int {
	fwd := make(map[byte]int)
	bwd := make(map[int]byte)
	for len(fwd) < 10 {
		for _, d := range d.I {
			switch bits.OnesCount8(d) {
			case 2:
				fwd[d] = 1
				bwd[1] = d
			case 3:
				fwd[d] = 7
				bwd[7] = d
			case 4:
				fwd[d] = 4
				bwd[4] = d
			case 5:
				if one, ok := bwd[1]; ok {
					if one|d == d {
						fwd[d] = 3
						bwd[3] = d
					} else if six, ok := bwd[6]; ok {
						if eight, ok := bwd[8]; ok {
							if six|d == eight {
								fwd[d] = 2
								bwd[2] = d
							} else {
								fwd[d] = 5
								bwd[5] = d
							}
						}
					}
				}
			case 6:
				if four, ok := bwd[4]; ok {
					if four|d == d {
						fwd[d] = 9
						bwd[9] = d
					} else if one, ok := bwd[1]; ok && one|d == d {
						fwd[d] = 0
						bwd[0] = d
					}
				}
				if one, ok := bwd[1]; ok {
					if eight, ok := bwd[8]; ok {
						if one|d == eight {
							fwd[d] = 6
							bwd[6] = d
						}
					}
				}
			case 7:
				fwd[d] = 8
				bwd[8] = d
			default:
				panic(d)
			}
		}
	}
	n := 0
	for _, d := range d.O {
		n = 10*n + fwd[d]
	}
	return n
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var displays []Display
	for sc.Scan() {
		var d Display
		ll := strings.Split(sc.Text(), " ")
		assert(len(ll) == 15)
		for i := 0; i < 10; i++ {
			d.I[i] = Pack(ll[i])
		}
		assert(ll[10] == "|")
		for i := 0; i < 4; i++ {
			d.O[i] = Pack(ll[11+i])
		}
		displays = append(displays, d)
	}
	ck(sc.Err())

	{
		n := 0
		for _, d := range displays {
			for _, o := range d.O {
				switch bits.OnesCount8(o) {
				case 2, 3, 4, 7:
					n++
				}
			}
		}
		fmt.Printf("a) %d\n", n)
	}

	{
		n := 0
		for _, d := range displays {
			n += d.Decode()
		}
		fmt.Printf("b) %d\n", n)
	}
}
