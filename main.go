package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

var COORD4 = [4]Pair{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var COORD8 = [8]Pair{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

// ---------- input handler ----------
var MINUS = func(i int) int {
	return i - 1
}

// ---------- scan ----------
func scanI(f func(i int) int) int {
	var v int
	fmt.Fscan(r, &v)
	if f != nil {
		v = f(v)
	}
	return v
}

func scanMakeSliceI(n int, f func(i int) int) []int {
	arr := make([]int, n)
	scanSliceI(arr, f)
	return arr
}

func scanSliceI(v []int, f func(i int) int) {
	for i := 0; i < len(v); i++ {
		fmt.Fscan(r, &v[i])
		if f != nil {
			v[i] = f(v[i])
		}
	}
}

func scanVectorI(size int, f func(i int) int) ([]int, []int) {
	x, y := make([]int, size), make([]int, size)
	for row := 0; row < size; row++ {
		x[row], y[row] = scanI(f), scanI(f)
	}
	return x, y
}

func scanS(f func(s string) string) string {
	var v string
	fmt.Fscan(r, &v)
	if f != nil {
		return f(v)
	}
	return v
}

func scanMakeSliceS(n int, f func(s string) string) []string {
	return scanSliceS(make([]string, n), f)
}

func scanSliceS(v []string, f func(s string) string) []string {
	for i := 0; i < len(v); i++ {
		fmt.Fscan(r, &v[i])
		if f != nil {
			v[i] = f(v[i])
		}
	}
	return v
}

func scanVectorS(rowSize, columnSize int, f func(s string) string) [][]rune {
	vec := make([][]rune, rowSize)
	for i := 0; i < rowSize; i++ {
		vec[i] = make([]rune, columnSize)
		ss := scanS(f)
		for j, s := range ss {
			vec[i][j] = s
		}
	}
	return vec
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
	printF(format+"\n", v...)
}

func printDebugMetrixI(v [][]int) {
	digit := len(strconv.Itoa(len(v)))
	format := "%" + strconv.Itoa(digit) + "d"

	header := "  "
	for i := 0; i < digit; i++ {
		header += " "
	}

	columnNumber := make([]interface{}, len(v))
	for i := 0; i < len(v); i++ {
		header += format + " "
		columnNumber[i] = i
	}
	printF(header+"\n", columnNumber...)

	for row := 0; row < len(v); row++ {
		printF(format+" [", row)
		for column := 0; column < len(v[row]); column++ {
			printF(format, v[row][column])
			if column < len(v[row])-1 {
				printF(" ")
			}
		}
		printV("]")
	}
}

func printDebugMetrixR(v [][]rune) {
	header := "   "
	columnNumber := make([]interface{}, len(v[0]))
	for i := 0; i < len(v[0]); i++ {
		header += "%d "
		columnNumber[i] = i
	}
	printF(header+"\n", columnNumber...)

	for row := 0; row < len(v); row++ {
		printF("%d [", row)
		for column := 0; column < len(v[row]); column++ {
			printF("%s", string(v[row][column]))
			if column < len(v[row])-1 {
				printF(" ")
			}
		}
		printV("]")
	}
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

func buildAdjacencyMatrixI(size int, srcNodes, dstNodes []int, isDirect bool) [][]int {
	size += 1
	vec := make([][]int, size)
	for i := 0; i < size; i++ {
		vec[i] = make([]int, size)
	}

	for i := 0; i < len(srcNodes); i++ {
		from, to := srcNodes[i], dstNodes[i]
		vec[from][to] = 1
		if !isDirect {
			vec[to][from] = 1
		}
	}
	return vec
}

func copyMatrixI(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i, v := range src {
		dst[i] = make([]int, len(v))
		copy(dst[i], v)
	}
	return dst
}

func copyMatrixR(src [][]rune) [][]rune {
	dst := make([][]rune, len(src))
	for i, v := range src {
		dst[i] = make([]rune, len(v))
		copy(dst[i], v)
	}
	return dst
}

type Pair struct {
	y, x int
}

type QueueInt struct {
	inner []int
}

func newQueueInt(a []int) QueueInt {
	return QueueInt{inner: a}
}

func (q *QueueInt) pop() (*int, error) {
	if len(q.inner) == 0 {
		return nil, errors.New("queue is empty")
	} else if len(q.inner) == 1 {
		var v *int
		v, q.inner = &q.inner[0], []int{}
		return v, nil
	} else {
		var v *int
		v, q.inner = &q.inner[0], q.inner[1:]
		return v, nil
	}
}

func (q *QueueInt) push(v int) {
	q.inner = append(q.inner, v)
}

func (q *QueueInt) isEmpty() bool {
	return len(q.inner) == 0
}

type QueuePair struct {
	inner []Pair
}

func newQueuePair(a []Pair) QueuePair {
	return QueuePair{inner: a}
}

func (q *QueuePair) pop() (*Pair, error) {
	if len(q.inner) == 0 {
		return nil, errors.New("queue is empty")
	} else if len(q.inner) == 1 {
		var v *Pair
		v, q.inner = &q.inner[0], []Pair{}
		return v, nil
	} else {
		var v *Pair
		v, q.inner = &q.inner[0], q.inner[1:]
		return v, nil
	}
}

func (q *QueuePair) push(v Pair) {
	q.inner = append(q.inner, v)
}

func (q *QueuePair) isEmpty() bool {
	return len(q.inner) == 0
}

func main() {
	defer w.Flush()
}
