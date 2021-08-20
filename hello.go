package main

import "fmt"

func main() {
	fmt.Println("This is Sudhir")
	foo()
	//declare and assign a value can be done using :=
	// the value can be changed of a declared value using equal sign =
	x := 42

	fmt.Println(x)

	x = 99
	fmt.Println(x)
	y := 100 + 24
	fmt.Println(y)
	check()
}

func foo() {
	fmt.Print("I am in foo")
}

func check() {
	n, _ := fmt.Println(37, true)
	fmt.Println(n)
	//fmt.Println(err)
	//n returns number of bytes and err returns if there are any errors
	//can pass boolen and integer without ""

	//_ is given instead of err in case we dont want to declare a variable because there should be no unused variables
}
