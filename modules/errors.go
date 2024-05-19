package modules

import (
	"errors"
	"fmt"
)

func Errors() {

	err01 := errors.New("something wrong")
	fmt.Printf("pointer: %[1]p\ntype: %[1]T\nvalue: %[1]v\n", err01)
	fmt.Println(err01.Error())
	fmt.Println(err01)
	fmt.Println("----------")
	err02 := errors.New("something wrong")
	fmt.Printf("pointer: %[1]p\ntype: %[1]T\nvalue: %[1]v\n", err01)
	fmt.Printf("err01: %[1]p\nerr02: %[1]p\n", err01, err02)
	fmt.Println(err01 == err02)

	err0 := fmt.Errorf("add info: %w", errors.New("origin error"))
	fmt.Printf("pointer: %[1]p\ntype: %[1]T\nvalue: %[1]v\n", err0)

}
