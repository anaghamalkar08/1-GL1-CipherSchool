package main

import "fmt"

func main() {
	number := 10
	fmt.Println(number)

	var getInt func(int) int
	getInt = func(x int) int {
		fmt.Println("Is a function")
		return 20 + x
	}
	getInt(1)
	// getInt = func(x int) int {
	// 	fmt.Println("IOn a second function")
	// 	return 10 + x
	// }

	j := func(x int) int {
		fmt.Println("In a first function")
		return 20 + x
	}(19)
	fmt.Println(j)
}

func g(getInt func(int) int, u int) {
	j := getInt(78)
	fmt.Println(j)
}

//function is a first class member i.e we can use function as a type
