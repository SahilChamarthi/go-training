package main

import (
	"project-3/stores"
	"project-3/stores/mysql"
	"project-3/stores/postgres"
)

func main() {

	u := stores.User{
		Name:  "ajay",
		Email: "ajay@email.com",
	}
	m := mysql.New("mysql")
	stores.StorerInterface = m
	stores.StorerInterface.Create(u)

	u1 := stores.User{
		Name:  "sahil",
		Email: "sahil163@gmail.com",
	}

	p := postgres.New("postgres")
	stores.StorerInterface = p
	stores.StorerInterface.Create(u1)

}
