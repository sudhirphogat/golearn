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
}

func foo() {
	fmt.Print("I am in foo")
}
