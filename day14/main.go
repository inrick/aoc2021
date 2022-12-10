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

type Rules map[string]string

func (r Rules) Add(left, right []byte) {
	r[string(left)] = string(right)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	assert(sc.Scan())
	stateInit := sc.Text()
	assert(sc.Scan() && sc.Text() == "")
	rules := make(Rules)
	for sc.Scan() {
		var left, right []byte
		fmt.Sscanf(sc.Text(), "%s -> %s", &left, &right)
		rules.Add(left, right)
	}
	ck(sc.Err())

	{
		state := []byte(stateInit)
		for i := 0; i < 10; i++ {
			state = step(state, rules)
		}
		nn := make(map[byte]int)
		for _, b := range state {
			nn[b]++
		}
		min, max := MinMax(nn)
		fmt.Printf("a) %d\n", max-min)
	}

	{
		pairs := make(map[string]int)
		for i := 1; i < len(stateInit); i++ {
			pairs[stateInit[i-1:i+1]]++
		}
		for i := 0; i < 40; i++ {
			pairs = stepPairs(pairs, rules)
		}

		// Since every character (except for the first and last ones) is part of
		// two pairs, it suffices to count the occurrences using the first
		// character in each pair and handling the last character as a special
		// case. (Or, the second character and treat the first one as a special
		// case.)
		c := make(map[byte]int)
		for p, n := range pairs {
			c[p[0]] += n
		}
		min, max := MinMax(c)
		// Check if the last character is the min or max one.
		last := stateInit[len(stateInit)-1]
		if min == c[last] {
			min++
		} else if max == c[last] {
			max++
		}

		fmt.Printf("b) %d\n", max-min)
	}
}

func step(state0 []byte, rules Rules) []byte {
	var state []byte
	state = append(state, state0[0])
	for i := 1; i < len(state0); i++ {
		left := state0[i-1 : i+1]
		if right, ok := rules[string(left)]; ok {
			state = append(state, []byte(right)[0], state0[i])
		} else {
			state = append(state, state0[i])
		}
	}
	return state
}

func stepPairs(pairsInit map[string]int, rules Rules) map[string]int {
	// Note that the new pairs is empty since every insertion breaks the pair
	// from which it came.
	pairs := make(map[string]int)
	for from, to := range rules {
		pairs[string(from[0])+to] += pairsInit[from]
		pairs[to+string(from[1])] += pairsInit[from]
	}
	return pairs
}

func MinMax(nn map[byte]int) (int, int) {
	min, max := 1<<63-1, -(1 << 63)
	for _, n := range nn {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}
