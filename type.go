package main

import "fmt"

var a int

//we can define our own type
// the under type is int
type testtype int

var b testtype

func main() {

	a = 42
	b = 43

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

}
