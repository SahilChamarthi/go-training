package main

import (
	"fmt"
	"project-1/models"
)

func main() {

	u1 := models.Users{

		Name:        "sahil",
		Password:    "sahil123",
		Email:       "sahil123@gmail.com",
		Permissions: map[string]bool{"admin": true},
	}

	fmt.Println(u1.Name, u1.Password, u1.Email, u1.Permissions["admin"])
	//fmt.Println(u1)

}
