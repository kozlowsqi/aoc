package main

type heap[V any] func(V, V) bool

type ptree[V any] struct {
	value    V
	subtrees []ptree[V]
}

func (h heap[V]) meld(a, b *ptree[V]) *ptree[V] {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	if h(a.value, b.value) {
		a.subtrees = append([]ptree[V]{*b}, a.subtrees...)
		return a
	} else {
		b.subtrees = append([]ptree[V]{*a}, b.subtrees...)
		return b
	}
}

func (h heap[V]) mergepairs(a []ptree[V]) *ptree[V] {
	if len(a) == 0 {
		return nil
	}

	if len(a) == 1 {
		return &a[0]
	}

	return h.meld(h.meld(&a[0], &a[1]), h.mergepairs(a[2:]))
}

func (h heap[V]) queue(a **ptree[V], v V) {
	*a = h.meld(*a, &ptree[V]{value: v})
}

func (h heap[V]) dequeue(a **ptree[V]) *V {
	if a == nil || *a == nil {
		return nil
	}

	value := (*a).value
	*a = h.mergepairs((*a).subtrees)

	return &value
}
