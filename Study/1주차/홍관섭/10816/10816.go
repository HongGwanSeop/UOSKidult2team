package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var in1, in2 int
	var m map[int]int
	m = make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	in1, _ = strconv.Atoi(scanner.Text())
	var h int = 0
	var val int
	var exists bool
	for i := 0; i < in1; i++ {
		scanner.Scan()
		h, _ = strconv.Atoi(scanner.Text())
		val, exists = m[h]
		if exists {
			val++
			m[h] = val
		} else {
			m[h] = 1
		}
	}
	scanner.Scan()
	in2, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < in2; i++ {
		scanner.Scan()
		h, _ = strconv.Atoi(scanner.Text())
		val, exists = m[h]
		if exists {
			fmt.Fprintf(writer, "%d ", val)
		} else {
			fmt.Fprintf(writer, "%d ", 0)
		}
	}
	writer.Flush()
}
