package main

import "fmt"

func main() {
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	// To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings
	//  of length 3 (initially zero-valued).
	fmt.Println("###Make a slice of strings of length 3 initally zero values###")
	s := make([]string, 3)
	fmt.Println("empty", s)

	fmt.Println("###Set and get just slices just like arrays###")
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get s[2]:", s[2])

	fmt.Println("len(s) =", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy from c --> s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy from c --> s:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	//  slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("slices up to (but excluding) s[5]=", l)

	fmt.Println(len(l))
	// counts from s[0], s
	l = s[2:]
	fmt.Println("Print from s[2] to the last element:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declare and initialize", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD=", twoD)

}
