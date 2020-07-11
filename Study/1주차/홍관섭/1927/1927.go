package main

//timeout됨
import (
	"bufio"
	"os"
	"strconv"
)

var arr [200100]int
var heaplen int = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var num int
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	num, _ = strconv.Atoi(scanner.Text())
	min := bufio.NewWriter(os.Stdout)
	var h string

	var input int
	for i := 1; i <= num; i++ {
		scanner.Scan()
		input, _ = strconv.Atoi(scanner.Text())
		if input == 0 {
			h = strconv.Itoa((Heappop()))
			min.WriteString(h + "\n")
			min.Flush()
		} else {
			Heapin(input)
		}
	}

}
func Heapin(dataa int) {
	heaplen++
	arr[heaplen] = dataa
	ind := heaplen
	var parent int
	parent = ind >> 1
	temp := 0
	for {
		if arr[ind] < arr[parent] && ind != 0 {
			temp = arr[ind]
			arr[ind] = arr[parent]
			arr[parent] = temp
			ind = parent
			parent = ind >> 1
		} else {
			break
		}
	}
}

func Heappop() int {
	if heaplen == 0 {
		return 0
	} else {
		temp := arr[1]
		arr[1] = arr[heaplen]
		arr[heaplen] = 0
		heaplen--
		var temp2 int = 0
		parent := 1
		var child, left, right int
		left = 2
		right = 3
		for left <= heaplen {
			if heaplen == left {
				child = left
			} else {
				if arr[left] < arr[right] {
					child = left
				} else {
					child = right
				}
			}
			if arr[child] < arr[parent] {
				temp2 = arr[child]
				arr[child] = arr[parent]
				arr[parent] = temp2
				parent = child
				left = parent << 1
				right = left + 1
			} else {
				break
			}
		}
		return temp
	}
}
