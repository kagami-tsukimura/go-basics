package modules

import (
	"fmt"
	"unsafe"
)

const SECRET = "abc"

var (
	i int
	s string
	b bool
)

type controller interface {
	speedUp() int
	speedDown() int
}

type vehicle struct {
	speed       int
	enginePower int
}

type bycycle struct {
	speed      int
	humanPower int
}

func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}

func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed
}

func (b *bycycle) speedUp() int {
	b.speed += 10 * b.humanPower
	return b.speed
}

func (b *bycycle) speedDown() int {
	b.speed -= 5 * b.humanPower
	return b.speed
}

func speedUpDown(c controller) {
	fmt.Printf(("current speed: %v\n"), c.speedUp())
	fmt.Printf(("current speed: %v\n"), c.speedDown())
}

func (v vehicle) String() string {
	// v(値レシーバー)のにvehicleにString()追加
	// Sprintf: formatした値をStringとして返す。→標準出力しない
	return fmt.Sprintf("Vehicle currnet speed: %v, enginePower: %v", v.speed, v.enginePower)
}

func checkType(i any) {
	// 型判定
	switch i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown")
	}
}

func Interfaces() {
	// var v vehicle
	// v.speed = 50
	// v.enginePower = 10
	// speedUpDown(&v)

	// 宣言時に構造体のpointerを取得
	v := &vehicle{
		speed:       0,
		enginePower: 10,
	}
	// v: speedUp, speedDownのメソッド実装済のため、controllerのinterfaceを満たすとみなされる
	speedUpDown(v)
	fmt.Println("----------")

	b := &bycycle{
		speed:      0,
		humanPower: 5,
	}
	speedUpDown(b)
	fmt.Println("----------")
	fmt.Println(v)
	fmt.Println("----------")

	var i1 interface{}
	// any: interface{}と同様(別名)
	var i2 any
	fmt.Printf("i1: %v %T %v\n", i1, i1, unsafe.Sizeof(i1))
	fmt.Printf("i2: %v %T %v\n", i2, i2, unsafe.Sizeof(i2))

	fmt.Println("----------")

	checkType(i1)
	checkType(i2)
	i2 = 1
	checkType(i2)
	i2 = "hello"
	checkType(i2)
	i2 = true
	checkType(i2)
	fmt.Println("----------")

	checkType(v)
	checkType(v.speed)
	checkType(v.enginePower)
	fmt.Println("----------")
}
