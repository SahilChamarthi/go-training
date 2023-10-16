package main

import "fmt"

type user struct {
	Name  string
	Age   int
	marks []int
}

func main() {

	ar := make([]int, 5, 10)

	fmt.Println(len(ar), cap(ar), ar)
	fmt.Printf()

}
