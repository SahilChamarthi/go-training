package main

import (
	"fmt"
)

func main() {

	a := []int{10, 20, 30, 40, 50}

	fmt.Printf("%p\n", a)

	//a = append(a, 60)

	i := a[1:5]
	fmt.Printf("%p\n", i)

	i[0] = 999

	fmt.Println(a)
	fmt.Println(i)
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", i)

	fmt.Println(cap(a), len(a))
	fmt.Println(cap(i), len(i))

}
