package main

import "fmt"

func main() {
	s := "THis is a string"
	bs := []byte(s)
	//we can convert the values into binary, hexa, oct
	fmt.Println(s)
	fmt.Println(bs)

	fmt.Printf("%b\t\t%#U\t\t%T\n", s, s, s)
	fmt.Printf("%b\t\t%#U\t\t%T\n", bs, bs, bs)
}
