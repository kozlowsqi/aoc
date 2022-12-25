package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
)

//go:embed *.input
var input embed.FS

type point struct {
	x, y int
}

type position struct {
	point
	cost int
	prev *position
}

type asterisk struct {
	open   []*position
	closed []*position

	c   *position
	end point
}

func (a *asterisk) next() bool {
	a.c = a.open[0]

	a.open = a.open[1:]
	a.closed = append(a.closed, a.c)

	return a.c.point != a.end
}

func (a *asterisk) insert(dx, dy int) {
	p := point{a.c.x + dx, a.c.y + dy}

	for prev := a.c; prev != nil; prev = prev.prev {
		if prev.point == p {
			return
		}
	}

	for _, n := range a.closed {
		if n.point == p {
			return
		}
	}

	for _, n := range a.open {
		if n.point == p {
			return
		}
	}

	l := &position{
		point: p,
		cost:  a.c.cost + 1,
		prev:  a.c,
	}

	for i := range a.open {
		if a.open[i].cost < l.cost {
			continue
		}

		a.open = append(a.open[:i], append([]*position{l}, a.open[i:]...)...)
		return
	}

	a.open = append(a.open, l)
}

func parse(f io.Reader) [][]int {
	fb := bufio.NewReader(f)

	board := [][]int{[]int{}}
	for {
		r, _, err := fb.ReadRune()
		if err == io.EOF {
			break
		}
		die(err)

		if r == '\n' {
			board = append(board, []int{})
		} else {
			board[len(board)-1] = append(board[len(board)-1], int(r))
		}
	}

	return board[:len(board)-1]
}

func p1(f io.Reader) string {
	board := parse(f)

	var start, end point
	for y := range board {
		for x := range board[y] {
			if board[y][x] == 'S' {
				start.x = x
				start.y = y
				board[y][x] = 'a'
			}
			if board[y][x] == 'E' {
				end.x = x
				end.y = y
				board[y][x] = 'z'
			}
		}
	}

	a := &asterisk{
		open: []*position{
			&position{point: start, cost: 0},
		},
		end: end,
	}

	for a.next() {
		if a.c.x > 0 && abs(board[a.c.y][a.c.x]-board[a.c.y][a.c.x-1]) <= 1 {
			a.insert(-1, 0)
		}
		if a.c.x < len(board[a.c.y])-1 && abs(board[a.c.y][a.c.x]-board[a.c.y][a.c.x+1]) <= 1 {
			a.insert(1, 0)
		}
		if a.c.y > 0 && abs(board[a.c.y][a.c.x]-board[a.c.y-1][a.c.x]) <= 1 {
			a.insert(0, -1)
		}
		if a.c.y < len(board)-1 && abs(board[a.c.y][a.c.x]-board[a.c.y+1][a.c.x]) <= 1 {
			a.insert(0, 1)
		}
	}

	return fmt.Sprint(a.c.cost)
}

func p2(f io.Reader) string {
	board := parse(f)

	a := &asterisk{}
	for y := range board {
		for x := range board[y] {
			if board[y][x] == 'S' {
				board[y][x] = 'a'
			}
			if board[y][x] == 'E' {
				board[y][x] = 'z'
				a.end.x = x
				a.end.y = y
			}

			if board[y][x] == 'a' {
				a.open = append(a.open, &position{point: point{x: x, y: y}, cost: 0})
			}
		}
	}

	for a.next() {
		if a.c.x > 0 && abs(board[a.c.y][a.c.x]-board[a.c.y][a.c.x-1]) <= 1 {
			a.insert(-1, 0)
		}
		if a.c.x < len(board[a.c.y])-1 && abs(board[a.c.y][a.c.x]-board[a.c.y][a.c.x+1]) <= 1 {
			a.insert(1, 0)
		}
		if a.c.y > 0 && abs(board[a.c.y][a.c.x]-board[a.c.y-1][a.c.x]) <= 1 {
			a.insert(0, -1)
		}
		if a.c.y < len(board)-1 && abs(board[a.c.y][a.c.x]-board[a.c.y+1][a.c.x]) <= 1 {
			a.insert(0, 1)
		}
	}

	return fmt.Sprint(a.c.cost)
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

func abs[N int8 | int](a N) N {
	if a < 0 {
		return -a
	}
	return 1
}
