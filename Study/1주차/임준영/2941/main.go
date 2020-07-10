package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string
	fmt.Scan(&s)
	letters := []string{"c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z="}
	for i := 0; i < len(letters); i++ {
		r, _ := regexp.Compile(letters[i])
		s = r.ReplaceAllString(s, "a")
	}
	fmt.Println(len(s))
}
