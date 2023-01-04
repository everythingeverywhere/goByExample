package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("var a [5]int is empty", a)

	a[4] = 100

	fmt.Println("set: a", a)
	fmt.Println("get: a[4]", a[4])

	fmt.Println("length: len(a)=", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declare and initialize on one line: b=", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Two dimentional array", twoD)

}
