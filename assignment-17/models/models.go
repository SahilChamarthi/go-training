package models

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"s_name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}
