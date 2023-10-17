package main

import (
	"fmt"
	"io"
)

type student struct {
}

func (s student) Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (s student) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func main() {
	var s student
	var r io.Reader = s
	var w io.Writer = s
	var rw io.ReadWriter = s
	_, _, _ = r, w, rw

	var dw DoWalk = s
	var dr DoRun = s
	var dwr DoWalkRun = s
	_, _, _ = dw, dr, dwr
	s.Walk()
	s.Run()

}

type DoWalk interface {
	Walk()
}

func (s student) Walk() {
	fmt.Println("student walks")
}

type DoRun interface {
	Run()
}

func (s student) Run() {
	fmt.Println("student runs")
}

type DoWalkRun interface {
	DoWalk
	DoRun
}
