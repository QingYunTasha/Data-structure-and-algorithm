package heap

type Element interface {
	Less(e Element) bool
}

type Heap []Element

func NewHeap(e Element) Heap { return Heap{e} }

func (h Heap) Len() int { return len(h) - 1 }

func (h Heap) swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(e Element) {
	*h = append(*h, e)
	h.up(len(*h) - 1)
}

func (h *Heap) Pop() Element {
	n := len(*h) - 1
	res := (*h)[1]
	(*h)[1] = (*h)[n]
	*h = (*h)[:n]
	h.down(1)
	return res
}

func (h Heap) up(i int) {
	for j := i >> 1; i > 1 && h[j].Less(h[i]); i, j = j, j>>1 {
		h.swap(i, j)
	}
}

func (h Heap) down(i int) {
	for j := i << 1; j < len(h); i, j = j, j<<1 {
		if j+1 < len(h) && h[j].Less(h[j+1]) {
			j++
		}
		if h[j].Less(h[i]) {
			break
		}
		h.swap(i, j)
	}
}
