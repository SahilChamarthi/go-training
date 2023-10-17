package postgres

import (
	"fmt"
	"project-3/stores"
)

type Conn struct {
	db string
}

func (c Conn) Create(u stores.User) error {
	fmt.Println(c.db, "data created", u)
	return nil
}

func (c Conn) Update(u stores.User) error {
	fmt.Println(c.db, "data updated", u)
	return nil
}

func (c Conn) Delete(u stores.User) error {
	fmt.Println(c.db, "data deleted", u)
	return nil
}

func New(db string) Conn {
	return Conn{db: db}
}
