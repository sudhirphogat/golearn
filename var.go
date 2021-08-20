package main

import "fmt"

var y = 43 // scope of this variable is entire program (package level)

//** limit the scope of the variable so declare insid the function to avoid confussion
var z int

// this will assign 0 value to Z if we do not assign any values

func main() {
	x := 42 // := shout declaration needs to be printed inside the curly braces
	fmt.Println(x)

	//var y = 43 //var can be declared ouside the function also
	fmt.Println(y)

	food()
	//fmt.Println(z) // No error will be thrown in case I do not use the var z
}

//example of var scope
func food() {
	fmt.Println("again:", y)
	fmt.Printf("%T", y)
}
