package main

import (
	"embed"
	"fmt"
	"io"
	"time"
)

//go:embed *.input
var input embed.FS

func p1(f io.Reader, y int) string {
	valves := parse(f)

	var open *ptree[node]

	cmp.queue(&open, node{
		v: valves["AA"],
		t: 30,
	})

	for open != nil {
		n := cmp.dequeue(&open)
		exhausted := true

	b:
		for l, c := range n.v.costs {
			for _, p := range n.p {
				if l == p {
					continue b
				}
			}

			v := valves[l]
			t := n.t - c

			if t < 0 {
				continue b
			}

			h := 0
		c:
			for l, c := range v.costs {
				if t-c < 0 {
					continue c
				}

				if l == n.v.label {
					continue c
				}

				for _, p := range n.p {
					if l == p {
						continue c
					}
				}

				h += valves[l].frate * (t - c)
			}

			cmp.queue(&open, node{
				v: v,
				t: t,
				g: n.g + t*v.frate,
				h: h,
				p: append(n.p, n.v.label),
			})

			exhausted = false
		}

		if exhausted {
			return fmt.Sprint(*n)
		}
	}

	return ""
}

func p2(f io.Reader, m int) string {
	t1 := time.Now()
	valves := parse(f)

	fmt.Println("P: ", time.Since(t1))

	var open, closed *ptree[node]

	cmp.queue(&open, node{
		v: valves["AA"],
		t: 26,
	})

	t2 := time.Now()
	for open != nil {
		n := cmp.dequeue(&open)
		cmp.queue(&closed, *n)

	b:
		for l, c := range n.v.costs {
			for _, p := range n.p {
				if l == p {
					continue b
				}
			}

			v := valves[l]
			t := n.t - c

			if t < 0 {
				continue b
			}

			h := 0
		c:
			for l, c := range v.costs {
				if t-c < 0 {
					continue c
				}

				if l == n.v.label {
					continue c
				}

				for _, p := range n.p {
					if l == p {
						continue c
					}
				}

				h += valves[l].frate * (t - c)
			}

			cmp.queue(&open, node{
				v: v,
				t: t,
				g: n.g + t*v.frate,
				h: h,
				p: append(n.p, n.v.label),
			})
		}

	}

	fmt.Println("G: ", time.Since(t2))
	t3 := time.Now()

	var ns []node
	for closed != nil {
		ns = append(ns, *cmp.dequeue(&closed))
	}

	var m1, m2 node
	for _, n1 := range ns {

		p0 := map[string]int{}

		p0[n1.v.label]++
		p0["AA"]--

		for _, p := range n1.p {
			p0[p]++
		}
	v:
		for _, n2 := range ns {

			if p0[n2.v.label] > 0 {
				continue v
			}

			for _, v := range n2.p {
				if p0[v] > 0 {
					continue v
				}
			}

			if m1.g+m2.g < n1.g+n2.g {
				m1, m2 = n1, n2
			}
		}
	}

	fmt.Println("C: ", len(ns), time.Since(t3))
	return fmt.Sprint(m1, "\n", m2, "\n", m1.g+m2.g)
}

func main() {
	f, err := input.Open("test.input")
	die(err)
	fmt.Println(p1(f, 10))

	f, err = input.Open("puzzle.input")
	die(err)
	fmt.Println(p1(f, 2000000))

	f, err = input.Open("test.input")
	die(err)
	fmt.Println(p2(f, 20))

	f, err = input.Open("puzzle.input")
	die(err)
	fmt.Println(p2(f, 4000000))
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
