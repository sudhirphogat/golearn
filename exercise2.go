package main

import "fmt"

const (
	a = iota
	b
	c
)

func main() {
	s := "THis is a string"
	bs := []byte(s)
	d := 42
	//we can convert the values into binary, hexa, oct
	fmt.Println(s)
	fmt.Println(bs)

	fmt.Printf("%b\t\t%#U\t\t%T\n", s, s, s)
	fmt.Printf("%b\t\t%#U\t\t%T\n", bs, bs, bs)
	//Iota gives an index to each value from where it is declared
	fmt.Println(a, b, c)

	fmt.Printf("%d\t\t%b\t\t%#x\n", d, d, d)

	//to shift 1 decimal to left
	e := d << 1

	fmt.Printf("%d\t\t%b\t\t%#x\n", e, e, e)
}
