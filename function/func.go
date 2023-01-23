package main

import "fmt"

//function ---> func keyword is used

func main() {
	w := new(int)
	name := new(string)
	t := "asdfgh"
	result, x := c(w, &t)
	fmt.Println(result)
	fmt.Println(x)
	fmt.Println(*name)

	// r, _ := b(1, 2, 3, 4, 5)
	// fmt.Print(r)
}
func a() (int, string) {
	return 122, "anagha"
}

func b(args ...int) (bool, int) {
	for _, value := range args {
		fmt.Println(value)
	}
	return true, 11

}
func c(w *int, name *string) (i int, j string) {
	i = 10
	j = "gaurav"
	*w = 100
	return
}
