package main

import (
	"fmt"
	"time"
)

type UserError struct {
	time_ time.Time
	type_ string
}

func (ue *UserError) Error() string {
	return fmt.Sprint(ue.time_, " Error: ", ue.type_)
}

func execute() error {
	return &UserError{time.Now(), "User Defined Error returned"}
}

func main() {
	if err := execute(); err != nil {
		fmt.Println(err)
	}

}
