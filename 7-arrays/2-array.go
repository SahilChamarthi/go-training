package main

import "fmt"

func main() {

	a := []int{10, 20, 30}
	b := a
	a = append(a, 100)

	fmt.Println(cap(a))
	fmt.Println(b)
}
