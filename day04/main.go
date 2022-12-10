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
	var game Game
	ParseDrawn(sc, &game)
	for sc.Scan() {
		assert(sc.Text() == "")
		ParseBoard(sc, &game)
	}
	ck(sc.Err())

	fmt.Printf("a) %d\n", game.Win())
	fmt.Printf("b) %d\n", game.Lose())
}

type Board [5][5]int

type Game struct {
	Drawn  []int
	Boards []Board
}

type Seen []bool

func (g *Game) Win() int {
	seen := make(Seen, len(g.Drawn))
	for _, n := range g.Drawn {
		seen[n] = true
		for _, b := range g.Boards {
			if b.Won(seen) {
				return n * b.SumUnseen(seen)
			}
		}
	}
	return -1
}

func (g *Game) Lose() int {
	seen := make(Seen, len(g.Drawn))
	hasWon := make([]bool, len(g.Boards))
	lastWon := -1
	for _, n := range g.Drawn {
		seen[n] = true
		for j, b := range g.Boards {
			if !hasWon[j] && b.Won(seen) {
				hasWon[j] = true
				lastWon = n * b.SumUnseen(seen)
			}
		}
	}
	return lastWon
}

func (b *Board) Won(seen Seen) bool {
	var rows [5]bool
	var cols [5]bool
	for i := 0; i < 5; i++ {
		rows[i] = true
		cols[i] = true
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			s := seen[b[i][j]]
			rows[i] = rows[i] && s
			cols[j] = cols[j] && s
		}
	}
	for i := 0; i < 5; i++ {
		if rows[i] || cols[i] {
			return true
		}
	}

	return false
}

func (b *Board) SumUnseen(seen Seen) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			sum += bool2int(!seen[b[i][j]]) * b[i][j]
		}
	}
	return sum
}

func ParseDrawn(sc *bufio.Scanner, game *Game) {
	assert(sc.Scan())
	elems := strings.Split(sc.Text(), ",")
	for _, x := range elems {
		n, err := strconv.Atoi(x)
		ck(err)
		game.Drawn = append(game.Drawn, n)
	}
}

func ParseBoard(sc *bufio.Scanner, game *Game) {
	var board Board
	for i := 0; i < 5; i++ {
		assert(sc.Scan())
		elems := strings.Split(sc.Text(), " ")
		j := 0
		for _, x := range elems {
			if x == "" {
				continue
			}
			n, err := strconv.Atoi(x)
			ck(err)
			board[i][j] = n
			j++
		}
		assert(j == 5)
	}
	game.Boards = append(game.Boards, board)
}

func bool2int(b bool) int {
	var i int
	if b {
		i = 1
	}
	return i
}
