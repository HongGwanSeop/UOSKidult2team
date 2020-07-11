/*백준 10816 숫자카드 2 
현재 시간초과. 이분탐색으로 진행해야 한다*/

package main
import "fmt"

func main(){
	var m,n int
	fmt.Scan(&n)
	var cards [500001]int
	var targets [500001]int
	for i:=0;i<n;i++{
		fmt.Scan(&cards[i])
	}
	fmt.Scan(&m)
	for i:=0;i<m;i++{
		fmt.Scan(&targets[i])
	}
	/*
	fmt.Println(n,m)
	for i:=0;i<n;i++{
		fmt.Print(cards[i]," ")
	}*/
	var count [500001]int
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if targets[i]==cards[j]{
				count[i]++
			}
		}
	}
	for i:=0;i<m;i++{
		fmt.Print(count[i]," ")
	}
		
}