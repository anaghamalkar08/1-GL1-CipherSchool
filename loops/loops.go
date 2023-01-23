package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}

	}

	j := 0
	for j < 4 {
		if j == 8 {
			continue
		}
		fmt.Println(j)
		break
	}

	nums := []int{1, 2, 3, 4}

	for _, v := range nums {
		if v == 3 {
			break

		}
		fmt.Println(v)
	}

}
