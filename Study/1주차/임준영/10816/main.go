package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var m [20000002]int
	var i, j int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d ", &i)

	for k := 0; k < i; k++ {
		fmt.Fscanf(reader, "%d ", &j)
		m[j+10000001] += 1
	}

	fmt.Fscanf(reader, "%d ", &i)

	for k := 0; k < i; k++ {
		fmt.Fscanf(reader, "%d ", &j)
		fmt.Printf("%d ", m[j+10000001])
	}
}
