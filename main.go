package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"strconv"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

var D4 = [4]Pair{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var D4P = func(p Pair) [4]Pair {
	return [4]Pair{
		{p.y + D4[0].y, p.x + D4[0].x},
		{p.y + D4[1].y, p.x + D4[1].x},
		{p.y + D4[2].y, p.x + D4[2].x},
		{p.y + D4[3].y, p.x + D4[3].x},
	}
}
var D8 = [8]Pair{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
var D8P = func(p Pair) [8]Pair {
	return [8]Pair{
		{p.y + D8[0].y, p.x + D8[0].x},
		{p.y + D8[1].y, p.x + D8[1].x},
		{p.y + D8[2].y, p.x + D8[2].x},
		{p.y + D8[3].y, p.x + D8[3].x},
		{p.y + D8[4].y, p.x + D8[4].x},
		{p.y + D8[5].y, p.x + D8[5].x},
		{p.y + D8[6].y, p.x + D8[6].x},
		{p.y + D8[7].y, p.x + D8[7].x},
	}
}

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

func scanVectorS(rowSize, columnSize int, f func(i, j int, r rune) rune) [][]rune {
	vec := make([][]rune, rowSize)
	for i := 0; i < rowSize; i++ {
		vec[i] = make([]rune, columnSize)
		s := scanS(nil)
		for j, r := range s {
			if f != nil {
				r = f(i, j, r)
			}
			vec[i][j] = r
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

// NOTE: expect argument is 0 ~ 9
func itor(i int) rune {
	return rune('0' + i)
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

func (p *Pair) isInBound(height, width int) bool {
	return p.y >= 0 && p.x >= 0 && p.y < height && p.x < width
}

type IntQueue struct {
	inner *list.List
}

func newIntQueue() IntQueue {
	return IntQueue{inner: list.New()}
}

func (q *IntQueue) popFront() int {
	e := q.inner.Front()
	q.inner.Remove(e)
	return e.Value.(int)
}

func (q *IntQueue) popBack() int {
	e := q.inner.Back()
	q.inner.Remove(e)
	return e.Value.(int)
}

func (q *IntQueue) popAll() []int {
	pairs := make([]int, q.inner.Len())
	l := q.inner.Len()
	for i := 0; i < l; i++ {
		pairs[i] = q.popFront()
	}
	return pairs
}

func (q *IntQueue) pushFront(v int) {
	q.inner.PushFront(v)
}

func (q *IntQueue) pushBack(v int) {
	q.inner.PushBack(v)
}

func (q *IntQueue) isEmpty() bool {
	return q.inner.Len() == 0
}

type PairQueue struct {
	inner *list.List
}

func newPairQueue() PairQueue {
	return PairQueue{inner: list.New()}
}

func (q *PairQueue) popFront() Pair {
	e := q.inner.Front()
	q.inner.Remove(e)
	return e.Value.(Pair)
}

func (q *PairQueue) popBack() Pair {
	e := q.inner.Back()
	q.inner.Remove(e)
	return e.Value.(Pair)
}

func (q *PairQueue) popAll() []Pair {
	pairs := make([]Pair, q.inner.Len())
	l := q.inner.Len()
	for i := 0; i < l; i++ {
		pairs[i] = q.popFront()
	}
	return pairs
}

func (q *PairQueue) pushFront(v Pair) {
	q.inner.PushFront(v)
}

func (q *PairQueue) pushBack(v Pair) {
	q.inner.PushBack(v)
}

func (q *PairQueue) isEmpty() bool {
	return q.inner.Len() == 0
}

func main() {
	defer w.Flush()
}
