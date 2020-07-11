package main

import "fmt"

func main() {
	var inp string
	cnt := 0
	fmt.Scanln(&inp)
	for i := 0; i < len(inp); i++ {
		if inp[i] == ('c') && (i+1 < len(inp)) {
			if inp[i+1] == '=' {
				cnt++
				i++
			} else if inp[i+1] == '-' {
				cnt++
				i++
			} else {
				cnt++
			}
		} else if inp[i] == 'd' && (i+1 < len(inp)) {
			if i+2 < len(inp) {
				if inp[i+1] == '-' {
					cnt++
					i++
				} else if inp[i+1] == 'z' && inp[i+2] == '=' {
					cnt++
					i += 2
				} else {
					cnt++
				}
			} else if inp[i+1] == '-' {
				cnt++
				i++
			} else {
				cnt++
			}
		} else if (inp[i] == 'l' || inp[i] == 'n') && (i+1 < len(inp)) {
			if inp[i+1] == 'j' {
				cnt++
				i++
			} else {
				cnt++
			}
		} else if (inp[i] == 'z' || inp[i] == 's') && (i+1 < len(inp)) {
			if inp[i+1] == '=' {
				cnt++
				i++
			} else {
				cnt++
			}
		} else {
			cnt++
		}
	}
	fmt.Println(cnt)
}
