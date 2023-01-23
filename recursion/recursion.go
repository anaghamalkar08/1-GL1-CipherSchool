package main

import "fmt"

func main() {
	//fmt.Println(fact(3))
	//fmt.Println(fibo(5))
	rec(5)
}

/*func fact(number int) int {
	if number == 1 || number == 0 {
		return 1
	}
	if number < 0 {
		fmt.Println("Invalid Number")
		return -1
	}
	result := number * fact(number-1)

	return result
}*/

// func fibo(n int) int {
// 	if n == 0 || n == 1 {
// 		return n
// 	}
// 	res := fibo(n-1) + fibo(n-2)
// 	return res
// }

func rec(num int) {
	if num == 0 {
		return
	}
	if num%2 == 0 {
		//fmt.Println(num + 1)
		rec(num - 1)
		//rec(num - 2)
		fmt.Println(num - 1)
	} else {
		//fmt.Println(num + 2)
		rec(num - 1)
		//rec(num - 2)
		fmt.Println(num - 1)
	}
	fmt.Println(num - 1)

}
