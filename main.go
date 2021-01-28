package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

// ---------- scan ----------
func scanI() int {
	var v int
	fmt.Fscan(reader, &v)
	return v
}

func scanMakeSliceI(n int) []int {
	arr := make([]int, n)
	scanSliceI(arr)
	return arr
}

func scanIPs(is ...*int) {
	l := len(is)
	for i := 0; i < l; i++ {
		*is[i] = scanI()
	}
}

func scanSliceI(v []int) {
	for i := 0; i < len(v); i++ {
		fmt.Fscan(reader, &v[i])
	}
}

func scanVectorI(length int, vv ...[]int) {
	columns := len(vv)
	for row := 0; row < length; row++ {
		for column := 0; column < columns; column++ {
			scanIPs(&vv[column][row])
		}
	}
}

func scanS() string {
	var v string
	fmt.Fscan(reader, &v)
	return v
}

func scanMakeSliceS(n int) []string {
	return scanSliceS(make([]string, n))
}

func scanSPs(ss ...*string) {
	l := len(ss)
	for i := 0; i < l; i++ {
		*ss[i] = scanS()
	}
}

func scanSliceS(v []string) []string {
	for i := 0; i < len(v); i++ {
		fmt.Fscan(reader, &v[i])
	}
	return v
}

func scanVectorS(length int, vv ...[]string) {
	columns := len(vv)
	for row := 0; row < length; row++ {
		for column := 0; column < columns; column++ {
			scanSPs(&vv[column][row])
		}
	}
}

// ---------- print ----------
func printV(v ...interface{}) {
	fmt.Fprintln(writer, v...)
}

func printDebugV(v ...interface{}) {
	l := len(v)
	format := ""
	for i := 0; i < l; i++ {
		format += "%+v "
	}
	fmt.Fprintf(writer, format, v...)
}

func printIF(cond bool, t interface{}, e interface{}) {
	if cond {
		printV(t)
	} else {
		printV(e)
	}
}

// ---------- util ----------
func isEven(i int) bool {
	return i%2 == 0
}

// NOTE: expect argument is '0' ~ '9'
func rtoi(r rune) int {
	return int(r - '0')
}

func absI(i int) int {
	return int(math.Abs(float64(i)))
}

func appendFrontI(i int, v []int) []int {
	return append([]int{i}, v...)
}

func appendFrontS(s string, v []string) []string {
	return append([]string{s}, v...)
}

func main() {
	defer writer.Flush()
}
