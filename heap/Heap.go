package heap

type Heap[T comparable] struct {
	Array   []T
	Greater func(a, b T) bool
}

func swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func (h *Heap[T]) root() (_ T) {
	if h.Len() <= 0 {
		return
	}
	return h.Array[0]
}
func (h *Heap[T]) flot() {
	loc := len(h.Array) - 1
	for loc > 0 {
		parent := (loc - 1) / 2
		// fmt.Println(parent, loc)
		if h.Greater(h.Array[parent], h.Array[loc]) {
			break
		}
		swap(&h.Array[parent], &h.Array[loc])
		loc = parent
	}
}
func (h *Heap[T]) sink() {
	loc := 0
	len := len(h.Array)
	for loc < len {
		who := loc
		max := h.Array[loc]
		if left := loc*2 + 1; left < len && h.Greater(h.Array[left], max) {
			who = left
			max = h.Array[left]
		}
		if right := loc*2 + 2; right < len && h.Greater(h.Array[right], max) {
			who = right
			max = h.Array[right]
		}
		if loc == who {
			break
		}
		swap(&h.Array[loc], &h.Array[who])
		loc = who
	}
}
func (h *Heap[T]) Push(e T) {
	h.Array = append(h.Array, e)
	h.flot()
}
func (h *Heap[T]) Pop() (_ T) {
	if h.Len() <= 0 {
		return
	}
	ret := h.Array[0]
	swap(&h.Array[0], &h.Array[len(h.Array)-1])
	h.Array = append(h.Array[:len(h.Array)-1])
	h.sink()
	return ret
}
func (h Heap[T]) Len() int {
	return len(h.Array)
}
