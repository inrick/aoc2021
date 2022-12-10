package main

import (
	"bufio"
	"fmt"
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

func bool2int(b bool) int {
	var i int
	if b {
		i = 1
	}
	return i
}

type Node string

func (n Node) OnlyVisitOnce() bool {
	return n != Node(strings.ToUpper(string(n)))
}

func (n Node) In(seen []Node) int {
	c := 0
	for _, m := range seen {
		c += bool2int(n == m)
	}
	return c
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	edges := make(map[Node][]Node)
	for sc.Scan() {
		ss := strings.Split(sc.Text(), "-")
		assert(len(ss) == 2)
		node1, node2 := Node(ss[0]), Node(ss[1])
		edges[node1] = append(edges[node1], node2)
		edges[node2] = append(edges[node2], node1)
	}
	ck(sc.Err())

	fmt.Printf("a) %d\n", search1(edges, []Node{"start"}, Node("start")))
	fmt.Printf("b) %d\n", search2(edges, []Node{"start"}, Node("start"), false))
}

func search1(edges map[Node][]Node, visited []Node, node Node) int {
	if node == "end" {
		return 1
	}
	n := 0
	for _, other := range edges[node] {
		if other.OnlyVisitOnce() && other.In(visited) == 1 {
			continue
		}
		n += search1(edges, append(visited, other), other)
	}
	return n
}

func search2(edges map[Node][]Node, visited []Node, node Node, visitedTwice bool) int {
	if node == "end" {
		return 1
	}
	n := 0
	for _, other := range edges[node] {
		if other.OnlyVisitOnce() && other.In(visited) >= 1 {
			if visitedTwice || other == "start" {
				continue
			}
			n += search2(edges, append(visited, other), other, true)
		} else {
			n += search2(edges, append(visited, other), other, visitedTwice)
		}
	}
	return n
}
