package models

import "errors"

var studentData = map[uint64]Student{101: {101, "sahil", 22, "A"},
	102: {101, "king", 23, "A+"}}

func FetchStudent(stud_id uint64) (Student, error) {

	stud, ok := studentData[stud_id]

	if !ok {
		return Student{}, errors.New("student not found with that id")
	}

	return stud, nil

}
