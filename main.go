package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

// ---------- scan ----------
func scanI() int {
	var v int
	fmt.Fscan(r, &v)
	return v
}

func scanMakeSliceI(n int) []int {
	arr := make([]int, n)
	scanSliceI(arr)
	return arr
}

func scanIPs(is ...*int) {
	for i := 0; i < len(is); i++ {
		*is[i] = scanI()
	}
}

func scanSliceI(v []int) {
	for i := 0; i < len(v); i++ {
		fmt.Fscan(r, &v[i])
	}
}

func scanVectorI(length int, vv ...[]int) {
	for row := 0; row < length; row++ {
		for column := 0; column < len(vv); column++ {
			scanIPs(&vv[column][row])
		}
	}
}

func scanS() string {
	var v string
	fmt.Fscan(r, &v)
	return v
}

func scanMakeSliceS(n int) []string {
	return scanSliceS(make([]string, n))
}

func scanSPs(ss ...*string) {
	for i := 0; i < len(ss); i++ {
		*ss[i] = scanS()
	}
}

func scanSliceS(v []string) []string {
	for i := 0; i < len(v); i++ {
		fmt.Fscan(r, &v[i])
	}
	return v
}

func scanVectorS(length int, vv ...[]string) {
	for row := 0; row < length; row++ {
		for column := 0; column < len(vv); column++ {
			scanSPs(&vv[column][row])
		}
	}
}

// ---------- print ----------
func printV(v ...interface{}) {
	fmt.Fprintln(w, v...)
}

func printF(format string, v ...interface{}) {
	fmt.Fprintf(w, format, v...)
}

func printDebugV(v ...interface{}) {
	format := ""
	for i := 0; i < len(v); i++ {
		format += "%+v "
	}
	format += "\n"
	fmt.Fprintf(w, format, v...)
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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func containI(target int, array []int) bool {
	for _, e := range array {
		if target == e {
			return true
		}
	}
	return false
}

func containS(target string, array []string) bool {
	for _, e := range array {
		if target == e {
			return true
		}
	}
	return false
}

func containR(target rune, array []rune) bool {
	for _, e := range array {
		if target == e {
			return true
		}
	}
	return false
}

func buildAdjacencyMatrixI(size int, a, b []int, isDirect bool) [][]int {
	size += 1
	vec := make([][]int, size)
	for i := 0; i < size; i++ {
		vec[i] = make([]int, size)
	}

	for i := 0; i < len(a); i++ {
		from, to := a[i], b[i]
		vec[from][to] = 1
		if !isDirect {
			vec[to][from] = 1
		}
	}
	return vec
}

func main() {
	defer w.Flush()
}
