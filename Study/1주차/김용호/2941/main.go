package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var input string
	var croatian = []string{"c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z="}

	rd := bufio.NewReader(os.Stdin)
	fmt.Fscanln(rd, &input)

	for i := 0; i < len(croatian); i++ {
		if strings.Contains(input, croatian[i]) {
			input = strings.Replace(input, croatian[i], "0", -1)
		}
	}
	fmt.Println(len(input))
}
