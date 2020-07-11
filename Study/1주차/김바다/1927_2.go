/*백준 1927 ver1에서 입출력방식 바꾸고, heap을 go에서 제공하는 것으로 바꿈*/

package main
import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type IntHeap []int
 
func (h IntHeap) Len() int {
    return len(h)
}
 
func (h IntHeap) Less(i, j int) bool {
    return h[i] < h[j]
}
 
func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
 
func (h *IntHeap) Push(element interface{}) {
    *h = append(*h, element.(int))
}
 
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    element := old[n-1]
    *h = old[0 : n-1]
    return element
}
func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	h := &IntHeap{}
	heap.Init(h)

	scanner.Scan()
	C, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < C; i++ {

		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			if h.Len() == 0 {
				fmt.Fprintln(writer, 0)
			} else {
				fmt.Fprintln(writer, heap.Pop(h))
			}
		} else {
			heap.Push(h, n)
		}
	}
	writer.Flush()
}