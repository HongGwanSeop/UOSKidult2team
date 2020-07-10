/*백준 2941 크로아티아 알파벳*/
package main

import "fmt"

func main() {
	var s string
	num := 0
	fmt.Scanln(&s)
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			switch s[i] {
			case 'c':
				if s[i+1] == '=' || s[i+1] == '-' {
					i++
				}
			case 'd':
				if s[i+1] == '-' {
					i++
				} else {
					if i+2 < len(s) {
						if s[i+1] == 'z' && s[i+2] == '=' {
							i = i + 2
						}
					}
				}
			case 'l', 'n':
				if s[i+1] == 'j' {
					i++
				}
			case 's', 'z':
				if s[i+1] == '=' {
					i++
				}
			}
		}
		num++
	}
	fmt.Print(num)

}
