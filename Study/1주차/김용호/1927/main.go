package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Intheap []int

func (h Intheap) Len() int {
	return len(h)
}

func (h Intheap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Intheap) Swap(i, j int) {
	if len(h) > 1 {
		h[i], h[j] = h[j], h[i]
	}
}

func (h *Intheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Intheap) Pop() interface{} {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	h := &Intheap{}
	heap.Init(h)

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	for i := 0; i < N; i++ {

		sc.Scan()
		x, _ := strconv.Atoi(sc.Text())
		if x == 0 {
			if h.Len() == 0 {
				fmt.Fprintln(wr, 0)
			} else {
				fmt.Fprintln(wr, heap.Pop(h))
			}
		} else {
			heap.Push(h, x)
		}
	}
	wr.Flush()
}
