package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func makeSliceI(n int) []int {
	return make([]int, n)
}

func scanI() int {
	var v int
	fmt.Fscan(reader, &v)
	return v
}

func scanSliceI(v []int) []int {
	for i := 0; i < len(v); i++ {
		fmt.Fscan(reader, &v[i])
	}
	return v
}

func scanSliceWithMakeI(n int) []int {
	return scanSliceI(makeSliceI(n))
}

func makeSliceS(n int) []string {
	return make([]string, n)
}

func scanS() string {
	var v string
	fmt.Fscan(reader, &v)
	return v
}

func scanSliceS(v []string) []string {
	for i := 0; i < len(v); i++ {
		fmt.Fscan(reader, &v[i])
	}
	return v
}

func scanSliceWithMakeS(n int) []string {
	return scanSliceS(makeSliceS(n))
}

func printV(v ...interface{}) {
	fmt.Fprintln(writer, v...)
}

func printIF(b bool, t interface{}, e interface{}) {
	if b {
		printV(t)
	} else {
		printV(e)
	}
}

func isEven(i int) bool {
	return i%2 == 0
}

// expect argument is '0' ~ '9'
func rtoi(r rune) int {
	return int(r - '0')
}

func absI(i int) int {
	return int(math.Abs(float64(i)))
}

func appendFrontI(i int, v []int) []int {
	return append([]int{i}, v...)
}

func main() {
	defer writer.Flush()
}
