/*백준 10816 숫자카드 2 */

package main
import (
	"fmt"
	"sort"
)

func lower_binary(arr []int, target, size int) int{
	var mid, start, end int
	end = size-1

	for end>start{
		mid = (start+end)/2
		if arr[mid] >= target {
			end = mid
		} else {
			start = mid+1
		}
	}
	return end
}
func upper_binary(arr []int, target, size int) int{
	var mid, start, end int
	end = size-1

	for end>start{
		mid = (start+end)/2
		if arr[mid] > target {
			end = mid
		} else {
			start = mid+1
		}
	}
	return end
}
func main(){
	var m,n,temp,lower,upper int
	fmt.Scan(&n)
	cards := make([]int,0)
	targets := make([]int,0)
	var count [500001]int
	for i:=0;i<n;i++{
		fmt.Scan(&temp)
		cards=append(cards,temp)
	}
	fmt.Scan(&m)
	for i:=0;i<m;i++{
		fmt.Scan(&temp)
		targets=append(targets,temp)
	}
	sort.Sort(sort.IntSlice(cards))
	//sort.Sort(sort.IntSlice(targets))
	for i:=0;i<m;i++{
		lower = lower_binary(cards, targets[i], n)
		upper = upper_binary(cards, targets[i], n)
		if upper == n-1 && cards[n-1]==targets[i]{
			upper++
		}
		count[i] = upper - lower
	}
	for i:=0;i<m;i++{
		fmt.Print(count[i]," ")
	}	
}