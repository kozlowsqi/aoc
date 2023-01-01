package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type valve struct {
	label string
	frate int
	costs map[string]int
}

func (v valve) String() string {
	var str string

	str = fmt.Sprintf("%s(%d):\t", v.label, v.frate)

	for k, v := range v.costs {
		str = fmt.Sprintf("%s %s=%d", str, k, v)
	}

	return str
}

func findcost(src, dest string, conns map[string][]string) int {
	open, closed := []string{src}, map[string]struct{}{}

	for cost := 1; ; cost++ {
		for _, a := range open {
			if a == dest {
				return cost
			}

			closed[a] = struct{}{}
			open = open[1:]

			for _, b := range conns[a] {
				if _, ok := closed[b]; ok {
					continue
				}

				open = append(open, b)
			}
		}
	}
}

func parse(f io.Reader) map[string]*valve {
	fb := bufio.NewReader(f)

	valves, conns := map[string]*valve{}, map[string][]string{}
	for {
		l, err := fb.ReadString('\n')
		if errors.Is(err, io.EOF) {
			break
		}
		die(err)

		l = strings.TrimSpace(l)
		if len(l) == 0 {
			break
		}

		var v valve
		_, err = fmt.Sscanf(l, "Valve %s has flow rate=%d;", &v.label, &v.frate)
		die(err)

		if i := strings.Index(l, "valves"); i != -1 {
			conns[v.label] = strings.Split(l[i+7:], ", ")
		} else if i := strings.Index(l, "valve"); i != -1 {
			conns[v.label] = []string{l[i+6:]}
		}

		if v.frate > 0 || v.label == "AA" {
			v.costs = map[string]int{}
			valves[v.label] = &v
		}
	}

	for src := range valves {
		for dest := range valves {
			if src == dest {
				continue
			}
			valves[src].costs[dest] = findcost(src, dest, conns)
		}
	}

	return valves
}
