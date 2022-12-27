package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"strings"
)

//go:embed *.input
var input embed.FS

type point struct {
	x, y int
}

var zero = point{x: 0, y: 0}

func p1(f io.Reader) string {
	fb := bufio.NewReader(f)

	points := map[point]struct{}{}
	for {
		line, err := fb.ReadString('\n')
		if err == io.EOF {
			break
		}
		die(err)

		pointsRaw := strings.Split(strings.TrimSpace(line), " -> ")

		var p0, p1 point
		for _, pointRaw := range pointsRaw {
			p0 = p1

			_, err := fmt.Sscanf(pointRaw, "%d,%d", &p1.x, &p1.y)
			die(err)

			if p0 == zero {
				continue
			}

			points[p1] = struct{}{}

			if p0.y == p1.y {
				var dx int
				if p0.x > p1.x {
					dx = -1
				} else {
					dx = 1
				}

				for p := p0; p != p1; p.x += dx {
					points[p] = struct{}{}
				}
			}

			if p0.x == p1.x {
				var dy int
				if p0.y > p1.y {
					dy = -1
				} else {
					dy = 1
				}

				for p := p0; p != p1; p.y += dy {
					points[p] = struct{}{}
				}
			}
		}
	}

	var mY int // point of no return
	for p := range points {
		if p.y > mY {
			mY = p.y
		}
	}

	for i := 0; ; i++ {
		p := point{x: 500, y: 0}

		for {
			if p.y > mY {
				return fmt.Sprint(i)
			}

			p0 := point{x: p.x, y: p.y + 1}
			if _, ok := points[p0]; !ok {
				p = p0
				continue
			}

			p0 = point{x: p.x - 1, y: p.y + 1}
			if _, ok := points[p0]; !ok {
				p = p0
				continue
			}

			p0 = point{x: p.x + 1, y: p.y + 1}
			if _, ok := points[p0]; !ok {
				p = p0
				continue
			}

			points[p] = struct{}{}
			break
		}
	}
}

func p2(f io.Reader) string {
	fb := bufio.NewReader(f)

	points := map[point]struct{}{}
	for {
		line, err := fb.ReadString('\n')
		if err == io.EOF {
			break
		}
		die(err)

		pointsRaw := strings.Split(strings.TrimSpace(line), " -> ")

		var p0, p1 point
		for _, pointRaw := range pointsRaw {
			p0 = p1

			_, err := fmt.Sscanf(pointRaw, "%d,%d", &p1.x, &p1.y)
			die(err)

			if p0 == zero {
				continue
			}

			points[p1] = struct{}{}

			if p0.y == p1.y {
				var dx int
				if p0.x > p1.x {
					dx = -1
				} else {
					dx = 1
				}

				for p := p0; p != p1; p.x += dx {
					points[p] = struct{}{}
				}
			}

			if p0.x == p1.x {
				var dy int
				if p0.y > p1.y {
					dy = -1
				} else {
					dy = 1
				}

				for p := p0; p != p1; p.y += dy {
					points[p] = struct{}{}
				}
			}
		}
	}

	var mY int // point of no return
	for p := range points {
		if p.y > mY {
			mY = p.y
		}
	}

	for i := 0; ; i++ {
		p := point{x: 500, y: 0}
		if _, ok := points[p]; ok {
			return fmt.Sprint(i)
		}

		for {
			if p.y == mY+1 {
				points[p] = struct{}{}
				break
			}

			p0 := point{x: p.x, y: p.y + 1}
			if _, ok := points[p0]; !ok {
				p = p0
				continue
			}

			p0 = point{x: p.x - 1, y: p.y + 1}
			if _, ok := points[p0]; !ok {
				p = p0
				continue
			}

			p0 = point{x: p.x + 1, y: p.y + 1}
			if _, ok := points[p0]; !ok {
				p = p0
				continue
			}

			points[p] = struct{}{}
			break
		}
	}
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
