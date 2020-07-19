package main

import (
	"bufio"
	"fmt"
	"os"
)

var arr [][]int
var n int
var safe int = 1

func rainB(arr [][]bool) (safe int) {
	max := len(arr) - 1
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			if !arr[i][j] {
				continue
			}
			if arr[i][j+1] {
				if arr[i+1][j] {
					if arr[i+1][j+1] {
						continue
					} else {
						safe -= 1
						continue
					}
				} else {
					continue
				}
			} else if arr[i+1][j] {
				continue
			} else {
				safe += 1
			}
		}
	}
	return safe
}

func rain(arr [][]int, n int) (safe int) {
	max := len(arr)
	cp := make([][]bool, max+1)
	for i := 0; i < max; i++ {
		cp[i] = make([]bool, max+1)
		for j := 0; j < max; j++ {
			if arr[i][j] > n {
				cp[i][j] = true
			}
		}
	}
	cp[max] = make([]bool, max+1)
	safe = rainB(cp)
	return safe
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	arr = make([][]int, n)
	max, min := 1, 100
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &arr[i][j])
			t := arr[i][j]
			if t < min {
				min = t
			}
			if t > max {
				max = t
			}
		}
	}

	for i := min; i < max; i++ {
		r := rain(arr, i)
		if safe < r {
			safe = r
		}
	}
	fmt.Println(safe)
}
