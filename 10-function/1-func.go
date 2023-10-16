package main

import (
	"fmt"
	"strconv"
)

func main() {

	//fmt.Println(getName("hello", "sahil"))y
	fmt.Println(add("10", "e10"))

	ar := [5]int{4, 1, 78, 7, 2}

	for i := range ar {
		fmt.Println(ar[i])
	}

}

func getName(a, b string) string {

	return a + " " + b
}

func add(a, b string) (int, error) {

	x, er := strconv.Atoi(a)

	if er != nil {
		fmt.Println(a, " is not convertable to int")
		return 0, er
	}

	y, er := strconv.Atoi(b)
	if er != nil {
		fmt.Println(b, " is not convertable to int")
		return 0, er
	}

	return x + y, nil

}
