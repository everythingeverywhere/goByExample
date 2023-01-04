package main

import "fmt"

func main() {

	i := 1

	for i <= 3 {
		fmt.Print("i := 1 for i <= 3::  ")
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Print("j := 7; j <= 9; j++ ::  ")
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {

		if n%2 == 0 {
			fmt.Println("n%2 = ", n%2, ",")
			fmt.Println("B I A N C A")
			continue
		}
		fmt.Println("n = ", n)
		fmt.Println("B I A N C A")
	}

}
