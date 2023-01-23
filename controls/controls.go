package main

import "fmt"

//if else

func main() {
	fmt.Println("Enter a number:")
	var input int
	fmt.Scanln(&input)

	if input%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if x := 10; x%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	//Switch case
	a := 100
	switch a {
	case 100:
		a = 23
		fmt.Println(a)
		fallthrough
	case 2:
		a = 25
		fmt.Println(a)
	case 3:
		a = 27
		fmt.Println(a)
	default:
		b := 2
		fmt.Println(b)

	}

}
