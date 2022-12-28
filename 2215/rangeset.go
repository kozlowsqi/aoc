package main

type set struct {
	min, max int
}

type rangeset struct {
	sets []set
}

func (rs *rangeset) add(a set) {

	var sets []set
	for _, b := range rs.sets {
		if a.min > b.max || b.min > a.max {
			sets = append(sets, b)
			continue
		}

		a.min = min(a.min, b.min)
		a.max = max(a.max, b.max)
	}

	rs.sets = sort(append(sets, a))
}

func (rs *rangeset) sum() int {
	var sum int
	for _, b := range rs.sets {
		sum += b.max - b.min
	}

	return sum
}

func (rs *rangeset) clip(a set) {

	var sets []set
	for _, b := range rs.sets {
		if a.min <= b.max && b.min <= a.max {
			sets = append(sets, set{min: max(a.min, b.min), max: min(a.max, b.max)})
		}
	}

	rs.sets = sets
}

func sort(rs []set) []set {
	if len(rs) == 1 {
		return rs
	}

	h := sort(rs[:len(rs)/2])
	t := sort(rs[len(rs)/2:])

	var srs []set
	for len(h) != 0 && len(t) != 0 {
		if h[0].min < t[0].min {
			srs = append(srs, h[0])
			h = h[1:]
		} else {
			srs = append(srs, t[0])
			t = t[1:]
		}
	}

	srs = append(srs, h...)
	srs = append(srs, t...)

	return srs
}
