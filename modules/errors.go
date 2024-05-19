package modules

import (
	"errors"
	"fmt"
)

func Errors() {

	err01 := errors.New("something wrong")
	fmt.Printf("%[1]p %[1]T %[1]v\n", err01)
	fmt.Println("----------")
}
