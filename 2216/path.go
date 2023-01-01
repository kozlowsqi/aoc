package main

import "fmt"

type node2 struct {
	v1 *valve
	v2 *valve

	t1 int
	t2 int

	g int
	h int
	p []string
}

func (n node2) String() string {
	return fmt.Sprintf("%d/%d %d \t %s %s %v", n.t1, n.t2, n.g, n.v1.label, n.v2.label, n.p)
}

var cmp2 heap[node2] = func(a, b node2) bool {
	return a.g+a.h > b.g+b.h
}

type node struct {
	v *valve
	t int
	g int
	h int
	p []string
}

func (n node) String() string {
	return fmt.Sprintf("%d/%d \t %s %v", n.t, n.g, n.v.label, n.p)
}

var cmp heap[node] = func(a, b node) bool {
	return a.g+a.h > b.g+b.h
}
