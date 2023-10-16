package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	greet()
	fmt.Println("doinng more work in main func")
	fmt.Println("main end")

}

func greet() {

	data := os.Args[1:] // type []string
	if len(data) != 3 {
		log.Println("please provide name , age , marks")
		os.Exit(1)
		//return // stops the exec of the current func
	}
	name := data[0]
	ageString := data[1]
	marksString := data[2]
	age, err := strconv.Atoi(ageString)

	if err != nil {
		fmt.Println("age conversion not done properly:", err)
		os.Exit(2)
		//return
	}

	marks, err := strconv.Atoi(marksString)
	if err != nil {
		fmt.Println("marks conversion not done properly:", err)
		os.Exit(3)
		//return
	}

	fmt.Println(name, marks, age)

}
