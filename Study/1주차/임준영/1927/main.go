package main

import (
	"bufio"
	"fmt"
	"os"
)

type Heap struct {
	Array []int
	Size  int
}

func (h *Heap) push(i int) {
	c := h.Size      // child node
	p := (c - 1) / 2 // parent node

	h.Array = append(h.Array, i)
	h.Size++

	for {
		if c == 0 || h.Array[c] >= h.Array[p] {
			break
		}
		h.Array[c], h.Array[p] = h.Array[p], h.Array[c]
		c, p = p, (p-1)/2
	}
}

func (h *Heap) pop() int {
	// size := h.Size // child node
	// p := (c - 1) / 2 // parent node
	if h.Size == 0 {
		return 0
	}

	top := h.Array[0]
	h.Size--
	size := h.Size
	h.Array[0] = h.Array[size]
	h.Array = h.Array[:size]

	node := 0
	for {
		left, right := node*2+1, node*2+2
		if left >= size {
			break
		} else if left+1 == size {
			if h.Array[node] > h.Array[left] {
				h.Array[node], h.Array[left] = h.Array[left], h.Array[node]
			}
			break
		}
		if h.Array[node] <= h.Array[left] && h.Array[node] <= h.Array[right] {
			break
		} else if h.Array[node] > h.Array[left] && h.Array[node] > h.Array[right] {
			if h.Array[left] < h.Array[right] {
				h.Array[node], h.Array[left] = h.Array[left], h.Array[node]
				node = left
			} else {
				h.Array[node], h.Array[right] = h.Array[right], h.Array[node]
				node = right
			}
		} else if h.Array[node] > h.Array[left] {
			h.Array[node], h.Array[left] = h.Array[left], h.Array[node]
			node = left
		} else {
			h.Array[node], h.Array[right] = h.Array[right], h.Array[node]
			node = right
		}
	}

	return top
}

func main() {
	var h Heap
	var n, t int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d ", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &t)
		if t == 0 {
			fmt.Println(h.pop())
		} else {
			h.push(t)
		}
	}
}
