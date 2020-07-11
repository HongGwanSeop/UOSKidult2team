/*백준 10816 숫자카드 2 */

package main
import (
	"fmt"
	"sort"
)

func main(){
	var m,n int
	fmt.Scan(&n)
	cards := make([]int,0)
	targets := make([]int,0)
	var temp int
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
	
	var left,right int
	index:=0
	for true{
		//fmt.Println("targets[index]: ",targets[index])
		flag := true
		left=0
		right=len(cards)-1
		for left<=right{
			mid:=(left+right)/2
			if cards[mid] > targets[index]{
				right = mid-1
			}else if cards[mid]<targets[index] {
				left=mid+1
			}else{ //같으면 (탐색성공)
				flag = false
				//fmt.Println(cards[mid],targets[index])
				count[index]++
				if mid<n-1{
					cards = append(cards[:mid],cards[mid+1:]...) //i번째 요소 삭제
				} else {
					cards = append(cards[:mid])
				}
				break
			}
		}
		if flag && index==m-1{
			break
		}
		if flag{
			index++
		}
		
		//fmt.Println(cards)
	}
	for i:=0;i<m;i++{
		fmt.Print(count[i]," ")
	}
	
		
}