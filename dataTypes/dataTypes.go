package main

import "fmt"

func main() {
	var data int
	data = 10
	data1 := 100
	data = 1000
	fmt.Println(data)
	fmt.Println(data1)

	//data types
	//a) Primitive
	//int,float,bool,unit,string,complex, byte, rune
	//b) Non-Primitive
	// struct, map, chan, slice, interface, array
	name := "anagha"
	fmt.Println(name)

	const pi = 3.14 // implicit
	//const pi float32=3.14 -> explicit
	fmt.Println(pi)

}
