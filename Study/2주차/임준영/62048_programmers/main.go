package main

import "fmt"

func gcd(a int, b int) int {
	if a%b > 0 {
		return gcd(b, a%b)
	} else {
		return b
	}
}

func solution(w int, h int) int64 {
	var answer int64 = int64(w * h)
	x := gcd(w, h)
	wx, hx := w/x, h/x

	answer -= int64((wx + hx - 1) * x)
	return answer
}

func main() {
	fmt.Println(solution(8, 12))
}
