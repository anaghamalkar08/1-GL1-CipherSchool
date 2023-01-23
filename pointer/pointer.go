package main

import "fmt"

func main() {
	//i:=10
	//j:=&i
	//*j=100
	// var i int
	// i = 10
	// var j *int // declaration
	//j:=new(int) ??this line will declare and assign(j is not empty)

	// j = &i
	// *j = 100
	// fmt.Println(*j)
	// fmt.Println(i)

	name := new(string)
	*name = "golang"
	fmt.Println(*name)

}
