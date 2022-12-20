package main

import (
	"embed"
	"fmt"
	"io"
)

//go:embed *.input
var input embed.FS

type position struct {
	x []int
	y []int

	log map[string]struct{}
}

func newposition(k int) position {
	return position{
		x:   make([]int, k),
		y:   make([]int, k),
		log: make(map[string]struct{}),
	}
}

func (p *position) move(x, y, k int) {
	for i := 0; i < k; i++ {
		p.log[fmt.Sprintf("%d,%d", p.x[len(p.x)-1], p.y[len(p.y)-1])] = struct{}{}

		p.y[0] += y
		p.x[0] += x

		for i := range p.x {
			if i == 0 {
				continue
			}

			h := p.x[i] <= p.x[i-1]+1 && p.x[i] >= p.x[i-1]-1
			v := p.y[i] <= p.y[i-1]+1 && p.y[i] >= p.y[i-1]-1

			if h && v {
				continue
			}

			var dx, dy int

			if p.x[i-1] < p.x[i] {
				dx--
			}

			if p.x[i-1] > p.x[i] {
				dx++
			}

			if p.y[i-1] < p.y[i] {
				dy--
			}

			if p.y[i-1] > p.y[i] {
				dy++
			}

			p.x[i] += dx
			p.y[i] += dy
		}
	}

	p.log[fmt.Sprintf("%d,%d", p.x[len(p.x)-1], p.y[len(p.y)-1])] = struct{}{}
}

func p1(f io.Reader) string {
	var p = newposition(2)
	for {
		var (
			m string
			n int
		)
		_, err := fmt.Fscanln(f, &m, &n)
		if err == io.EOF {
			break
		}
		die(err)

		if m == "U" {
			p.move(0, 1, n)
		}
		if m == "D" {
			p.move(0, -1, n)
		}
		if m == "R" {
			p.move(1, 0, n)
		}
		if m == "L" {
			p.move(-1, 0, n)
		}
	}

	return fmt.Sprint(len(p.log))
}

func p2(f io.Reader) string {
	var p = newposition(10)
	for {
		var (
			m string
			n int
		)
		_, err := fmt.Fscanln(f, &m, &n)
		if err == io.EOF {
			break
		}
		die(err)

		if m == "U" {
			p.move(0, 1, n)
		}
		if m == "D" {
			p.move(0, -1, n)
		}
		if m == "R" {
			p.move(1, 0, n)
		}
		if m == "L" {
			p.move(-1, 0, n)
		}
	}

	return fmt.Sprint(len(p.log))
}

func main() {
	f, err := input.Open("test.input")
	die(err)
	fmt.Println(p1(f))

	f, err = input.Open("puzzle.input")
	die(err)
	fmt.Println(p1(f))

	f, err = input.Open("test.input")
	die(err)
	fmt.Println(p2(f))

	f, err = input.Open("puzzle.input")
	die(err)
	fmt.Println(p2(f))
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
