package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var depths []int
	for sc.Scan() {
		n, err := strconv.Atoi(sc.Text())
		ck(err)
		depths = append(depths, n)
	}
	ck(sc.Err())
	increased := 0
	for i := 1; i < len(depths); i++ {
		if depths[i-1] < depths[i] {
			increased++
		}
	}

	fmt.Printf("a) %d\n", increased)

	increased3 := 0
	for i := 3; i < len(depths); i++ {
		a := sum(depths[i-3 : i])
		b := sum(depths[i-2 : i+1])
		if b > a {
			increased3++
		}
	}

	fmt.Printf("b) %d\n", increased3)
}

func sum(xx []int) int {
	n := 0
	for _, x := range xx {
		n += x
	}
	return n
}
