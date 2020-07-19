package main

import (
	"bufio"
	"fmt"
	"os"
)

func lcheck(arr *[][]int, i int, j int, n int) {
	i--
	if i >= 0 {
		if (*arr)[i][j] <= 100 && (*arr)[i][j] > 0 {
			(*arr)[i][j] = 0
			if i > 0 {
				lcheck(arr, i, j, n)
			}
			if j > 0 {
				ucheck(arr, i, j, n)
			}
			if j < n {
				dcheck(arr, i, j, n)
			}
		}
	}
}
func rcheck(arr *[][]int, i int, j int, n int) {
	i++
	if i < n {
		if (*arr)[i][j] <= 100 && (*arr)[i][j] > 0 {
			(*arr)[i][j] = 0
			if i < n-1 {
				rcheck(arr, i, j, n)
			}
			if j > 0 {
				ucheck(arr, i, j, n)
			}
			if j < n {
				dcheck(arr, i, j, n)
			}
		}
	}

}
func ucheck(arr *[][]int, i int, j int, n int) {
	j--
	if j >= 0 {
		if (*arr)[i][j] <= 100 && (*arr)[i][j] > 0 {
			(*arr)[i][j] = 0
			if i > 0 {
				lcheck(arr, i, j, n)
			}
			if i < n-1 {
				rcheck(arr, i, j, n)
			}
			if j > 0 {
				ucheck(arr, i, j, n)
			}
		}
	}
}
func dcheck(arr *[][]int, i int, j int, n int) {
	j++
	if j < n {
		if (*arr)[i][j] <= 100 && (*arr)[i][j] > 0 {
			(*arr)[i][j] = 0
			if i > 0 {
				lcheck(arr, i, j, n)
			}
			if i < n-1 {
				rcheck(arr, i, j, n)
			}
			if j < n-1 {
				dcheck(arr, i, j, n)
			}
		}
	}
}

func check(arr *[][]int, i int, j int, n int) int {
	if (*arr)[i][j] > 100 && (*arr)[i][j] == 0 {
		return 0
	}
	(*arr)[i][j] = 0

	ucheck(arr, i, j, n)
	dcheck(arr, i, j, n)
	rcheck(arr, i, j, n)
	lcheck(arr, i, j, n)
	return 1
}

func cntt(arr [][]int, n int) int {
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] <= 100 && arr[i][j] > 0 {
				cnt += check(&arr, i, j, n)
			}
		}
	}
	return cnt
}

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	array := make([][]int, n)
	for i := range array {
		array[i] = make([]int, n)
	}
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &temp[j])
			array[i][j] = temp[j]
		}
	}
	max := 0
	var cnt int
	temp2 := make([][]int, n)
	for i := range array {
		temp2[i] = make([]int, n)
	}
	for k := 0; k <= 100; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				temp2[i][j] = array[i][j]
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if temp2[i][j] <= k {
					temp2[i][j] += 100
				}
			}
		}
		cnt = cntt(temp2, n)
		if max <= cnt {
			max = cnt
		}

	}
	fmt.Println(max)
}
