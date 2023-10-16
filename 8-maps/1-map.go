package main

import "fmt"

func main() {

	ar1 := [5]int{1, 2, 3}
	ar2 := [5]int{1, 2, 4}

	fmt.Println(ar1 == ar2)
}
