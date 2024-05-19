package modules

import (
	"errors"
	"fmt"
	"os"
)

var ErrCustom = errors.New("not found")

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
	fmt.Println(errors.Unwrap(err0))
	fmt.Printf("Unwrap type: %[1]T\n", errors.Unwrap(err0))
	fmt.Println("----------")

	err1 := fmt.Errorf("add info: %v", errors.New("origin error"))
	fmt.Println(err1)
	fmt.Printf("type: %[1]T\n", err1)
	// 引数にUnwrapしていないメソッド: nil
	// 値を返す: fmt.Errorfで%v → %w
	fmt.Println(errors.Unwrap(err1))
	fmt.Println("----------")

	// %w: ErrCustomにテキスト追加
	err2 := fmt.Errorf("in repository layer: %w", ErrCustom)
	fmt.Println(err2)
	// %w: 付加情報は累積
	err2 = fmt.Errorf("in service layer: %w", err2)
	fmt.Println(err2)

	// err2を自動でUnwrapして、ErrCustomに一致しているか1つずつ確認
	if errors.Is(err2, ErrCustom) {
		fmt.Println("matched")
	}

}

func fileCheck(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in checker: %w", err)
	}
	defer f.Close()
	return nil
}
