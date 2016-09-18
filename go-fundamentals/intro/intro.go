// go doc builtin
/*
package builtin // import "builtin"

Package builtin provides documentation for Go's predeclared identifiers. The
items documented here are not actually in package builtin but their
descriptions here allow godoc to present documentation for the language's
special identifiers.

const true = 0 == 0 ...
const iota = 0
var nil Type
func append(slice []Type, elems ...Type) []Type
func close(c chan<- Type)
func delete(m map[Type]Type1, key Type)
func panic(v interface{})
func print(args ...Type)
func println(args ...Type)
func recover() interface{}
func complex(r, i FloatType) ComplexType
func imag(c ComplexType) FloatType
func real(c ComplexType) FloatType
func make(Type, size IntegerType) Type
func new(Type) *Type
func cap(v Type) int
func copy(dst, src []Type) int
func len(v Type) int
type ComplexType complex64
type FloatType float32
type IntegerType int
type Type int
type Type1 int
type bool bool
type byte byte
type complex128 complex128
type complex64 complex64
type error interface { ... }
type float32 float32
type float64 float64
type int int
type int16 int16
type int32 int32
type int64 int64
type int8 int8
type rune rune
type string string
type uint uint
type uint16 uint16
type uint32 uint32
type uint64 uint64
type uint8 uint8
type uintptr uintptr
*/
// use "go run" to run this go file.
// also you can "go build" this go file and run the result file.
package main

import "fmt"

func myfunc() {
	i := 0
Here:
	fmt.Printf("%d\n", i)
	i++
	if i < 10 {
		goto Here
	}
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func main() {
	// define variables

	var a int
	fmt.Printf("a=%d\n", a)

	b := 1 // can only be used inside a function.
	fmt.Printf("b=%d\n", b)

	// multiple declarations
	var (
		c int
		d string
	)

	fmt.Printf("c=%d d=%s\n", c, d)

	e, f := 1, 2
	fmt.Printf("e=%d f=%d\n", e, f)

	_, g := 3, 4 // _ will be dropped.
	fmt.Printf("g=%d\n", g)

	var h int64
	fmt.Printf("h=%d\n", h)

	const x = 1
	fmt.Printf("print constant x=%d\n", x)

	const (
		y = iota
		z = iota
	)
	fmt.Printf("wtf is iota? y=%d z=%d\n", y, z)

	const (
		s  string = "hello"
		a1        = 1
	)
	fmt.Printf("s=%s a1=%d\n", s, a1)

	s1 := "hello"
	c1 := []rune(s1)
	c1[0] = 'c'
	s2 := string(c1)
	fmt.Printf("s2=%s\n", s2)

	msg := "hello" +
		" my friend"

	fmt.Printf("msg=%s\n", msg)

	msg = `hello my 
	ok, ok
	workd`
	fmt.Printf("msg=%s\n", msg)

	// rune is alias of int32

	var comp complex64 = 5 + 5i
	fmt.Printf("comp=%v\n", comp)

	// if else
	w := 10
	if w > 0 {
		fmt.Printf("w > 0\n")
	} else {
		fmt.Printf("w <= 0\n")
	}

	if w = 10; w > 5 {
		fmt.Printf("w > 5\n")
	} else {
		fmt.Printf("w <= 5\n")
	}

	if true && true {
		fmt.Printf("it should be here.\n")
	}

	if true || false { //left brace should be in the same line with if.
		fmt.Printf("it should be here.\n")
	}

	myfunc()

	for j := 0; j < 10; j++ {
		fmt.Printf("j=%d\n", j)
	}

	j := 0
	for j < 10 {
		fmt.Printf("j=%d\n", j)
		j++
	}

	j = 0
	for {
		fmt.Printf("j = %d\n", j)
		j++
		if j > 10 {
			break
		}
	}

	list := []string{"a", "b", "c", "d"}
	for k, v := range list {
		fmt.Printf("%d->%s\n", k, v)
	}

	for pos, char := range "hello" {
		fmt.Printf("%d->%c\n", pos, char)
	}

	res := unhex('8')
	fmt.Printf("res = %d\n", res)

	// fallthrough
	i := 0
	switch i {
	case 0:
		fallthrough //fallthrough make case connected.
	case 1:
		fmt.Printf("I'm in case 1.\n")
	default:
		fmt.Printf("I'm in default.\n")
	}

	i = 1
	switch i {
	case 1, 2, 3, 4:
		fmt.Printf("Ok, you are in my list.\n")
	}

	println("hello, println.\n")

	// bubble sort
	arr := []int{1, 3, 2, 4}
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	for _, v := range arr {
		println(v)
	}

	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = 10 - i
	}

	for _, v := range array {
		println(v)
	}

	//向函数传递数组，是值拷贝，而不是引用拷贝。。。。
	arr1 := [...]int{1, 2, 3, 4}
	for _, v := range arr1 {
		println(v)
	}

	sl := make([]int, 10)
	for i := 0; i < len(sl); i++ {
		println(sl[i])
	}

	sl2 := arr1[2:3]
	println(len(sl2))
	println(cap(sl2))

	t0 := arr1[:]
	t1 := append(t0, 2)
	t2 := append(t1, 3, 4, 5)

	t3 := append(t2, t0...)

	for _, v := range t3 {
		println(v)
	}

	// test copy
	var y0 = []int{0, 2, 3}
	var y1 = make([]int, 6)

	_ = copy(y1, y0[0:])
	fmt.Printf("y2:\n")
	for _, v := range y1 {
		println(v)
	}

	// test map
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, // comma is necessary
	}

	// calculate how many days in a year
	days := 0
	for _, n := range monthdays {
		days += n
	}
	fmt.Printf("days=%d\n", days)

	// check if key exists
	var value int
	var present bool
	value, present = monthdays["Jan"]
	v, ok := monthdays["Ok"]

	if ok == false {
		println("Good!")
	} else {
		println(v)
	}
	if present {
		println("Good!")
		println(value)
	}

	// delete a key/value in a map
	delete(monthdays, "Jan")
}
