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
	var fish0 []int
	assert(sc.Scan())
	for _, s := range strings.Split(sc.Text(), ",") {
		n, err := strconv.Atoi(s)
		ck(err)
		fish0 = append(fish0, n)
	}
	assert(!sc.Scan())
	ck(sc.Err())

	fmt.Printf("a) %d\n", fishAfterDays(fish0, 80))
	fmt.Printf("b) %d\n", fishAfterDays(fish0, 256))
}

type School struct {
	N, Val int
}

func fishAfterDays(fish0 []int, days int) int {
	school := make([]School, len(fish0), len(fish0)+days)
	for i, f := range fish0 {
		school[i] = School{N: 1, Val: f}
	}
	for day := 0; day < days; day++ {
		newfish := 0
		for i, s := range school {
			if s.Val == 0 {
				school[i].Val = 6
				newfish += s.N
			} else {
				school[i].Val--
			}
		}
		school = append(school, School{N: newfish, Val: 8})
	}

	n := 0
	for _, s := range school {
		n += s.N
	}
	return n
}
