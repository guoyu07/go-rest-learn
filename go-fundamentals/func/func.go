package main

import "fmt"

// implement a recursive function.
func rec(i int) {
	if i == 10 {
		return
	}
	rec(i + 1)
	fmt.Printf("%d ", i)
}

// test defer sentence.
func deferTest() {
	for i := 0; i < 10; i++ {
		defer println(i)
	}

}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 0
}

// test variable length arguments.
func varArgument(arg ...int) {
	for i := 0; i < len(arg); i++ {
		println(arg[i])
	}
}

// test anonymous func
func anonymousFunc() {
	a := func() {
		fmt.Printf("hello, anonymous function.\n")
	}

	a()
}

// callback is simple, don't show it here.

func main() {
	rec(0)
	deferTest()
	println(f())
	varArgument(10, 10, 10)
	anonymousFunc()
}
