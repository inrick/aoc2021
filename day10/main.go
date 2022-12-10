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

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	ck(sc.Err())

	{
		n := 0
		for _, line := range lines {
			err, _ := validate(line)
			n += err
		}
		fmt.Printf("a) %d\n", n)
	}

	{
		var scores []int
		for _, line := range lines {
			err, rem := validate(line)
			if err != 0 {
				continue
			}
			score := 0
			for i := len(rem); i > 0; i-- {
				score = score*5 + completionScore[matching[rem[i-1]]]
			}
			scores = append(scores, score)
		}

		sort.Slice(scores, func(i, j int) bool {
			return scores[i] < scores[j]
		})
		fmt.Printf("b) %d\n", scores[len(scores)/2])
	}
}

var errorScore = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completionScore = map[byte]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var matching = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func left(a byte) bool {
	switch a {
	case '(', '[', '{', '<':
		return true
	default:
		return false
	}
}

func validate(line string) (int, []byte) {
	var stack []byte
	for _, b := range []byte(line) {
		if left(b) {
			stack = append(stack, b)
			continue
		}
		n := len(stack)
		a := stack[n-1]
		stack = stack[:n-1]
		if matching[a] != b {
			return errorScore[b], stack
		}
	}
	return 0, stack
}
