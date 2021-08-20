package main

import "fmt"

func main() {
	fmt.Println("This is Sudhir")
	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
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
