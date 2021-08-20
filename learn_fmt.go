package main

import "fmt"

var y = 45

func main() {
	fmt.Println(y)
	//refer to https://pkg.go.dev/fmt
	fmt.Printf(" %d this is a printf statemnet %v , %+v\n", y, y, y)
	fmt.Printf(" %d this is a printf statemnet %#v , %+v %T\n", y, y, y, y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%O\n", y)
	fmt.Printf("%U\n", y)
}
