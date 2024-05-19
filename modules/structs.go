package modules

import (
	"fmt"
	"unsafe"
)

// struct: 構造体
type Task struct {
	Title    string
	Estimate int
}

func Structs() {
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
