package main

import (
	"errors"
	"fmt"
	"go-basics/modules"
	"os"
	"strings"
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

func funcDefer() {
	// defer: 関数終了時に遅れて実行
	// deferが複数: 下から上に実行
	defer fmt.Println("defer func finally")   // 3
	defer fmt.Println("defer func seminally") // 2
	fmt.Println("Hello World")                // 1
	fmt.Println("----------")
}

func trim(files ...string) []string {
	// 要素数: 0, capacity: 可変長引数の数
	out := make([]string, 0, len(files))
	for _, f := range files {
		// TrimSuffix: 第2引数にある値を末尾要素から取り除く
		out = append(out, strings.TrimSuffix(f, ".csv"))
	}
	return out
}

func fileChecker(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		// string, error
		return "", errors.New("file not found")
	}
	// 関数終了時に、リソースの解放
	defer f.Close()

	// string, error
	return name, nil
}

func addExt(f func(file string) string, name string) {
	// 無名関数を引数に持つ関数
	// 第1引数: 無名関数f、第2引数: name

	// 第1引数の無名関数fに第2引数nameを渡して、無名関数を実行。戻り値を出力
	fmt.Println(f(name))
}

func multiply() func(int) int {
	// 無名関数をreturnする関数

	return func(n int) int {
		return n * 1000
	}
}

func countUp() func(int) int {
	// グローバル変数でも同じ挙動になるが、closureとして扱うことで関数内に閉じ込める
	// →他の箇所から使えないようにする

	count := 0
	return func(n int) int {
		count += n
		return count
	}
}

func functions() {
	funcDefer()

	files := []string{"file1.csv", "file2.csv", "file3.csv"}
	// か延長引数: ...型
	fmt.Printf("trim: %v\n", trim(files...))
	fmt.Println("----------")

	targetFile := "main.go"
	name, err := fileChecker(targetFile)
	// name, err := fileChecker(trim(files...)[0])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("name: %v\n", name)
	fmt.Println("----------")

	i := 1
	// 無名関数
	// NOTE: 無名関数の終わりに(i): 無名関数を呼び出し
	func(i int) {
		fmt.Printf("i: %v\n", i)
	}(i)
	fmt.Println("----------")

	// 無名関数を代入
	f1 := func(i int) int {
		return i + 1
	}
	// NOTE: 後から無名関数を呼び出し: 変数に格納してから呼び出し
	fmt.Printf("f1: %v\n", f1(1))
	fmt.Println("----------")

	f2 := func(file string) string {
		return file + ".csv"
	}
	addExt(f2, trim(files...)[1])
	fmt.Println("----------")

	// f3: 戻り値の無名関数が格納
	f3 := multiply()
	// NOTE: 無名関数を呼び出し
	fmt.Printf("f3: %v\n", f3(10))
	fmt.Println("----------")

	f4 := countUp()
	// NOTE: 無名関数を呼び出し
	for i := 1; i < 10; i++ {
		fmt.Printf("f4: %v\n", f4(i))
	}
	fmt.Println("----------")
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

func interfaces() {
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

func main() {

	// modules.ModulePackage()
	// modules.Variables()
	// modules.Pointers()
	// modules.Slices()
	modules.Structs()
	// functions()
	// interfaces()

}
