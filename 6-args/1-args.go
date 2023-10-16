package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(os.Args)
	// 	fmt.Printf("%T \n", os.Args)
	// 	fmt.Println(os.Args[0])
	// 	fmt.Println(os.Args[1])
	// 	fmt.Println(os.Args[2])
	// 	fmt.Println(os.Args[3])

	// a := os.Args[1:]

	// fmt.Println(a[0])
	// fmt.Println(a[1])
	// fmt.Print(a[2])

	data := os.Args[0:]

	// if len(data) != 3 {
	// 	log.Println("all values given")
	// 	return
	// }
	fmt.Println(data)

}
