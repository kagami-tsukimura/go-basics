package main

import (
	"errors"
	"fmt"
	"go-basics/calculator"
	"os"
	"strings"
	"unsafe"

	"github.com/joho/godotenv"
)

const SECRET = "abc"

type Os int

const (
	MAC Os = iota + 1
	WINDOWS
	LINUX
)

var (
	i int
	s string
	b bool
)

// struct: 構造体
type Task struct {
	Title    string
	Estimate int
}

func modulePackage() {
	// module, package
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calculator.Offset)
	fmt.Println("----------")

	sumResultNumber := calculator.Sum(1, 2)
	sumResultFloat := calculator.Sum(1, 2.1)
	fmt.Printf("a + b + offset = %g\n", sumResultNumber)
	fmt.Printf("a + b + offset = %g\n", sumResultFloat)
	fmt.Printf("(a * b) + offset = %g\n", calculator.Multiply(1, 2))
	fmt.Println("----------")
}

func variables() {
	// variables
	// var
	var i int
	fmt.Println(i)
	var j int = 2
	fmt.Println(j)
	// 型推論
	var k = 4
	fmt.Println(k)
	fmt.Println("----------")

	// :=
	l := 1
	fmt.Println(l)
	fmt.Printf("i: %v %T\n", i, i)
	ui := uint16(6)
	// v: value
	// t: type
	fmt.Printf("ui: %v %T\n", ui, ui)
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)
	fmt.Println("----------")

	f := 3.1
	fmt.Printf("f: %v %T\n", f, f)
	s := "hello"
	fmt.Printf("s: %v %T\n", s, s)
	b := true
	fmt.Printf("b: %v %T\n", b, b)
	fmt.Println("----------")

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v, title: %v\n", pi, title)
	fmt.Println("----------")

	// 型変換
	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Printf("x: %v, y: %v, z: %v\n", x, y, z)
	fmt.Println("----------")

	// 定数
	fmt.Printf("Mac: %v, Windows: %v, Linux: %v\n", MAC, WINDOWS, LINUX)
	fmt.Println("----------")

	// 変数の値を変更
	// l = 1 → 2
	fmt.Println(l)
	l = 2
	fmt.Println(l)
	l += 1
	fmt.Println(l)
	l *= 2
	fmt.Println(l)
	fmt.Println("----------")
}

func pointers() {
	var ui1 uint16
	// pointer: メモリ内の1byte番地
	// 先頭アドレス
	fmt.Printf("memory address of ui1: %p\n", &ui1)
	var ui2 uint16
	// pointer: メモリ内の1byte番地
	// 先頭アドレス
	fmt.Printf("memory address of ui2: %p\n", &ui2)
	fmt.Println("----------")

	// pointer変数: *<型>
	// NOTE: pointer変数は型の宣言が必須
	// NOTE: pointerで変数の先頭番地が分かり、型で長さが分かる
	var p1 *uint16
	// NIL
	fmt.Printf("value of p1: %v\n", p1)
	// ui1のアドレス情報をp1に代入
	fmt.Println("--- Assign ui1 address information to p1 ---")
	p1 = &ui1
	fmt.Printf("value of p1: %v\n", p1)
	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
	// 先頭アドレス
	fmt.Printf("memory address of p1: %p\n", &p1)
	// dereference: *<pointer> pointer変数が指し示すui1の値
	fmt.Printf("value of ui1(dereference): %v\n", *p1)
	// ui1の値を変更
	*p1 = 1
	fmt.Printf("value of ui1: %v\n", ui1)
	fmt.Println("----------")

	// Double pointer
	// **: p1の先頭アドレスを別のpointer変数(pp1)に格納
	var pp1 **uint16 = &p1
	// p1の先頭アドレス = pp1の先頭アドレス
	fmt.Printf("value of pp1: %v\n", pp1)
	fmt.Printf("value of *pp1: %v\n", *pp1)
	fmt.Printf("size of pp1: %d[bytes]\n", unsafe.Sizeof(pp1))
	fmt.Printf("value of p1(dereference): %v\n", *pp1)
	fmt.Printf("value of ui1(dereference): %v\n", **pp1)
	// ui1の値を変更
	**pp1 = 10
	fmt.Printf("value of ui1: %v\n", ui1)
	fmt.Println("----------")

	// 制御文
	// スコープ内外で異なるメモリ番地
	println("異なるメモリ番地 -> ':='")
	ok, result := true, "A"
	fmt.Printf("memory address of result(out scope): %p\n", &result)
	if ok {
		// := スコープ内でのみ有効
		result := "OK"
		// メモリ番地も異なる
		fmt.Printf("memory address of result(in scope): %p\n", &result)
		// println("OK")
		println(result)
	} else {
		result := "NG"
		println(result)
	}
	// "OK"のスコープ外のため、"A"
	// fmt.Printf("result: %v\n", "A")
	fmt.Printf("result: %v\n", result)
	fmt.Println("----------")

	// スコープ内外で同じメモリ番地
	println("同じメモリ番地 -> '='")
	ok2, result2 := true, "A"
	fmt.Printf("memory address of result(out scope): %p\n", &result)
	if ok2 {
		// = コロンを外すと、同じメモリ番地
		result2 = "OK"
		fmt.Printf("memory address of result(in scope): %p\n", &result)
		println(result2)
	} else {
		result2 = "NG"
		println(result2)
	}
	fmt.Printf("result2: %v\n", result2)

	fmt.Println("----------")
}

func slices() {
	var a1 [3]int
	var a2 = [3]int{10, 20, 30}
	// := 要素の値が必須
	// [...] 要素数を自動で判定するため、要素数の指定が不要
	a3 := [...]int{10, 20}
	fmt.Println(a1, a2, a3)
	fmt.Printf("%v %v %v\n", a1, a2, a3)
	// len: 配列の要素数
	// cap: 配列の容量
	fmt.Printf("len(a1), cap(a1): %v, %v\n", len(a1), cap(a1))
	fmt.Printf("len(a2), cap(a2): %v, %v\n", len(a2), cap(a2))
	fmt.Printf("len(a3), cap(a3): %v, %v\n", len(a3), cap(a3))
	// 配列: 要素数を動的に変更不可、要素数の異なる配列は型も異なる
	fmt.Printf("%T %T %T\n", a1, a2, a3)
	fmt.Println("----------")

	// slice: 要素数を動的に変更可能
	// 要素数が空の配列: slice
	var s1 []int
	s2 := []int{}
	fmt.Printf("s1_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s2, len(s2), cap(s2))
	// カーリーブラケットなし: nil
	fmt.Println("s1 is nil?: ", s1 == nil)
	// カーリーブラケットあり: not nil
	fmt.Println("s2 is nil?: ", s2 == nil)
	fmt.Println("----------")

	// sliceに要素を追加
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("s1_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s1, len(s1), cap(s1))
	fmt.Println("----------")

	// sliceにsliceを追加
	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...)
	fmt.Printf("s1_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s1, len(s1), cap(s1))
	fmt.Println("----------")

	// make: capacity確保
	s4 := make([]int, 0, 2)
	fmt.Printf("s4_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s4, len(s4), cap(s4))
	// capacityを超えた追加も可能
	s4 = append(s4, 1, 2, 3, 4)
	fmt.Printf("s4_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s4, len(s4), cap(s4))
	fmt.Println("----------")

	// [0, 0, 0, 0]
	s5 := make([]int, 4, 6)
	fmt.Printf("s5_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s5, len(s5), cap(s5))
	fmt.Println("----------")

	// NOTE: メモリを切り取ると、メモリを共有する
	// s5: [0, 0, 0, 0]
	// s6: [   0, 0   ]
	// s5[1]とs5[2]がs6[0]s6[1]
	s6 := s5[1:3]
	s6[1] = 10
	// s5[2] = 10になる
	fmt.Printf("s5_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s6, len(s6), cap(s6))
	fmt.Println("----------")

	s6 = append(s6, 2)
	// s5[3] = 2になる
	fmt.Printf("s5_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s6, len(s6), cap(s6))
	fmt.Println("----------")

	// NOTE: メモリをコピーすると、メモリを共有しない
	// capacity省略: lengthに基づいて自動計算
	sc6 := make([]int, len(s5[1:3]))
	fmt.Printf("s5 source of copy: %v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6 dst copy before: %v %v %v\n", sc6, len(sc6), cap(sc6))

	copy(sc6, s5[1:3])
	fmt.Printf("s5 source of copy: %v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6 dst copy after: %v %v %v\n", sc6, len(sc6), cap(sc6))

	// copy後に値を書き換え
	sc6[1] = 10
	// copy元（s5）の値は変更されない
	fmt.Printf("s5 source of copy: %v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6 dst rewrite after: %v %v %v\n", sc6, len(sc6), cap(sc6))
	fmt.Println("----------")

	s5 = make([]int, 4, 6)
	// [1:3:3]指定値(3)-1のインデックスまでメモリ共有
	fs6 := s5[1:3:3]
	fmt.Printf("s5_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("fs6_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", fs6, len(fs6), cap(fs6))
	fmt.Println("----------")

	fs6[0] = 6
	fs6[1] = 7
	fs6 = append(fs6, 8)
	// s5[1]: 6, s5[2]: 7,
	// s5[3]: 0 (メモリ共有はインデックス2までのため、appendは共有されない)
	fmt.Printf("s5_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("fs6_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", fs6, len(fs6), cap(fs6))
	fmt.Println("----------")

	s5[3] = 9
	fmt.Printf("s5_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("fs6_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", fs6, len(fs6), cap(fs6))
	fmt.Println("----------")

	// map(key: string, value: int)
	var m1 map[string]int
	// := 空の{}で明示的に指定
	m2 := map[string]int{}
	// m1 == nil
	fmt.Printf("m1: %v, %v\n", m1, m1 == nil)
	// m2 != nil
	fmt.Printf("m2: %v, %v\n", m2, m2 == nil)
	fmt.Println("----------")

	m2["A"] = 10
	m2["B"] = 20
	m2["C"] = 0
	// len: 3, m2["A"] = 10
	fmt.Printf("m2: %v, len: %v, m2['A']: %v\n", m2, len(m2), m2["A"])
	delete(m2, "A")
	// len: 2, m2["A"] = 0
	fmt.Printf("m2: %v, len: %v, m2['A']: %v\n", m2, len(m2), m2["A"])
	fmt.Println("----------")

	// 存在しない値の0と存在する値の0の区別: valueの第2引数
	// 存在しない: false
	v, ok := m2["A"]
	fmt.Printf("m2['A']_v: %v, ok: %v\n", v, ok)
	// 存在する: true
	v, ok = m2["C"]
	fmt.Printf("m2['C']_v: %v, ok: %v\n", v, ok)
	fmt.Println("----------")

	// mapの取り出し(ループ処理)
	for k, v := range m2 {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
	fmt.Println("----------")
}

func structs() {
	// struct
	task1 := Task{
		Title:    "Learning Golang",
		Estimate: 3,
	}
	task1.Title = "Learning Go"
	fmt.Printf("Task: %v\n", task1)
	// %+v: 構造体のフィールド名も追加で出力
	fmt.Printf("Task: %+v\n", task1)
	fmt.Printf("Task: %[1]T: %[1]v, [%v]\n", task1, task1.Title)
	fmt.Println("----------")

	var task2 Task = task1
	task2.Title = "new"
	// task1, 2は別のメモリ領域
	fmt.Printf("task1: %v, task2: %v\n", task1.Title, task2.Title)
	fmt.Println("----------")

	// pointerアドレスの格納
	task1p := &Task{
		Title:    "Learning concurrency",
		Estimate: 2,
	}
	fmt.Printf("task1p: %v\n", task1p)
	// アドレス表示
	fmt.Printf("task1p: %p\n", &task1p)
	// *で始まるデータ型: pointer変数の型
	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))
	fmt.Println("----------")

	// dereference
	// (*task1p).Title = "Changed Learning"
	// 構造体のフィールドにdeterence: *を省略可
	task1p.Title = "Changed Learning by Task1"
	fmt.Printf("task1p: %v\n", *task1p)
	fmt.Println("----------")

	var task2p *Task = task1p
	task1p.Title = "Changed Learning by Task2"
	// dereference
	fmt.Printf("task1p: %+v\n", task1p)
	fmt.Printf("task2p: %+v\n", task2p)
	fmt.Println("----------")

	// // 値が初期値（3）から変更されない
	// // task receiverはコピーの値を変更するため、元の値は変更なし
	// task1.extendEstimate()
	// fmt.Printf("task1 extendEstimate: %+v\n", task1.Estimate)
	// pointer receiverで元の値を変更する
	fmt.Printf("task1 before: %+v\n", task1.Estimate)
	// 先頭アドレス取得
	// (&task1).extendEstimateByPointer()
	// 自動でpointer変換して取得
	task1.extendEstimateByPointer()
	fmt.Printf("task1 extendEstimateByPointer: %+v\n", task1.Estimate)
	fmt.Println("----------")
}

// func (task Task) extendEstimate() {
// 	task.Estimate += 10
// }

func (task *Task) extendEstimateByPointer() {
	task.Estimate += 10
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

// 無名関数を引数に持つ関数
// 第一引数: 無名関数、第二引数: name
func addExt(f func(file string) string, name string) {
	fmt.Println(f(name))
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

}

func main() {

	// modulePackage()
	// variables()
	// pointers()
	// slices()
	// structs()
	functions()

}
