package modules

import (
	"fmt"
	"unsafe"
)

func Pointers() {
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
