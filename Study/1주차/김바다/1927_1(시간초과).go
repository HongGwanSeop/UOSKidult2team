/*백준 1927*/

package main
import "fmt"

var heap [100001]int
func insert_min_heap(size, n int){
	size++
	i:=size
	for i!=1 && n < heap[i/2]{
		heap[i] = heap[i/2]
		i/=2
	}
	heap[i]=n
}
func delete_min_heap(size int) int{
	var parent, child, item, temp int
	item = heap[1]
	temp = heap[size]
	size=size-1
	parent = 1
	child =2
	for child<=size{
		if child < size && heap[child]>heap[child+1]{
			child++
		}
		if temp<=heap[child]{
			break
		}
		heap[parent] = heap[child]
		parent=child
		child*=2
	}
	heap[parent]=temp
	return item
}
func main(){
	var C,n,size int
	fmt.Scan(&C)
	size=0
	for i:=0;i<C;i++{
		fmt.Scan(&n)
		if n==0{
			if size==0{
				fmt.Println(0)
			} else{
				//최솟값 출력하기 
				fmt.Println(delete_min_heap(size))
				size--
			}
		} else{
			insert_min_heap(size,n)
			size++
		}
		//fmt.Println(heap)
	}	
}