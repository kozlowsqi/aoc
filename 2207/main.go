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

type dir struct {
	parent *dir
	files  map[string]int
	dirs   map[string]*dir
}

func newdir(cd *dir) *dir {
	return &dir{
		parent: cd,
		files:  make(map[string]int),
		dirs:   make(map[string]*dir),
	}
}

func (d *dir) msize(f int) int {
	size := d.size()

	if size < f {
		return 0
	}

	for _, v := range d.dirs {
		s := v.msize(f)

		if s >= f && s < size {
			size = s
		}
	}

	return size
}

func (d *dir) dsize() int {
	size := d.size()

	if size > 100000 {
		size = 0
	}

	for _, v := range d.dirs {
		size += v.dsize()
	}

	return size
}

func (d *dir) size() int {
	var size int

	for _, v := range d.files {
		size += v
	}

	for _, v := range d.dirs {
		size += v.size()
	}

	return size
}

func p1(f io.Reader) string {
	fb := bufio.NewReader(f)

	root := newdir(nil)
	cd := root
	for {
		l, err := fb.ReadString('\n')
		if err == io.EOF {
			break
		}
		die(err)

		switch {
		case strings.HasPrefix(l, "$ cd /"):
			break
		case strings.HasPrefix(l, "$ cd .."):
			cd = cd.parent
		case strings.HasPrefix(l, "$ cd "):
			cd = cd.dirs[strings.TrimPrefix(l, "$ cd ")]
		case strings.HasPrefix(l, "$ ls"):
			break
		case strings.HasPrefix(l, "dir "):
			cd.dirs[strings.TrimPrefix(l, "dir ")] = newdir(cd)
		default:
			var n string
			var s int
			_, err := fmt.Sscanf(l, "%d %s", &s, &n)
			if err != nil {
				panic(err)
			}

			cd.files[n] = s
		}
	}

	return fmt.Sprint(root.dsize())
}

func p2(f io.Reader) string {
	fb := bufio.NewReader(f)

	root := newdir(nil)
	cd := root
	for {
		l, err := fb.ReadString('\n')
		if err == io.EOF {
			break
		}
		die(err)

		switch {
		case strings.HasPrefix(l, "$ cd /"):
			break
		case strings.HasPrefix(l, "$ cd .."):
			cd = cd.parent
		case strings.HasPrefix(l, "$ cd "):
			cd = cd.dirs[strings.TrimPrefix(l, "$ cd ")]
		case strings.HasPrefix(l, "$ ls"):
			break
		case strings.HasPrefix(l, "dir "):
			cd.dirs[strings.TrimPrefix(l, "dir ")] = newdir(cd)
		default:
			var n string
			var s int
			_, err := fmt.Sscanf(l, "%d %s", &s, &n)
			if err != nil {
				panic(err)
			}

			cd.files[n] = s
		}
	}

	s := 30000000 - (70000000 - root.size())
	return fmt.Sprint(root.msize(s))
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
